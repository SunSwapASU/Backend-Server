package handlers

import (
	"github.com/branogarbo/sunswap_backend/models"
	"github.com/branogarbo/sunswap_backend/prisma"
	"github.com/gofiber/fiber/v2"
)

func CreateCategory(c *fiber.Ctx) error {
	var fields models.CategoryCreate

	if err := c.BodyParser(&fields); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request body parser threw error",
		})
	}

	return nil
}

func UpdateCategory(c *fiber.Ctx) error {
	var search models.CategorySearch

	if err := c.BodyParser(&search); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request body parser threw error",
		})
	}

	return nil
}

func GetCategory(c *fiber.Ctx) error {
	var search models.CategorySearch

	if err := c.BodyParser(&search); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request body parser threw error",
		})
	}

	return nil
}

func GetAllCategories(c *fiber.Ctx) error {
	categories, err := prisma.Client.Category.FindMany().Exec(prisma.Ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Database threw error while finding categories",
		})
	}

	return c.JSON(categories)
}

func DeleteCategory(c *fiber.Ctx) error {
	var search models.CategorySearch

	if err := c.BodyParser(&search); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request body parser threw error",
		})
	}

	return nil
}
