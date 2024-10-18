package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Parse and validate path, query, and JSON body parameters from a Fiber request.
//
// path is an optional pointer to a struct with `params` and `validate` tags for fiber.ParamsParser()
// query is an optional pointer to a struct with `query` and `validate` tags for fiber.QueryParser()
// body is an optional pointer to a struct with `json` and `validate` tags for fiber.BodyParser()
//
// An error is returned if either parsing or validation failed.
// An effort is made to make these error messages useful - the error is designed to be sent to the user.
func Parse(c *fiber.Ctx, path, query, body any) error {
	// These are designed to be created only once and cached
	fiberValidator := validator.New()

	if path != nil {
		if err := c.ParamsParser(path); err != nil {
			return fmt.Errorf("path parsing failed: %w", err)
		}

		if err := fiberValidator.Struct(path); err != nil {
			return fmt.Errorf("path validation failed: %w", err)
		}
	}

	if query != nil {
		if err := c.QueryParser(query); err != nil {
			return fmt.Errorf("query parsing failed: %w", err)
		}

		if err := fiberValidator.Struct(query); err != nil {
			return fmt.Errorf("query validation failed: %w", err)
		}
	}

	if body != nil {
		if err := c.BodyParser(body); err != nil {
			return fmt.Errorf("body parsing failed: %w", err)
		}

		if err := fiberValidator.Struct(body); err != nil {
			return fmt.Errorf("body validation failed: %w", err)
		}
	}

	return nil
}
