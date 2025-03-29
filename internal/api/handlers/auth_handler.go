// Handlder login, oauth2, email
package handlers

import (
	"coop-gardens-be/internal/models"
	"coop-gardens-be/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	AuthUC *usecase.AuthUsecase
}

type AuthLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (h *AuthHandler) Signup(c echo.Context) error {
	var req struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
		FullName string `json:"full_name" validate:"required"`
		Role     string `json:"role"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	// Create user object
	user := &models.User{
		Email:    req.Email,
		Password: req.Password, // This will be hashed in the usecase
		FullName: req.FullName,
	}

	// Default role if not specified
	role := req.Role
	if role == "" {
		role = "User"
	}

	// Validate that the role is one of the allowed roles
	if role != "User" && role != "Admin" && role != "Farmer" {
		return c.JSON(http.StatusBadRequest, "Invalid role specified")
	}

	err := h.AuthUC.SignupWithRole(user, role)
	if err != nil {
		return c.JSON(http.StatusConflict, err.Error())
	}

	return c.JSON(http.StatusOK, "User registered successfully, please verify email")
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req AuthLoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	token, err := h.AuthUC.Login(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
