package main

import (
	"log"

	"coop-gardens-be/config"
	_ "coop-gardens-be/docs"
	"coop-gardens-be/internal/api/routers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// @title Swagger API Test
// @version 1.0
// @description This is a sample API Coop-Gardens server.
// @host localhost:8080
// @BasePath /api/v1

// @host localhost:8080
// @BasePath /v1
func main() {
	// Load .env
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	// Group API v1
	apiV1 := e.Group("/api/v1")

	routers.LoginRoutes(apiV1)

	// Kết nối DB
	config.InitDB()

	log.Println("🚀 Server đang chạy tại: http://localhost:8080")
	e.Start(":8080")
}
