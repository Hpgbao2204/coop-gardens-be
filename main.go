package main

import (
	"log"
	"os"

	"coop-gardens-be/config"
	"coop-gardens-be/internal/api/handlers"
	"coop-gardens-be/internal/api/routes"
	"coop-gardens-be/internal/repository"
	"coop-gardens-be/internal/usecase"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env file not found — using environment variables from Render")
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://coop-gardens-be-no2t.onrender.com", "http://localhost:3000"},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))
	config.InitDB()

	// User auth
	userRepo := &repository.UserRepository{
		DB: config.DB,
	}
	authUC := &usecase.AuthUsecase{UserRepo: userRepo}
	authHandler := &handlers.AuthHandler{AuthUC: authUC}

	// Crop
	cropRepo := repository.NewCropRepository(config.DB)
	cropUsecase := usecase.NewCropUsecase(cropRepo)
	cropHandler := handlers.NewCropHandler(cropUsecase)

	// Cloudinary image upload
	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		log.Fatal("Error initializing Cloudinary:", err)
	}
	uploadUC := usecase.NewUploadImageUsecase(cld)
	uploadHandler := handlers.NewUploadImageHandler(uploadUC)

	// Season
	seasonRepo := repository.NewSeasonRepository(config.DB)
	seasonUsecase := usecase.NewSeasonUsecase(seasonRepo)
	seasonHandler := handlers.NewSeasonHandler(seasonUsecase)

	// Task
	taskRepo := repository.NewTaskRepository(config.DB)
	taskUsecase := usecase.NewTaskUsecase(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskUsecase)

	// Inventory
	inventoryRepo := repository.NewInventoryRepository(config.DB)
	inventoryUsecase := usecase.NewInventoryUsecase(inventoryRepo)
	inventoryHandler := handlers.NewInventoryHandler(inventoryUsecase)

	// Blog + Feedback module
	blogRepo := repository.NewBlogRepository(config.DB)
	blogUsecase := usecase.NewBlogUsecase(blogRepo)
	blogHandler := handlers.NewBlogHandler(blogUsecase)

	// Product + Order module
	poRepo := repository.NewProductOrderRepository(config.DB)
	poUsecase := usecase.NewProductOrderUsecase(poRepo)
	poHandler := handlers.NewProductOrderHandler(poUsecase)

	// Dashboard
	dashboardRepo := repository.NewDashboardRepository(config.DB)
	dashboardUsecase := usecase.NewDashboardUsecase(dashboardRepo)
	dashboardHandler := handlers.NewDashboardHandler(dashboardUsecase)

	// Crop Growth Log
	cropGrowthLogRepo := repository.NewCropGrowthLogRepository(config.DB)
	cropGrowthLogUsecase := usecase.NewCropGrowthLogUsecase(cropGrowthLogRepo, cropRepo)
	cropGrowthLogHandler := handlers.NewCropGrowthLogHandler(cropGrowthLogUsecase)

	// Define API groups
	apiV1 := e.Group("/v1") // Auth endpoints
	apiV2 := e.Group("/v2") // Feature endpoints

	auth := apiV1.Group("/auth")
	auth.POST("/signup", authHandler.Signup)
	auth.POST("/login", authHandler.Login)

	// Common routes - chỉ cần JWT, không cần role cụ thể
	common := apiV1.Group("/common")
	routes.CommonRoutes(common, userRepo)

	// Register endpoints
	routes.AuthRoutes(apiV1, authHandler)
	routes.AdminRoutes(apiV1.Group("/admin"), userRepo)
	routes.FarmerRoutes(apiV1.Group("/farmer"), userRepo)
	routes.UserRoutes(apiV1.Group("/user"), userRepo)

	routes.CropRoutes(apiV2.Group("/crops"), cropHandler, userRepo, cropGrowthLogHandler)
	routes.SeasonRoutes(apiV2.Group("/seasons"), seasonHandler, userRepo)
	routes.UploadImageRoutes(apiV2.Group("/upload"), uploadHandler)
	routes.TaskRoutes(apiV2.Group("/tasks"), taskHandler, userRepo)
	routes.InventoryRoutes(apiV2.Group("/inventory"), inventoryHandler)
	routes.BlogRoutes(apiV2.Group("/blog"), blogHandler)
	routes.ProductOrderRoutes(apiV2.Group("/product-order"), poHandler)
	routes.DashboardRoutes(apiV2.Group("/dashboard"), dashboardHandler)
	routes.CropGrowthLogRoutes(apiV2.Group("/crop-growth-logs"), cropGrowthLogHandler)

	log.Println("🚀 Server đang chạy tại: http://localhost:8080")
	e.Start(":8080")
}
