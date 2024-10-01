package routes

import (
	"log"
	"os"

	"github.com/branogarbo/sunswap_backend/handlers"
	"github.com/branogarbo/sunswap_backend/prisma"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var s *fiber.App

func Run() {
	prisma.Connect()

	defer func() {
		if err := prisma.Client.Disconnect(); err != nil {
			log.Fatal(err)
		}
	}()

	////////////////////////////////////////////////////////////////

	s = fiber.New()

	jwtCheck := jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		TokenLookup: "cookie:token",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": err.Error(),
			})
		},
	})

	s.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("ALLOWED_ORIGINS"),
		AllowCredentials: true,
	}))

	////////////////////////////////////////////////////////////////

	authRoutes := s.Group("/auth")
	privateRoutes := s.Group("/private", jwtCheck)
	creators := privateRoutes.Group("/create")
	readers := privateRoutes.Group("/read")
	updaters := privateRoutes.Group("/update")
	deleters := privateRoutes.Group("/delete")

	creators.Post("/item", handlers.CreateItem)
	creators.Post("/category", handlers.CreateCategory)

	readers.Get("/user", handlers.GetUser)
	readers.Get("/users", handlers.GetAllUsers)
	readers.Get("/item", handlers.GetItem)
	readers.Get("/items", handlers.GetAllItems)
	readers.Get("/category", handlers.GetCategory)
	readers.Get("/categories", handlers.GetAllCategories)

	updaters.Post("/user", handlers.UpdateUser)
	updaters.Post("/item", handlers.UpdateItem)
	updaters.Post("/category", handlers.UpdateCategory)

	deleters.Delete("/user", handlers.DeleteUser)
	deleters.Delete("/item", handlers.DeleteItem)
	deleters.Delete("/category", handlers.DeleteCategory)

	authRoutes.Post("/register", handlers.RegisterUser)
	authRoutes.Post("/login", handlers.LoginUser)
	authRoutes.Post("/logout", handlers.LogoutUser)

	////////////////////////////////////////////////////////////////

	s.All("*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	log.Fatal(s.Listen(":3000"))
}
