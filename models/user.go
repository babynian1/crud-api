package models 

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama string
	Email string
	Umur int
	
}