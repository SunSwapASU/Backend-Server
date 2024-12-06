package handlers

import (
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
	var creds models.RegisterCreds

	if err := c.BodyParser(&creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request body parser threw error",
		})
	}

	registeredUser, err := prisma.Client.User.FindFirst(
		db.User.Or(
			db.User.Username.Equals(creds.Username),
			db.User.Email.Equals(creds.Email),
		),
	).Exec(prisma.Ctx)
	if registeredUser != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "A user with this username or email address already exists",
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
	var creds models.LoginCreds

	if err := c.BodyParser(&creds); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request body parser threw error",
		})
	}

	user, err := prisma.Client.User.FindUnique(
		db.User.Email.Equals(creds.Email),
	).With(
		db.User.Items.Fetch(),
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

	itemsJSON, err := c.App().Config().JSONEncoder(user.Items())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not encode user items to JSON",
		})
	}

	JWTclaims := jwt.MapClaims{
		"userId":           user.ID,
		"username":         user.Username,
		"email":            user.Email,
		"items":            string(itemsJSON),
		"firstName":        user.FirstName,
		"lastName":         user.LastName,
		"preferredContact": user.PreferredContact,
		"campusNames":      user.CampusName,
		"major":            user.Major,
		"gradYear":         user.GradYear,
		"bio":              user.Bio,
		"exp":              time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTclaims)

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not sign JWT",
			"error":   err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:  "token",
		Value: signedToken,
	})

	return c.JSON(JWTclaims)
}

func LogoutUser(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:  "token",
		Value: "",
	})

	return c.JSON(fiber.Map{
		"message": "User logged out successfully",
	})
}
