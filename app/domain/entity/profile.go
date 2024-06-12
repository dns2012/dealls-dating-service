package entity

import (
	"gorm.io/gorm"
	"time"
)

type Profile struct {
	gorm.Model
	UserID uint `json:"user_id" db:"user_id"`
	FullName string `json:"full_name" db:"full_name"`
	ImageUrl string `json:"image_url" db:"image_url"`
	BirthAt time.Time `json:"birth_at" db:"birth_at"`
	Gender uint `json:"gender" db:"gender"`
	Company string `json:"company" db:"company"`
	JobTitle string `json:"job_title" db:"job_title"`
}