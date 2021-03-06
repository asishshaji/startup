package usecase

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/asishshaji/startup/apps/auth"
	"github.com/asishshaji/startup/apps/auth/repository"
	model "github.com/asishshaji/startup/models"
	"github.com/dgrijalva/jwt-go/v4"
)

// AuthClaims creates struct
type AuthClaims struct {
	jwt.StandardClaims
	User *model.User `json:"user"`
}

// AuthUseCase creates struct
type AuthUseCase struct {
	userRepo       repository.UserRepository
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

// NewAuthUseCase is the constructor
func NewAuthUseCase(
	userRepo repository.UserRepository,
	hashSalt string,
	signingKey []byte,
	tokenTTLSeconds time.Duration) *AuthUseCase {
	return &AuthUseCase{
		userRepo:       userRepo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: time.Second * tokenTTLSeconds,
	}
}

// SignUp creates a new user account
func (a *AuthUseCase) SignUp(ctx context.Context, username, password string) error {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))

	user, _ := a.userRepo.GetUser(ctx, username, fmt.Sprintf("%x", pwd.Sum(nil)))

	if user != nil {
		return errors.New("User already exits")
	}

	log.Println("Creating new User")
	user1 := &model.User{
		Username: username,
		Password: fmt.Sprintf("%x", pwd.Sum(nil)),
	}
	return a.userRepo.CreateUser(ctx, user1)

}

// SignIn logs in the user
func (a *AuthUseCase) SignIn(ctx context.Context, username, password string) (string, error) {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	user, err := a.userRepo.GetUser(ctx, username, password)
	if err != nil {
		return "", err
	}

	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(a.signingKey)
}

// ParseToken parses the jwt token from request
func (a *AuthUseCase) ParseToken(ctx context.Context, accessToken string) (*model.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return a.signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, auth.ErrInvalidAccessToken
}
