package handlers

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/branogarbo/sunswap_backend/models"
	"github.com/branogarbo/sunswap_backend/prisma"
	"github.com/branogarbo/sunswap_backend/prisma/db"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *fiber.Ctx) error {
	var creds models.Creds

	err := json.Unmarshal(c.Body(), &creds)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "JSON unmarshal threw error",
		})
	}

	registeredUser, err := prisma.Client.User.FindUnique(
		db.User.Email.Equals(creds.Email),
	).Exec(prisma.Ctx)
	if registeredUser != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "A user with this email address already exists",
		})
	} else if err != nil && !errors.Is(err, db.ErrNotFound) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Database threw error while finding registering user",
		})
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 14)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Password could not be hashed",
		})
	}

	_, err = prisma.Client.User.CreateOne(
		db.User.Username.Set(creds.Username),
		db.User.Email.Set(creds.Email),
		db.User.Password.Set(string(passwordHash)),
	).Exec(prisma.Ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create user in database",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

func LoginUser(c *fiber.Ctx) error {
	var creds models.Creds

	err := json.Unmarshal(c.Body(), &creds)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "JSON unmarshal threw error",
		})
	}

	user, err := prisma.Client.User.FindUnique(
		db.User.Email.Equals(creds.Email),
	).Exec(prisma.Ctx)
	if errors.Is(err, db.ErrNotFound) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Incorrect email address or password",
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Database threw error while finding login user",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Incorrect email address or password",
		})
	}

	JWTclaims := jwt.MapClaims{
		"username": creds.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTclaims)

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not sign JWT",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:  "token",
		Value: signedToken,
	})

	return c.JSON(fiber.Map{"token": signedToken})
}

func LogoutUser(c *fiber.Ctx) error {
	c.ClearCookie("token")

	return c.JSON(fiber.Map{
		"message": "User logged out successfully",
	})
}
