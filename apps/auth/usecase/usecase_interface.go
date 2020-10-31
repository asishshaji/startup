package usecase

import (
	"context"

	model "github.com/asishshaji/startup/models"
)

// UseCase creates interface for authentication
type UseCase interface {
	SignUp(ctx context.Context, username, password string) error
	SignIn(ctx context.Context, username, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*model.User, error)
}
