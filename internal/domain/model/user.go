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
	Books  []Book `gorm:"many2many:users_books"`
}

func (user *User) HashPassword() {
	hashByte, _ := bcrypt.GenerateFromPassword([]byte(user.Passwd), 8)
	user.Passwd = string(hashByte)
}

func (user *User) VerifyPasswd(inputPasswd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Passwd), []byte(inputPasswd))
	if err != nil {
		return err
	}
	return nil
}
