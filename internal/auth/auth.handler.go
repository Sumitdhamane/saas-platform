package auth

import (
	"github.com/gofiber/fiber/v2"

	"github.com/sumitdhamane/saas-platform/configs"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHandler(cfg *configs.Config) fiber.Handler {

	return func(c *fiber.Ctx) error {

		var req LoginRequest

		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "invalid payload",
			})
		}

		token, err := Login(
			req.Email,
			req.Password,
			cfg,
		)

		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"token": token,
		})
	}
}
