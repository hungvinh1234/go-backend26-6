package usecase

import (
	"go-template/repository"
	"go-template/usecase/account"
	"go-template/usecase/hometown"
	"go-template/usecase/university"
)

// Usecase .
type Usecase struct {
	Account    account.IUsecase
	Hometown   hometown.IUsecase
	University university.IUsecase
}

// New .
func New(repo *repository.Repository) *Usecase {
	return &Usecase{
		Account:    account.New(repo.Account),
		Hometown:   hometown.New(repo.Hometown),
		University: university.New(repo.University),
	}
}
