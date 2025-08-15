package handlers

import (
	"github.com/lucasbpereira/emi/database"
	"github.com/lucasbpereira/emi/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Criar Post
func CreatePost(c *fiber.Ctx) error {
	post := new(models.Post)
	if err := c.BodyParser(post); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON inválido"})
	}

	query := `INSERT INTO posts (title, content) VALUES ($1, $2) RETURNING id`
	err := database.DB.QueryRow(query, post.Title, post.Content).Scan(&post.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao inserir post"})
	}

	return c.Status(201).JSON(post)
}

// Listar Posts
func GetPosts(c *fiber.Ctx) error {
	var posts []models.Post
	err := database.DB.Select(&posts, "SELECT * FROM posts")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao buscar posts"})
	}
	return c.JSON(posts)
}

// Buscar Post por ID
func GetPost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post
	err := database.DB.Get(&post, "SELECT * FROM posts WHERE id=$1", id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Post não encontrado"})
	}
	return c.JSON(post)
}

// Atualizar Post
func UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")
	post := new(models.Post)

	if err := c.BodyParser(post); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON inválido"})
	}

	query := `UPDATE posts SET title=$1, content=$2 WHERE id=$3`
	_, err := database.DB.Exec(query, post.Title, post.Content, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao atualizar post"})
	}

	idInt, _ := strconv.Atoi(id)
	post.ID = idInt
	return c.JSON(post)
}

// Deletar Post
func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := database.DB.Exec("DELETE FROM posts WHERE id=$1", id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao deletar post"})
	}
	return c.SendStatus(204)
}
