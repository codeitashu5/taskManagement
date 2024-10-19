package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"taskManagement/models"
	"time"
)

const (
	// The RefreshTokenDuration can be longer - the calling application can store this value and use it to request new
	// access tokens. During token refresh, we return a new (accessToken, refreshToken) pair.

	RefreshTokenDuration = 24 * 30 * time.Hour // 30 days
)

// RefreshTokenClaims satisfies jwt.Claims interface.
type RefreshTokenClaims struct {
	jwt.RegisteredClaims
	Type   string             `json:"type"` // "refresh"
	UserID primitive.ObjectID `json:"uid"`
}

func (c RefreshTokenClaims) Valid() error {
	if err := baseValidation(c.RegisteredClaims, c.Type, TokenTypeRefresh); err != nil {
		return err
	}
	if c.UserID == primitive.NilObjectID {
		return errors.New("missing claims")
	}
	return nil
}

func NewRefreshToken(user models.User) (string, error) {
	return newSignedToken(RefreshTokenClaims{
		RegisteredClaims: baseClaims(RefreshTokenDuration),
		Type:             TokenTypeRefresh,
		UserID:           user.ID,
	})
}
