package main

import (
	"context"
	"log"
	"os"

	"github.com/branogarbo/sunswap_backend/prisma/db"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

var (
	s        *fiber.App
	dbclient *db.PrismaClient
	ctx      context.Context = context.Background()
)

func main() {
	godotenv.Load()

	// init db
	dbclient = db.NewClient()
	if err := dbclient.Prisma.Connect(); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := dbclient.Prisma.Disconnect(); err != nil {
			log.Fatal(err)
		}
	}()

	// init fiber
	s = fiber.New()

	s.Use(cors.New(cors.Config{
		AllowOrigins:     "http://127.0.0.1:5500",
		AllowCredentials: true,
	}))

	s.Post("/register", registerUser)
	s.Post("/login", loginUser)
	s.Post("/logout", logoutUser)

	s.Use(jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		TokenLookup: "cookie:token",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": err.Error(),
			})
		},
	}))

	s.Get("/restricted", restricted)

	s.All("*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	log.Fatal(s.Listen(":3000"))
}
