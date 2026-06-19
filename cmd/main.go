package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sumitdhamane/saas-platform/configs"
	"github.com/sumitdhamane/saas-platform/internal/auth"
	"github.com/sumitdhamane/saas-platform/internal/database"
	"github.com/sumitdhamane/saas-platform/internal/tenant"
	"github.com/sumitdhamane/saas-platform/internal/user"
)

func main() {
	cfg := configs.LoadConfig()

	err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	app.Post("/tenants", tenant.CreateTenantHandler)

	app.Post("/users", user.CreateUserHandler)

	app.Post(
		"/auth/login",
		auth.LoginHandler(cfg),
	)

	log.Fatal(app.Listen(":" + cfg.AppPort))

}
