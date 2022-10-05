package repository

import (
	"sql-test/internal/domain/model"
	"sql-test/internal/domain/repository"

	"gorm.io/gorm"
)

type mysqlUserRepo struct {
	DB *gorm.DB
}

func NewMysqlUserRepo(DB *gorm.DB) repository.UserRepo {
	DB.AutoMigrate(&model.User{})
	return &mysqlUserRepo{DB: DB}
}

func (userRepo *mysqlUserRepo) Create(user *model.User) error {
	err := userRepo.DB.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (userRepo *mysqlUserRepo) Delete(ID int) error {
	err := userRepo.DB.Delete(&model.User{}, ID).Error
	if err != nil {
		return err
	}
	return nil
}
