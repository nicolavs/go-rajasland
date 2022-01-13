package models

import "gorm.io/gorm"

// User model
type User struct {
	gorm.Model

	Title  string `json:"title"`
	Author string `json:"author"`
}
