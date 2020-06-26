package university

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

func (p pgRepository) GetUniversityList(ctx context.Context) (*[]model.University, error) {
	db := p.getDB(ctx)
	universities := []model.University{}

	// vi vo tinh truoc do dc khoi tao la con tro roi
	err := db.Find(&universities).Error
	//account o day da la con tro san roi
	return &universities, err
}
