package models

import (
	"time"

	"golang-jwt/utils"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Username     string
	PasswordHash string
	Role         string
}

type TokenClaims struct {
	jwt.RegisteredClaims
	Role string `json:"role"`
	Csrf string `json:"csrf"`
}

const (
	RefreshTokenValidTime = time.Hour * 72
	AuthTokenvalidTime    = time.Minute * 15
)

func GenrerateCSRFSecret() (string, error) {
	return utils.GenerateRandomString(32)
}
