package dtos

import "github.com/Abrahamthefirst/finecore-practice/internal/entities"

type LoginRequestDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=4,lte=20"`
}

type SignupRequestDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=4,lte=20"`
	Username string `json:"username" validate:"required,gte=3,lte=20"`
}

type SignupResponse struct {
	Message    string             `json:"message"`
	Data       SignupResponseBody `json:"data"`
	StatusCode uint               `json:"statusCode"`
}

type SignupResponseBody struct {
	User entities.User `json:"user"`
}

type LoginResponseBody struct {
	AccessToken string        `json:"access_token"`
	User        entities.User `json:"user"`
}
type LoginOkResponse struct {
	Message    string            `json:"message"`
	Data       LoginResponseBody `json:"data"`
	StatusCode uint              `json:"statusCode"`
}

type LoginServiceOkResponse struct {
	User         *entities.User `json:"user"`
	AccessToken  string         `json:"access_token"`
	RefreshToken string
}
