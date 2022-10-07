package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserId string `gorm:"index:idx_users_user_id_unique,unique"`
	Name   string
	Email  string
	Passwd string
}

func (user *User) HashPassword() {
	hashByte, _ := bcrypt.GenerateFromPassword([]byte(user.Passwd), 8)
	user.Passwd = string(hashByte)
}
