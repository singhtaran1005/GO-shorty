package handlers

import (
	"go-url-short/dto"

	"github.com/asaskevich/govalidator"

	"github.com/gofiber/fiber"
)

func Ping(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": "pong"})
}
func ShortenURL(c *fiber.Ctx) error {
	body := new(dto.Request)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse Json"})
	}

	// rate limiting logic
	// check if ip already present
	// max allowed 10 times in 30 min

	// check if input is actual URL
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL"})
	}
	// ex user can call localhost:blablaba

	// check for domain error

	if !RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "Invalid URL"})
	}

	// enforce https,SSL

	body.URL = EnforceHttp(body.URL)
	return nil
}

func ResolveURL() error {
	return nil
}
