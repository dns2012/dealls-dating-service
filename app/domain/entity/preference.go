package entity

import "gorm.io/gorm"

type Preference struct {
	gorm.Model
	UserID uint `json:"user_id" db:"user_id"`
	User User
	PreferenceUserID uint `json:"preference_user_id"`
	PreferenceUser User
	PreferenceType uint `json:"preference_type"`
}
