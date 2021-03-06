package account

import (
	"context"

	"go-template/model"
)

// Repository .
type Repository interface {
	GetByUserName(ctx context.Context, username string) (*model.Account, error)
	CreateUser(ctx context.Context, account *model.Account) (*model.Account, error)
	GetByUserId(ctx context.Context, id int64) (*model.Account, error)
	UpdateUser(ctx context.Context, account *model.Account) (*model.Account, error)

	GetUserList(ctx context.Context) (*[]model.Account, error)
	DeleteUser(ctx context.Context, account *model.Account) (*model.Account, error)
	GetByUserEmail(ctx context.Context, email string) (*model.Account, error)
	GetRefList(ctx context.Context, username string) (*[]model.Account, error)

	// GetHometownList(ctx context.Context) (*[]model.Hometown, error)
}
