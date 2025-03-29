package handlers

import (
	"net/http"

	"coop-gardens-be/internal/usecase"

	"github.com/labstack/echo/v4"
)

const MaxFileSize = 5 * 1024 * 1024

type UploadImageHandler struct {
	UploadUC *usecase.UploadImageUsecase
}

func NewUploadImageHandler(uploadUC *usecase.UploadImageUsecase) *UploadImageHandler {
	return &UploadImageHandler{UploadUC: uploadUC}
}

func (h *UploadImageHandler) UploadImage(c echo.Context) error {
	// Lấy file từ request
	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid file"})
	}

	// Kiểm tra kích thước file
	if file.Size > MaxFileSize {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "File quá lớn, vui lòng chọn file nhỏ hơn 5MB"})
	}

	// Mở file
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Cannot open file"})
	}
	defer src.Close()

	// Gửi file lên Cloudinary
	url, err := h.UploadUC.UploadImageToCloudinary(src, file.Filename)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Upload failed"})
	}

	// Trả về URL của ảnh
	return c.JSON(http.StatusOK, map[string]string{"url": url})
}
