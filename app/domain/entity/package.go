package entity

import "gorm.io/gorm"

type Package struct {
	gorm.Model
	Name string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Price uint `json:"price" db:"price"`
	UnlimitedSwap bool `json:"unlimited_swap" db:"unlimited_swap"`
	TotalSwapPerDay uint `json:"total_swap_per_day" db:"total_swap_per_day"`
}
