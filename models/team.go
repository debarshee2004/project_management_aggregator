package models

import "time"

type Team struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Users     []User    `json:"users,omitempty"`
	Projects  []Project `json:"projects,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
