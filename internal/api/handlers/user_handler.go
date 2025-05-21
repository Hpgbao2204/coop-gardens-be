package handlers

import (
	"coop-gardens-be/internal/models"
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

	// Tạo response với thông tin vai trò và dashboard URL
	type RoleInfo struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
		URL  string `json:"dashboard_url"`
	}

	type ProfileResponse struct {
		User  *models.User `json:"user"`
		Roles []RoleInfo   `json:"roles"`
	}

	response := ProfileResponse{
		User:  user,
		Roles: make([]RoleInfo, 0),
	}

	baseURL := c.Scheme() + "://" + c.Request().Host

	// Thêm thông tin về vai trò và URL tương ứng
	for _, role := range user.Roles {
		roleInfo := RoleInfo{
			ID:   role.ID,
			Name: role.Name,
		}

		// Tạo URL dựa vào role
		switch role.Name {
		case "Admin":
			roleInfo.URL = baseURL + "/v1/admin/dashboard"
		case "Farmer":
			roleInfo.URL = baseURL + "/v1/farmer/dashboard"
		case "User":
			roleInfo.URL = baseURL + "/v1/user/dashboard"
		default:
			roleInfo.URL = baseURL + "/v1/user/profile"
		}

		response.Roles = append(response.Roles, roleInfo)
	}

	return c.JSON(http.StatusOK, response)
}
