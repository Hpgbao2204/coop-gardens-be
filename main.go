package main

import (
	"log"

	"coop-gardens-be/config"
	"coop-gardens-be/internal/api/handlers"
	"coop-gardens-be/internal/api/routers"
	"coop-gardens-be/internal/repository"
	"coop-gardens-be/internal/usecase"

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

	// Group API v1
	apiV1 := e.Group("/api/v1")

	routers.AuthRoutes(apiV1, authHandler)

	log.Println("🚀 Server đang chạy tại: http://localhost:8080")
	e.Start(":8080")
}
