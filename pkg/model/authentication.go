package model

import (
	"gorm.io/gorm"
)

// User stores the information of a user of the application
type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password []byte `gorm:"unique"`
}
