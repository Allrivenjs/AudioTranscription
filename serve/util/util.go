package util

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

func NewJError(err error) *fiber.Map {
	if err != nil {
		return &fiber.Map{
			"error": err.Error(),
		}
	}
	return nil
}

func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}
