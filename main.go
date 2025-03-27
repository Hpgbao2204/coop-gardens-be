package main

import (
	"log"

	"coop-gardens-be/config"
	"coop-gardens-be/internal/api/handlers"
	"coop-gardens-be/internal/api/routers"
	"coop-gardens-be/internal/models"
	"coop-gardens-be/internal/repository"
	"coop-gardens-be/internal/usecase"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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

	initRoles(config.DB)

	// Group API v1
	apiV1 := e.Group("/api/v1")

	routers.AuthRoutes(apiV1, authHandler)
	routers.AdminRoutes(apiV1.Group("/admin"), userRepo)
	routers.FarmerRoutes(apiV1.Group("/farmer"), userRepo)
	routers.UserRoutes(apiV1.Group("/user"), userRepo)

	log.Println("🚀 Server đang chạy tại: http://localhost:8080")
	e.Start(":8080")
}

func initRoles(db *gorm.DB) {
	roles := []models.Role{
			{Name: "User"},
			{Name: "Admin"},
			{Name: "Farmer"},
	}

	for _, role := range roles {
			var existingRole models.Role
			if db.Where("name = ?", role.Name).First(&existingRole).RowsAffected == 0 {
					db.Create(&role)
					log.Printf("Created role: %s", role.Name)
			}
	}
}

