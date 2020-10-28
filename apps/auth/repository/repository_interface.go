package repository

import (
	"context"

	model "github.com/asishshaji/startup/models"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, username, password string) (*model.User, error)
}
