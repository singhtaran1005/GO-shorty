package handlers

import (
	"go-url-short/dto"
	"go-url-short/internal/database"

	"github.com/asaskevich/govalidator"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
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

func ResolveURL(c *fiber.Ctx) error {

	url := c.Params("url")
	r := database.CreateClient(0)
	defer r.Close()
	val, err := r.Get(c.Context(), url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Shorten url not found in the database"})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot connect to Database"})
	}

	//increment by 1
	rIncr := database.CreateClient(1)
	defer rIncr.Close()

	rIncr.Incr(c.Context(), "counter")

	return c.Redirect(val, 301)
}
