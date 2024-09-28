package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/branogarbo/sunswap_backend/prisma/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// init db
	dbclient := db.NewClient()
	if err := dbclient.Prisma.Connect(); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := dbclient.Prisma.Disconnect(); err != nil {
			log.Fatal(err)
		}
	}()

	ctx := context.Background()

	// init fiber
	s := fiber.New()

	s.Use(cors.New())

	s.Post("/register", func(c *fiber.Ctx) error {
		creds := struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}{}

		err := json.Unmarshal(c.Body(), &creds)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "JSON unmarshal threw error"})
		}

		registeredUser, err := dbclient.User.FindUnique(
			db.User.Email.Equals(creds.Email),
		).Exec(ctx)
		if registeredUser != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "A user with this email address already exists"})
		} else if err != nil && !errors.Is(err, db.ErrNotFound) {
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Database threw error while finding registering user"})
		}

		passwordHash, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 14)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Password could not be hashed"})
		}

		_, err = dbclient.User.CreateOne(
			db.User.Username.Set(creds.Username),
			db.User.Email.Set(creds.Email),
			db.User.Password.Set(string(passwordHash)),
		).Exec(ctx)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Could not create user in database"})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User registered successfully"})
	})

	s.All("*", func(c *fiber.Ctx) error {
		return c.SendString("bruh")
	})

	log.Fatal(s.Listen(":3000"))
}
