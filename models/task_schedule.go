package models

import "time"

type TaskSchedule struct {
    ID        int       `json:"id"`
    TaskID    int       `json:"task_id"`
    UserID    int       `json:"user_id"`
    StartTime time.Time `json:"start_time"`
    EndTime   time.Time `json:"end_time"`
    Status    string    `json:"status"`
}