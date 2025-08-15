package main

import (
	"github.com/lucasbpereira/emi/database"
	"github.com/lucasbpereira/emi/handlers"
	"github.com/lucasbpereira/emi/middleware"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Variáveis de ambiente (defina antes de rodar)
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "...")
	os.Setenv("DB_NAME", "...")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")

	database.Connect()

	app := fiber.New()

    // Auth
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)

	// Rota protegida
	app.Get("/me", middleware.AuthRequired, handlers.GetMe)
	app.Get("/me/members", middleware.AuthRequired, handlers.GetMembers)

	// Rotas de posts (protegidas)
	app.Post("/posts", middleware.AuthRequired, middleware.RoleOnly("Mestre", "Pastor", "Admin"), handlers.CreatePost)
	app.Get("/posts", handlers.GetPosts)
	app.Get("/posts/:id", handlers.GetPost)
	app.Put("/posts/:id", middleware.AuthRequired, middleware.RoleOnly("Mestre", "Pastor", "Admin"), handlers.UpdatePost)
	app.Delete("/posts/:id", middleware.AuthRequired, middleware.RoleOnly("Mestre", "Pastor", "Admin"), handlers.DeletePost)

	app.Post("/tasks", middleware.AuthRequired, middleware.RoleOnly("Admin", "Pastor", "Profeta", "Evangelista", "Mestre", "Apostolo"), handlers.CreateTask)
    app.Post("/tasks/schedule", middleware.AuthRequired, middleware.RoleOnly("Admin", "Pastor", "Profeta", "Evangelista", "Mestre", "Apostolo"), handlers.CreateTaskSchedule)
    app.Get("/tasks/mine", middleware.AuthRequired, handlers.GetUserSchedules)
    app.Put("/tasks/complete/:id", middleware.AuthRequired, handlers.CompleteSchedule)

	app.Listen(":3000")
}
