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
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	e := echo.New()
	config.InitDB()

	userRepo := &repository.UserRepository{
		DB: config.DB,
	}
	authUC := &usecase.AuthUsecase{UserRepo: userRepo}
	authHandler := &handlers.AuthHandler{AuthUC: authUC}

	cropRepo := repository.NewCropRepository(config.DB)
	cropUsecase := usecase.NewCropUsecase(cropRepo)
	cropHandler := handlers.NewCropHandler(cropUsecase)

	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		log.Fatal("Error initializing Cloudinary:", err)
	}
	uploadUC := usecase.NewUploadImageUsecase(cld)
	uploadHandler := handlers.NewUploadImageHandler(uploadUC)

	seasonRepo := repository.NewSeasonRepository(config.DB)
	seasonUsecase := usecase.NewSeasonUsecase(seasonRepo)
	seasonHandler := handlers.NewSeasonHandler(seasonUsecase)

	taskRepo := repository.NewTaskRepository(config.DB)
	taskUsecase := usecase.NewTaskUsecase(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskUsecase)

	apiV1 := e.Group("/api/v1")

	routes.AuthRoutes(apiV1, authHandler)
	routes.AdminRoutes(apiV1.Group("/admin"), userRepo)
	routes.FarmerRoutes(apiV1.Group("/farmer"), userRepo)
	routes.UserRoutes(apiV1.Group("/user"), userRepo)
	routes.CropRoutes(apiV1.Group("/crops"), cropHandler, userRepo)
	routes.SeasonRoutes(apiV1.Group("/seasons"), seasonHandler, userRepo)
	routes.UploadImageRoutes(apiV1.Group("/upload"), uploadHandler)
	routes.TaskRoutes(apiV1.Group("/tasks"), taskHandler, userRepo)

	log.Println("🚀 Server đang chạy tại: http://localhost:8080")
	e.Start(":8080")
}
