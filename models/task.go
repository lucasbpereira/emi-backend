package models

type Task struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    CreatedBy   int    `json:"created_by"`
}