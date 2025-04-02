package routes

import (
    "coop-gardens-be/internal/api/handlers"
    "coop-gardens-be/internal/api/middlewares"
    "coop-gardens-be/internal/repository"
    "github.com/labstack/echo/v4"
)

func TaskRoutes(g *echo.Group, taskHandler *handlers.TaskHandler, userRepo *repository.UserRepository) {
    // Apply JWT middleware
    g.Use(middlewares.JWTMiddleware)
    
    // Get tasks by season - available to all authenticated users
    g.GET("/season/:season_id", taskHandler.GetTasksBySeason)
    
    // Create task - requires Admin or Farmer role
    adminFarmerRoutes := g.Group("")
    adminFarmerRoutes.Use(middlewares.RoleMiddleware("Farmer", userRepo))
    
    // Root endpoint for creating tasks
    adminFarmerRoutes.POST("", taskHandler.CreateTask)
    
    // Update task status
    adminFarmerRoutes.PUT("/:task_id/status", taskHandler.UpdateTaskStatus)
}