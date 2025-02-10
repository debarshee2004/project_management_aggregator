package models

import "time"

type Todo struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	Status      string    `json:"status" gorm:"default:'pending'"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	User        User      `json:"user,omitempty"`
	ProjectID   uint      `json:"project_id" gorm:"not null"`
	Project     Project   `json:"project,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
