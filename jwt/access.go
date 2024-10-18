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
	//
	// Mercury will trust a valid access token for its entire duration and will not check the userDB until a new
	// token is generated. If a user is deleted, they will still have access until their token expires.
	//
	// There is a security/usability tradeoff here: the higher the duration, the greater the chance the token is
	// leaked or abused. But a value which is too low makes it hard to work with the Axle API for development (both
	// our own developers and new consumers of the Axle API that are integrating with Axle for the first time.)
	//
	// Right now, we're willing to err in favor of usability, so this is long enough to cover a workday.
	//
	// TODO: invalidate deleted task more quickly
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
