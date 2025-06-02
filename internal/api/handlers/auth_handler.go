package handlers

import (
	"net/http"

	"coop-gardens-be/internal/usecase"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	AuthUC *usecase.AuthUsecase
}

type AuthLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthSignupRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req AuthLoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request body",
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Validation error: " + err.Error(),
		})
	}

	token, err := h.AuthUC.Login(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "Login failed: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

func (h *AuthHandler) Signup(c echo.Context) error {
	var req AuthSignupRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request body",
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Validation error: " + err.Error(),
		})
	}

	token, err := h.AuthUC.Signup(req.Name, req.Email, req.Password, req.Phone, req.Address)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Signup failed: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}
