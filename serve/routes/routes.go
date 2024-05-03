package routes

import (
	"AudioTranscription/serve/security"
	"AudioTranscription/serve/util"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"net/http"
)

type Routes interface {
	Install(app *fiber.App)
}

func AuthRequired(ctx *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:    security.JwtSecretKey,
		SigningMethod: security.JwtSigningMethod,
		TokenLookup:   "header:Authorization",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.
				Status(http.StatusUnauthorized).
				JSON(util.NewJError(err))
		},
	})(ctx)
}
