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
	var req models.User
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	err := h.AuthUC.Signup(&req)
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