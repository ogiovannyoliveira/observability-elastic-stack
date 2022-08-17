package models

import "time"

type Event struct {
	ID        string     `json:"id" gorm:"primary_key"`
	Name      string     `json:"name" gorm:"not null"`
	Date      string     `json:"date" gorm:"not null"`
	CreatedAt time.Time  `json:"created_at" gorm:"not null"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"not null"`
}