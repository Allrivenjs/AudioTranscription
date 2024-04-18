package util

import (
	"github.com/gofiber/fiber/v2"
	"gopkg.in/asaskevich/govalidator.v9"
)

type Error = govalidator.Error
type Errors = govalidator.Errors

func ValidateInput(ctx *fiber.Ctx, input interface{}) *fiber.Map {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return ErrorsByField(err)
	}
	return nil
}

func ErrorsByField(e error) *fiber.Map {
	m := make(fiber.Map)
	if e == nil {
		return nil
	}
	switch e.(type) {
	case Error:
		m[e.(Error).Name] = e.(Error).Err.Error()
	case Errors:
		for _, item := range e.(Errors).Errors() {
			n := ErrorsByField(item)
			for k, v := range *n {
				m[k] = v
			}
		}
	}
	return &m
}
