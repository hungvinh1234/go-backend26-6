package hometown

import (
	"context"

	"go-template/model"
)

// Repository .
type Repository interface {
	GetHometownList(ctx context.Context) (*[]model.Hometown, error)
}
