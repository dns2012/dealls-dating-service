package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	PremiumAt time.Time `json:"premium_at" db:"premium_at"`
	Nickname string `json:"nickname" db:"nickname"`
	Email string `json:"email" db:"email" gorm:"unique"`
	Password string `json:"password" db:"password"`
	Profile Profile
	Preference []Preference
	UserPackage UserPackage
}