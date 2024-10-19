package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"taskManagement/models"
	"time"
)

const (
	// The AccessTokenDuration should be relatively short - think of the access token as a cached identity proof.
	AccessTokenDuration = 12 * time.Hour
)

// AccessTokenClaims satisfies jwt.Claims interface.
type AccessTokenClaims struct {
	jwt.RegisteredClaims
	Type   string             `json:"type"` // "access"
	UserID primitive.ObjectID `json:"uid"`
	Email  string             `json:"email"`
}

func (c AccessTokenClaims) Valid() error {
	if err := baseValidation(c.RegisteredClaims, c.Type, TokenTypeAccess); err != nil {
		return err
	}
	if c.UserID == primitive.NilObjectID || c.Email == "" {
		return errors.New("missing claims")
	}
	return nil
}

func NewAccessToken(user models.User) (string, error) {
	return newSignedToken(AccessTokenClaims{
		RegisteredClaims: baseClaims(AccessTokenDuration),
		Type:             TokenTypeAccess,
		UserID:           user.ID,
		Email:            user.Email,
	})
}

func ParseAccessToken(tokenStr string) (claims AccessTokenClaims, err error) {
	err = parse(tokenStr, &claims)
	return
}
