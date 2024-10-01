package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func CreateUser(c *fiber.Ctx) error {
	return nil
}

func UpdateUser(c *fiber.Ctx) error {
	return nil
}

func GetUser(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	username := claims["username"].(string)

	return c.JSON(fiber.Map{"message": "Welcome " + username})
}

func GetAllUsers(c *fiber.Ctx) error {
	return nil
}

func DeleteUser(c *fiber.Ctx) error {
	return nil
}
