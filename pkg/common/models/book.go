package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model         // adds ID, created_at, ect...
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}
