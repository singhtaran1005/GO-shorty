package handlers

import (
	"go-url-short/dto"
	"go-url-short/internal/database"
	"os"
	"strconv"
	"time"

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

	// store new user in api-quota if already pressent (check using ip-address) decrement its quota

	r2 := database.CreateClient(1)
	defer r2.Close()
	val, err := r2.Get(c.Context(), c.IP()).Result()
	if err == redis.Nil {
		// so new user store it in redis
		r2.Set(c.Context(), c.IP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
	} else {
		// user found with ip in redis
		val, _ = r2.Get(c.Context(), c.IP()).Result() // get value of quota, how many api calls left
		valInt, _ := strconv.Atoi(val)
		if valInt <= 0 {
			limit, _ := r2.TTL(r2.Context(), c.IP()).Result()
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "Rate limit exceeded",
				"rate_limit_reset": limit / time.Nanosecond / time.Minute})
		}
	}

	// check if input is actual URL
	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL"})
	}

	// ex user can call localhost:blablaba
	// check for domain error
	if !RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "Invalid URL"})
	}

	// enforce http,SSL
	body.URL = EnforceHttp(body.URL)

	// decrement by 1 everytime function is called
	r2.Decr(r2.Context(), c.IP())
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
