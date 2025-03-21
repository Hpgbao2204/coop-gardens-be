package handlers

import (
	"log"
	"net/http"

	"coop-gardens-be/usecase"

	"github.com/labstack/echo/v4"
)

type SignupRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

func Signup(c echo.Context) error {
	var req SignupRequest

	if err := c.Bind(&req); err != nil {
		log.Println("Signup bind error:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	log.Println("Signup attempt for:", req.Email)

	err := usecase.Signup(req.Email, req.Username, req.Password)
	if err != nil {
		log.Println("Signup failed:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	log.Println("Signup successful for:", req.Email)
	return c.JSON(http.StatusCreated, map[string]string{"message": "Signup successful"})
}
