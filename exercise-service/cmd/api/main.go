package main

import (
	exerciseRoute "exercise-service/internal/app/exercise/route"
	"exercise-service/internal/pkg/middleware"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := fiber.Config{
		ServerHeader:  "Gramedia",
		StrictRouting: true,
		CaseSensitive: true,
	}
	app := fiber.New(config)

	app.Use(logger.New(logger.Config{
		Format:     "${time} ${method} ${path} ${status}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
	}))

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).Send(nil)
	})

	homepage := app.Group("", middleware.SetSecurityHeader)

	exerciseRoute.NewRouteExercise(homepage)

	app.Listen(":" + os.Getenv("APP_PORT"))
}
