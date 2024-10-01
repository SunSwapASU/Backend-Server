package routes

import (
	"github.com/branogarbo/sunswap_backend/handlers"
	"github.com/gofiber/fiber/v2"
)

func userRoutes(r fiber.Router) {
	r.Get("/get", handlers.GetUser)
	r.Get("/getAll", handlers.GetAllUsers)
	r.Post("/update", handlers.UpdateUser)
	r.Delete("/delete", handlers.DeleteUser)
}
