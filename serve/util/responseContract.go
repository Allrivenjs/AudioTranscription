package util

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func ErrorResponse(err *fiber.Map) *fiber.Map {
	fmt.Println(err)
	return &fiber.Map{
		"errors": err,
		"status": "false",
		"data":   "",
	}
}

func SuccessResponse(data *fiber.Map) *fiber.Map {
	return &fiber.Map{
		"error":  "",
		"status": "true",
		"data":   data,
	}
}
