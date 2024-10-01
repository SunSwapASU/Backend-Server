package routes

import (
	"log"
	"os"

	"github.com/branogarbo/sunswap_backend/prisma"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var s *fiber.App

func Run() {
	prisma.Connect()

	defer func() {
		if err := prisma.Client.Disconnect(); err != nil {
			log.Fatal(err)
		}
	}()

	////////////////////////////////////////////////////////////////

	s = fiber.New()

	jwtCheck := jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		TokenLookup: "cookie:token",
		ContextKey:  "jwt",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": err.Error(),
			})
		},
	})

	s.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("ALLOWED_ORIGINS"),
		AllowCredentials: true,
	}))

	////////////////////////////////////////////////////////////////

	s.Route("/auth", authRoutes)

	p := s.Group("/private", jwtCheck)

	p.Route("/user", userRoutes)
	p.Route("/item", itemRoutes)
	p.Route("/category", categoryRoutes)

	////////////////////////////////////////////////////////////////

	s.All("*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	log.Fatal(s.Listen(":3000"))
}
