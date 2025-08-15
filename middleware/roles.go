package middleware

import "github.com/gofiber/fiber/v2"

func RoleOnly(roles ...string) fiber.Handler {
    return func(c *fiber.Ctx) error {
        userRole := c.Locals("role")
        if userRole == nil {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "Unauthorized",
            })
        }

        // Verifica se o role do usuário está na lista de roles permitidos
        for _, role := range roles {
            if userRole == role {
                return c.Next()
            }
        }

        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
            "error": "Forbidden - insufficient permissions",
        })
    }
}