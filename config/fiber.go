package config

import (
	"order-service/exception"

	"github.com/gofiber/fiber/v2"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
		// Concurrency:  1024 * 1024,
	}
}
