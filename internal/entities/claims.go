package entities

import "github.com/golang-jwt/jwt/v5"

type TokenClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

type VerificationClaims struct {
	UserID  uint   `json:"user_id"`
	Purpose string `json:"purpose"`
	jwt.RegisteredClaims
}
