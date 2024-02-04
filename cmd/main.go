package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	app := fiber.New()
	app.Use(logger.New())
	Routes(app)

	log.Fatal(app.Listen(os.Getenv("PORT")))
}
