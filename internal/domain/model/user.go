package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId string `gorm:"index"`
	Name   string
	Email  string
	Passwd string
}
