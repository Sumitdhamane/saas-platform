package user

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserRequest struct {
	TenantID  int64  `json:"tenant_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func CreateUserHandler(c *fiber.Ctx) error {

	var req CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid payload",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to hash password",
		})
	}

	userID, err := CreateUser(
		req.TenantID,
		req.FirstName,
		req.LastName,
		req.Email,
		string(hashedPassword),
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"id":      userID,
		"message": "user created successfully",
	})
}
