package handlers

import (
	"log"
	"net/http"

	"coop-gardens-be/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func Login(c echo.Context) error {
	var req LoginRequest

	if err := c.Bind(&req); err != nil {
		log.Println("Bind error:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	log.Println("Login attempt for:", req.Email)

	db := c.Get("db").(*gorm.DB)
	token, err := usecase.Login(db, req.Email, req.Password)
	if err != nil {
		log.Println("Login failed:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	log.Println("Login successful for:", req.Email)

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
