package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password,omitempty" gorm:"not null"`
	TeamID    *uint     `json:"team_id"`
	Team      *Team     `json:"team,omitempty"`
	Todos     []Todo    `json:"todos,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
