package models

type Event struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Name      string `json:"name" gorm:"not null"`
	Date      string `json:"date" gorm:"not null"`
	CreatedAt string `json:"created_at" gorm:"not null"`
	UpdatedAt string `json:"updated_at" gorm:"not null"`
}