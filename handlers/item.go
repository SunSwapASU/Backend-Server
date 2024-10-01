package handlers

import (
	"github.com/branogarbo/sunswap_backend/models"
	"github.com/branogarbo/sunswap_backend/prisma"
	"github.com/gofiber/fiber/v2"
)

func CreateItem(c *fiber.Ctx) error {
	var search models.ItemCreate

	if err := c.BodyParser(&search); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request body parser threw error",
		})
	}

	return nil
}

func UpdateItem(c *fiber.Ctx) error {
	var search models.ItemSearch

	if err := c.BodyParser(&search); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request body parser threw error",
		})
	}

	return nil
}

func GetItem(c *fiber.Ctx) error {
	var search models.ItemSearch

	if err := c.BodyParser(&search); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request body parser threw error",
		})
	}

	return nil
}

func GetAllItems(c *fiber.Ctx) error {
	items, err := prisma.Client.Item.FindMany().Exec(prisma.Ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Database threw error while finding items",
		})
	}

	return c.JSON(items)
}

func DeleteItem(c *fiber.Ctx) error {
	var search models.ItemSearch

	if err := c.BodyParser(&search); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request body parser threw error",
		})
	}

	return nil
}
