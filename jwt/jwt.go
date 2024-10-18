// Package jwt handles token parsing and generation for access, refresh, and link tokens.
package jwt

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"taskManagement/envornment"
	"taskManagement/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenResponse struct {
	AccessToken           string `json:"accessToken"`
	ExpiresIn             int    `json:"expiresIn"`
	RefreshToken          string `json:"refreshToken"`
	RefreshTokenExpiresIn int    `json:"refreshTokenExpiresIn"`
}

const (
	// TokenTypeAccess is the login/bearer token - the token used to authenticate with the Task Management .
	TokenTypeAccess = "access"

	// TokenTypeRefresh is the refresh token, which can be used to generate new access tokens after they expire.
	TokenTypeRefresh = "refresh"
)

// The API service associated with the given environment.

// Every token includes issuer and expiration information.
func baseClaims(duration time.Duration) jwt.RegisteredClaims {
	now := time.Now()

	return jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(duration)),
	}
}

func newSignedToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(envornment.GetJwtKey()))
}

func baseValidation(claims jwt.RegisteredClaims, tokenType, expectedType string) error {
	if err := claims.Valid(); err != nil {
		return err
	}
	if tokenType != expectedType {
		return errors.New("invalid type")
	}
	return nil
}

func parse(tokenStr string, claims jwt.Claims) error {
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
		return []byte(envornment.GetJwtKey()), nil
	}, jwt.WithValidMethods([]string{"HS512"}))

	if err != nil || !token.Valid {
		return fmt.Errorf("invalid token: %w", err)
	}

	return nil
}

func SendJWT(c *fiber.Ctx, user models.User) error {
	accessToken, err := NewAccessToken(user)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	refreshToken, err := NewRefreshToken(user)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	token := TokenResponse{
		AccessToken:           accessToken,
		ExpiresIn:             int(AccessTokenDuration.Seconds()),
		RefreshToken:          refreshToken,
		RefreshTokenExpiresIn: int(RefreshTokenDuration.Seconds()),
	}

	c.Cookie(&fiber.Cookie{
		Name:    "jwt",
		Value:   accessToken,
		Expires: time.Now().Add(AccessTokenDuration),
	})

	return c.Status(http.StatusCreated).JSON(token)
}
