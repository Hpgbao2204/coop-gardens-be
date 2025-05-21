package handlers

import (
	"coop-gardens-be/internal/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserRepo *repository.UserRepository
}

func NewUserHandler(userRepo *repository.UserRepository) *UserHandler {
	return &UserHandler{UserRepo: userRepo}
}

func (h *UserHandler) GetUserProfile(c echo.Context) error {
	// Lấy userID từ JWT context (đã được set bởi middleware)
	userID, ok := c.Get("user_id").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "User not authenticated"})
	}

	user, err := h.UserRepo.GetUserByID(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	// Loại bỏ password khỏi response
	user.Password = ""

	return c.JSON(http.StatusOK, user)
}
