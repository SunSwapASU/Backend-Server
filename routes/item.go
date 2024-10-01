package routes

import (
	"github.com/branogarbo/sunswap_backend/handlers"
	"github.com/gofiber/fiber/v2"
)

func itemRoutes(r fiber.Router) {
	r.Get("/create", handlers.CreateItem)
	r.Get("/get", handlers.GetItem)
	r.Get("/getAll", handlers.GetAllItems)
	r.Post("/update", handlers.UpdateItem)
	r.Delete("/delete", handlers.DeleteItem)
}
