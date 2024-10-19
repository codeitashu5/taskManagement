package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strings"
	"taskManagement/envornment"
	"taskManagement/jwt"
	"taskManagement/models"
	"taskManagement/mongoClient"
	"time"
)

const contextKey = "task-jwt-claims"

func Recover(c *fiber.Ctx) (err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("PANIC: %s %s: %v\n", c.Method(), c.OriginalURL(), r)
			err = c.SendStatus(http.StatusInternalServerError)
		}
	}()

	return c.Next()
}

// ParseToken ensures the JWT is valid and adds it to the Fiber context.
//
// This does not check the user against the DB, nor does it authorize the request.
func ParseToken(c *fiber.Ctx) error {
	token := utils.CopyString(strings.TrimPrefix(c.Get("Authorization"), "Bearer "))
	claims, err := jwt.ParseAccessToken(token)
	if err != nil {
		return c.Status(http.StatusUnauthorized).SendString(err.Error())
	}

	// Check if the token has an expiration time
	// Check if claims.ExpiresAt is nil
	if claims.ExpiresAt == nil {
		return c.Status(http.StatusUnauthorized).SendString("Token does not have an expiration time")
	}

	// Check if the token has expired
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return c.Status(http.StatusUnauthorized).SendString("Token expired")
	}

	// get the collection validate user
	userCollection := mongoClient.MongoDB.Database("taskManagmentDb").Collection(envornment.GetUserCollection())

	// check if the user is valid or not
	user := models.User{}
	err = userCollection.FindOne(context.Background(), bson.M{"_id": claims.UserID}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).SendString(err.Error())
		}

		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	// Attach the parsed JWT claims to the context for routes to read as needed
	c.Locals(contextKey, claims)
	return c.Next()
}

// ClaimsFromContext will read the JWT claims from the Fiber context.
//
// Each protected route handler should call this to validate the claim against the request details.
func ClaimsFromContext(c *fiber.Ctx) jwt.AccessTokenClaims {
	return c.Locals(contextKey).(jwt.AccessTokenClaims)
}
