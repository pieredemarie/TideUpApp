package dto

import "time"

type TaskRequest struct {
	Name      string     `json:"name" binding:"required"`
	Desc      string     `json:"desc"`
	ContextID *int       `json:"context_id"`
	Deadline  *time.Time `json:"deadline"`
}

type UpdateTaskRequest struct {
	Name      *string    `json:"name" binding:"required"`
	Desc      *string    `json:"desc"`
	ContextID *int       `json:"context_id"`
	Deadline  *time.Time `json:"deadline"`
	Completed *bool      `json:"completed"`
}

type ShowAllTasks struct {
	Limit int `json:"limit"`
}

type ShowEbbTasks struct {
	Limit int `json:"limit"`
}
