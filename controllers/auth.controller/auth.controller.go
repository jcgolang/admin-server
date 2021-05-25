package auth_controller

import (
	"admin-server/models"
	userService "admin-server/services/user.service"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.NewUser()
	user.Name = data["name"]
	user.Email = data["email"]
	user.Password = password

	if err := userService.Create(user); err != nil {
		return err
	}

	return c.JSON(user)
}
