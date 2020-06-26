package hometown

import (
	"context"
	"go-template/model"

	"github.com/jinzhu/gorm"
)

func NewPG(getDB func(ctx context.Context) *gorm.DB) Repository {
	return pgRepository{getDB}
}

type pgRepository struct {
	getDB func(ctx context.Context) *gorm.DB
}

func (p pgRepository) GetHometownList(ctx context.Context) (*[]model.Hometown, error) {
	db := p.getDB(ctx)
	hometowns := []model.Hometown{}

	// vi vo tinh truoc do dc khoi tao la con tro roi
	err := db.Find(&hometowns).Error
	//account o day da la con tro san roi
	return &hometowns, err
}
