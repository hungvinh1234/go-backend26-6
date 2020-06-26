package university

import (
	"context"

	"go-template/model"
)

// Repository .
type Repository interface {
	GetUniversityList(ctx context.Context) (*[]model.University, error)
}
