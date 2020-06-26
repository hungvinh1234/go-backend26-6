package repository

import (
	"context"

	"github.com/jinzhu/gorm"

	"go-template/repository/account"
	"go-template/repository/hometown"
	"go-template/repository/university"
)

// Repository .
type Repository struct {
	Account    account.Repository
	Hometown   hometown.Repository
	University university.Repository
}

// New .
func New(getClient func(ctx context.Context) *gorm.DB) *Repository {
	return &Repository{
		Account:    account.NewPG(getClient),
		Hometown:   hometown.NewPG(getClient),
		University: university.NewPG(getClient),
	}
}
