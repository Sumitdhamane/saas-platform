package tenant

import "github.com/gofiber/fiber/v2"

type CreateTenantRequest struct {
	Name string `json:"name"`
}

func CreateTenantHandler(c *fiber.Ctx) error {

	var req CreateTenantRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid payload",
		})
	}

	id, err := CreateTenant(req.Name)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"id": id,
	})
}