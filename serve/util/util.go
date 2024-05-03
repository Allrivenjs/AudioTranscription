package util

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
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

func NormalizeUrl(url string) string {
	return fmt.Sprintf("%s%s", os.Getenv("APP_URl"), url)
}
