package routes

import (
	"github.com/branogarbo/sunswap_backend/handlers"
	"github.com/gofiber/fiber/v2"
)

func authRoutes(r fiber.Router) {
	r.Post("/register", handlers.RegisterUser)
	r.Post("/login", handlers.LoginUser)
	r.Post("/logout", handlers.LogoutUser)
	r.Post("/setup_profile", handlers.ProfileSetup)
}
