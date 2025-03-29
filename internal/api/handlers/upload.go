package handlers

import (
	"net/http"

	"coop-gardens-be/internal/usecase"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/labstack/echo/v4"
)

type UploadHandler struct {
	UploadUC   *usecase.UploadUsecase
	Cloudinary *cloudinary.Cloudinary
}

func (h *UploadHandler) UploadImage(c echo.Context) error {
	// Nhận file từ request
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "File không hợp lệ"})
	}

	// Kiểm tra kích thước file
	if file.Size > usecase.MaxFileSize {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "File quá lớn, vui lòng chọn ảnh dưới 5MB"})
	}

	// Mở file
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Gọi usecase để upload ảnh
	url, err := h.UploadUC.UploadImage(c.Request().Context(), src, file.Size)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"url": url})
}
