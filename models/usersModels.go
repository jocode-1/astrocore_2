package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	id           uint `gorm:"primaryKey"`
	firstname    string
	lastname     string
	Phone_number uint
	usernumber   uint
	password     uint
	dob          uint
}
