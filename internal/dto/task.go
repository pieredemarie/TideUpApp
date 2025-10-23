package dto

import "time"

type CreateTaskRequest struct {
	Name      string     `json:"name" binding:"required"`
	Desc      string     `json:"desc"`
	ContextID *int       `json:"context_id"`
	Deadline  *time.Time `json:"deadline"`
}

type DeleteTaskRequest struct {
	ID int `json:"id"`
}

type ShowAllTasks struct {
	Limit int `json:"limit"`
}