package models

import "time"

type Task struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	ContextID *int `json:"context_id,omitempty"`
	Deadline *time.Time `json:"date,omitempty"`
	Completed bool `json:"completed"`
	CreatedAt *time.Time `json:"created_at"`

	Context Context `gorm:"foreignKey:ID;references:ContextID"`
}


