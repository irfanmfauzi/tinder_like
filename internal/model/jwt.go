package model

import (
	"tinder_like/internal/model/entity"

	"github.com/golang-jwt/jwt/v5"
)

// Struct for custom claims, extend jwt.RegisteredClaims if needed
type CustomClaim struct {
	User   entity.User   `json:"user"`
	Member entity.Member `json:"member"`
	jwt.RegisteredClaims
}
