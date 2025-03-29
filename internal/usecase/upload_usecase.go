package usecase

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

const MaxFileSize = 5 * 1024 * 1024 // Giới hạn 5MB

type UploadUsecase struct {
	Cloudinary *cloudinary.Cloudinary
}

func (uc *UploadUsecase) UploadImage(ctx context.Context, file multipart.File, fileSize int64) (string, error) {
	if fileSize > MaxFileSize {
		return "", errors.New("File quá lớn, vui lòng chọn ảnh dưới 5MB")
	}

	uploadResult, err := uc.Cloudinary.Upload.Upload(ctx, file, uploader.UploadParams{})
	if err != nil {
		return "", fmt.Errorf("Upload thất bại: %v", err)
	}
	return uploadResult.SecureURL, nil
}
