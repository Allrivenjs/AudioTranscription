package models

import (
	"gorm.io/gorm"
)

// User struct
// swagger:model
// @Description User struct register in the application
type User struct {
	gorm.Model
	FirstName string `gorm:"size:100;not null" json:"first_name"`
	LastName  string `gorm:"size:100;not null" json:"last_name"`
	Email     string `gorm:"size:100;unique;not null" json:"email"`
	Password  string `gorm:"not null" json:"-"`
}
