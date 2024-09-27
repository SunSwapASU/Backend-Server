package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	s := fiber.New()

	s.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "hello",
		})
	})

	log.Fatal(s.Listen(":3000"))
}
