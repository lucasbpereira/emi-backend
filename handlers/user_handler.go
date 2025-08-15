package handlers

import (
	"github.com/lucasbpereira/emi/database"
	"github.com/lucasbpereira/emi/models"
	"github.com/lucasbpereira/emi/types"

	"github.com/gofiber/fiber/v2"
)



// Listar Usuarios
func GetMembers(c *fiber.Ctx) error {
	var users []types.PublicUser
	err := database.DB.Select(&users, "SELECT id, name, role FROM users where role = 'Membro'")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao buscar usuarios"})
	}
	return c.JSON(users)
}

func GetMe(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Usuário não autenticado",
		})
	}

	var user models.User
	err := database.DB.Get(&user, "SELECT * FROM users WHERE id = $1", userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao buscar usuário"})
	}

	return c.JSON(user)
}