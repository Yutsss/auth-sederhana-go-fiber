package utilities

import (
	"auth-sederhana-go-fiber/dtos"
	errorUtils "auth-sederhana-go-fiber/utilities/error"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

type (
	JWTUtils interface {
		GenerateToken(userID int64) (string, errorUtils.CustomError)
		ExtractFromCookie(ctx *fiber.Ctx) (string, errorUtils.CustomError)
		ValidateToken(token string) (*jwt.Token, errorUtils.CustomError)
		GetPayload(token string) (dtos.AuthPayload, errorUtils.CustomError)
	}

	jwtUtils struct {
		secretKey string
		issuer    string
	}
)

func NewJWTUtils() JWTUtils {
	return &jwtUtils{
		secretKey: loadSecretKey(),
		issuer:    "FP_PBKK_GO",
	}
}

func loadSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		secretKey = "Template"
	}
	return secretKey
}

func loadExpTime() (int, error) {
	expTime := os.Getenv("JWT_EXPIRED")
	if expTime == "" {
		expTime = "1h"
	}

	// Parsing durasi dari string
	duration, err := time.ParseDuration(expTime)
	if err != nil {
		return 0, errorUtils.ErrInternalServer
	}

	// Mengembalikan durasi dalam detik
	return int(duration.Seconds()), nil
}

func (j *jwtUtils) GenerateToken(userId int64) (string, errorUtils.CustomError) {
	expTime, err := loadExpTime()

	if err != nil {
		return "", errorUtils.ErrInternalServer
	}

	claims := dtos.JwtCustomClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expTime) * time.Second)),
			Issuer:    j.issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(j.secretKey))

	if err != nil {
		return "", errorUtils.ErrInternalServer
	}

	return signedToken, nil
}

func (j *jwtUtils) ExtractFromCookie(ctx *fiber.Ctx) (string, errorUtils.CustomError) {
	cookie := ctx.Cookies("access_token")

	if cookie == "" {
		return "", errorUtils.ErrUnauthorized
	}

	return cookie, nil
}

func (j *jwtUtils) parseToken(t_ *jwt.Token) (any, error) {
	if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", t_.Header["alg"])
	}
	return []byte(j.secretKey), nil
}

func (j *jwtUtils) ValidateToken(token string) (*jwt.Token, errorUtils.CustomError) {
	parsedToken, err := jwt.Parse(token, j.parseToken)
	if err != nil {
		return nil, errorUtils.ErrUnauthorized
	}

	if !parsedToken.Valid {
		return nil, errorUtils.ErrUnauthorized
	}

	return parsedToken, nil
}

func (j *jwtUtils) GetPayload(token string) (dtos.AuthPayload, errorUtils.CustomError) {
	t, err := j.ValidateToken(token)
	if err != nil {
		return dtos.AuthPayload{}, err
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return dtos.AuthPayload{}, errorUtils.ErrInternalServer
	}

	return dtos.AuthPayload{
		UserID: int64(claims["user_id"].(float64)),
	}, nil
}
