package entity

import "gorm.io/gorm"

type UserPackage struct {
	gorm.Model
	UserID uint `json:"user_id" db:"user_id"`
	PackageID uint `json:"package_id" db:"package_id"`
	Package Package
}
