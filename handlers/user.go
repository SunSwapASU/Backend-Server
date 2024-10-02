package handlers

import (
	"github.com/branogarbo/sunswap_backend/models"
	"github.com/branogarbo/sunswap_backend/prisma"
	"github.com/branogarbo/sunswap_backend/prisma/db"
	"github.com/gofiber/fiber/v2"
)

func UpdateUser(c *fiber.Ctx) error {
	var search models.UserSearch

	if err := c.BodyParser(&search); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request body parser threw error",
		})
	}

	return nil
}

func GetUser(c *fiber.Ctx) error {
	var search models.UserSearch

	if err := c.BodyParser(&search); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request body parser threw error",
		})
	}

	return nil
}

func GetAllUsers(c *fiber.Ctx) error {
	users, err := prisma.Client.User.FindMany().With(
		db.User.Items.Fetch(),
	).Exec(prisma.Ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Database threw error while finding users",
		})
	}

	return c.JSON(users)
}

func DeleteUser(c *fiber.Ctx) error {
	var search models.UserSearch

	if err := c.BodyParser(&search); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request body parser threw error",
		})
	}

	return nil
}
