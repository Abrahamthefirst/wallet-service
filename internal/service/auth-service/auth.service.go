package authservice

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Abrahamthefirst/finecore-practice/internal/db/repository"
	"github.com/Abrahamthefirst/finecore-practice/internal/dtos"
	"github.com/Abrahamthefirst/finecore-practice/internal/entities"
	custom_http "github.com/Abrahamthefirst/finecore-practice/internal/http/webutil"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepository *repository.UserRepository
}

func NewAuthService(userRepository *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (s *AuthService) SignUp(ctx context.Context, dto dtos.SignupRequestDto) (*entities.User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	user := &entities.User{
		Email:    dto.Email,
		Password: string(hashedPassword),
		Username: dto.Username,
	}

	newUser, err := s.userRepository.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("%w: user with email %s already exists", custom_http.ErrConflict, dto.Email)
	}

	return newUser, nil
}

func (a *AuthService) Login(ctx context.Context, dto dtos.LoginRequestDto) (*dtos.LoginServiceOkResponse, error) {
	user, err := a.userRepository.FindByEmail(ctx, dto.Email)

	if err != nil {
		if errors.Is(err, custom_http.RecordNotFound) {
			return nil, custom_http.ErrInvalidCredentials

		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password)); err != nil {
		return nil, custom_http.ErrInvalidCredentials
	}

	atClaims := entities.TokenClaims{
		UserID: uint(user.ID),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "my-auth-server",
		},
	}

	rtClaims := entities.TokenClaims{
		UserID: uint(user.ID),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "my-auth-server",
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))
	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("REFRESH_TOKEN_SECRET")))

	if err != nil {
		return nil, custom_http.ErrInternal
	}

	return &dtos.LoginServiceOkResponse{User: user, AccessToken: accessTokenString, RefreshToken: refreshTokenString}, nil

}
