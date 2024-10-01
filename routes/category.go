package routes

import (
	"github.com/branogarbo/sunswap_backend/handlers"
	"github.com/gofiber/fiber/v2"
)

func categoryRoutes(r fiber.Router) {
	r.Get("/create", handlers.CreateCategory)
	r.Get("/get", handlers.GetCategory)
	r.Get("/getAll", handlers.GetAllCategories)
	r.Post("/update", handlers.UpdateCategory)
	r.Delete("/delete", handlers.DeleteCategory)
}
