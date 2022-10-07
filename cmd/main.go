package main

import (
	"log"
	"sql-test/internal/delivery/http"
	"sql-test/internal/repository"
	"sql-test/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Print("db connect error: ", err)
	}

	router := gin.Default()

	repo := repository.NewMysqlUserRepo(db)
	usecase := usecase.NewUserUsecase(repo)
	http.BindHttpHandler(router, usecase)

	router.Run(":8080")
}
