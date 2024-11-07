package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims
}

func Generate(userId string, expiresIn time.Duration, signingKey []byte) (string, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "awesomeProject",
			Subject:   userId,
			ID:        uuid.NewString(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}

func GeneratePair(userId string, signingKey []byte) (string, string, error) {
	accessTok, err := Generate(userId, 15*time.Minute, signingKey)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate access token: %w", err)
	}

	refTok, err := Generate(userId, 24*time.Hour, signingKey)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return accessTok, refTok, nil
}

func Verify(key, token string) (jwt.Claims, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return nil, errors.New("jwt token parse")
	}

	if !jwtToken.Valid {
		return nil, errors.New("invalid token")
	}

	return jwtToken.Claims, nil
}
