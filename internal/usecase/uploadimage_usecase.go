package usecase

import (
	"context"
	"io"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type UploadImageUsecase struct {
	Cloudinary *cloudinary.Cloudinary
}

func NewUploadImageUsecase(cld *cloudinary.Cloudinary) *UploadImageUsecase {
	return &UploadImageUsecase{Cloudinary: cld}
}

func (uc *UploadImageUsecase) UploadImageToCloudinary(file io.Reader, filename string) (string, error) {
	uploadResult, err := uc.Cloudinary.Upload.Upload(context.Background(), file, uploader.UploadParams{
		PublicID: filename,
		Folder:   "uploads",
	})
	if err != nil {
		return "", err
	}
	return uploadResult.SecureURL, nil
}
