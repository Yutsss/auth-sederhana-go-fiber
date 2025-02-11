package dtos

import "github.com/golang-jwt/jwt/v4"

type (
	AuthRequest struct {
		AccessToken string
	}

	JwtCustomClaims struct {
		UserID int64 `json:"user_id"`
		jwt.RegisteredClaims
	}

	AuthPayload struct {
		UserID int64 `json:"user_id"`
	}
)
