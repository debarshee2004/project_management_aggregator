package models

import "time"

type Project struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	TeamID      uint      `json:"team_id" gorm:"not null"`
	Team        Team      `json:"team,omitempty"`
	Todos       []Todo    `json:"todos,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
