package university

import (
	"context"
	"go-template/model"
)

// IUsecase . Xu ly logic
// Khai bao cac ham se dung ben usecase
type IUsecase interface {
	ShowUniversityList(ctx context.Context) (*[]model.University, error)
}
