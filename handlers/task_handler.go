package handlers

import (
    "github.com/lucasbpereira/emi/database"

    "github.com/gofiber/fiber/v2"
	"github.com/lucasbpereira/emi/models"
)

func CreateTask(c *fiber.Ctx) error {
    var task models.Task
    if err := c.BodyParser(&task); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Dados inválidos"})
    }

    userID := c.Locals("user_id").(int)

    err := database.DB.QueryRow(`
        INSERT INTO tasks (name, description, created_by)
        VALUES ($1, $2, $3) RETURNING id
    `, task.Name, task.Description, userID).Scan(&task.ID)

    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(task)
}

// Criar escala para tarefa existente
func CreateTaskSchedule(c *fiber.Ctx) error {
    var s models.TaskSchedule
    if err := c.BodyParser(&s); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Dados inválidos"})
    }

    _, err := database.DB.Exec(`
        INSERT INTO task_schedules (task_id, user_id, start_time, end_time, status)
        VALUES ($1, $2, $3, $4, $5)
    `, s.TaskID, s.UserID, s.StartTime, s.EndTime, s.Status)

    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(fiber.Map{"message": "Escala criada com sucesso"})
}


func GetUserSchedules(c *fiber.Ctx) error {
    userIDVal := c.Locals("user_id")
    userID, ok := userIDVal.(int)
    if !ok {
        return c.Status(400).JSON(fiber.Map{"error": "ID de usuário inválido"})
    }

    query := `
        SELECT ts.id, ts.task_id, ts.user_id, ts.start_time, ts.end_time, ts.status,
               t.name AS task_name, t.description
        FROM task_schedules ts
        JOIN tasks t ON t.id = ts.task_id
        WHERE ts.user_id = $1 
          AND ts.status != 'completed' 
          AND ts.end_time >= NOW()
        ORDER BY ts.start_time ASC
    `

    rows, err := database.DB.Query(query, userID)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    defer rows.Close()

    var schedules []fiber.Map
    for rows.Next() {
        var s models.TaskSchedule
        var taskName, desc string

        if err := rows.Scan(&s.ID, &s.TaskID, &s.UserID, &s.StartTime, &s.EndTime, &s.Status, &taskName, &desc); err != nil {
            return c.Status(500).JSON(fiber.Map{"error": err.Error()})
        }

        schedules = append(schedules, fiber.Map{
            "id":          s.ID,
            "task_id":     s.TaskID,
            "task_name":   taskName,
            "description": desc,
            "start_time":  s.StartTime,
            "end_time":    s.EndTime,
            "status":      s.Status,
        })
    }

    return c.JSON(schedules)
}


// Marcar escala como concluída
func CompleteSchedule(c *fiber.Ctx) error {
    id := c.Params("id")
    userID := c.Locals("user_id").(int)

    _, err := database.DB.Exec(`
        UPDATE task_schedules
        SET status = 'completed'
        WHERE id = $1 AND user_id = $2
    `, id, userID)

    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(fiber.Map{"message": "Tarefa concluída"})
}
