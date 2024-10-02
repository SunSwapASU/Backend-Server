package handlers

import (
	"github.com/branogarbo/sunswap_backend/models"
	"github.com/branogarbo/sunswap_backend/prisma"
	"github.com/branogarbo/sunswap_backend/prisma/db"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func CreateItem(c *fiber.Ctx) error {
	var fields models.ItemCreate

	if err := c.BodyParser(&fields); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request body parser threw error",
		})
	}

	ownerId := c.Locals("jwt").(*jwt.Token).Claims.(jwt.MapClaims)["userId"].(string)

	// create item with no category
	// if category does not exist, create category with no items
	// update item to have categoryName
	// update category to append item

	newItem, err := prisma.Client.Item.CreateOne(
		db.Item.Owner.Link(
			db.User.ID.Equals(ownerId),
		),
		db.Item.Name.Set(fields.Name),
		db.Item.Condition.Set(fields.Condition),
		db.Item.Description.Set(fields.Description),
	).Exec(prisma.Ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Database could not create item",
		})
	}

	newCategory, err := prisma.Client.Category.CreateOne(
		db.Category.Name.Set(fields.CategoryName),
		db.Category.Description.Set(fields.Description),
	).Exec(prisma.Ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			// "message": "Database could not create category",
			"error": err.Error(),
		})
	}

	_, err = prisma.Client.Item.FindUnique(
		db.Item.ID.Equals(newItem.ID),
	).Update(
		db.Item.Category.Link(
			db.Category.Name.Equals(fields.CategoryName),
		),
	).Exec(prisma.Ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Database could not update item to include new category",
		})
	}

	_, err = prisma.Client.Category.FindUnique(
		db.Category.ID.Equals(newCategory.ID),
	).With(
		db.Category.Items.Fetch(),
	).Update(
		db.Category.Items.Push(newItem),
	).Exec(prisma.Ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Database could not update category",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Item created successfully",
	})
}

func UpdateItem(c *fiber.Ctx) error {
	var fields models.ItemSearch

	if err := c.BodyParser(&fields); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request body parser threw error",
		})
	}

	return nil
}

func GetItem(c *fiber.Ctx) error {
	var fields models.ItemSearch

	if err := c.BodyParser(&fields); err != nil {
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
	var fields models.ItemSearch

	if err := c.BodyParser(&fields); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request body parser threw error",
		})
	}

	return nil
}
