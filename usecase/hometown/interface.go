package hometown

import (
	"context"
	"go-template/model"
)

// IUsecase . Xu ly logic
// Khai bao cac ham se dung ben usecase
type IUsecase interface {
	ShowHometownList(ctx context.Context) (*[]model.Hometown, error)

	// Upload(form *multipart.Form) (*UploadResponse, error)
}
