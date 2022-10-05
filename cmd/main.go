package main

import (
	"log"
	"sql-test/internal/domain/model"
	"sql-test/internal/repository"
	"sql-test/internal/usecase"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Print("db connect error: ", err)
	}

	repo := repository.NewMysqlUserRepo(db)
	usecase := usecase.NewUserUsecase(repo)

	user := model.User{
		UserId: "test01",
		Name:   "測試使用者",
		Email:  "test@email.com",
	}
	usecase.UserRegister(&user)
}
