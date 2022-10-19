package main

import (
	"fmt"
	"log"
	"sql-test/internal/delivery/http"
	"sql-test/internal/repository"
	"sql-test/internal/usecase"

	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg, err := ini.Load("../config/env.ini")
	if err != nil {
		log.Print("ini loading error")
	}
	dbUser := cfg.Section("database").Key("db_user").String()
	dbPasswd := cfg.Section("database").Key("db_passwd").String()
	dbHost := cfg.Section("database").Key("db_host").String()
	dbPort := cfg.Section("database").Key("db_port").String()
	dbName := cfg.Section("database").Key("db_name").String()

	httpHost := cfg.Section("http").Key("http_host").String()
	httpPort := cfg.Section("http").Key("http_port").String()

	dsnFormat := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	httpFormat := "%s:%s"

	dsn := fmt.Sprintf(dsnFormat, dbUser, dbPasswd, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Print("db connect error: ", err)
	}

	router := gin.Default()

	userRepo := repository.NewMysqlUserRepo(db)
	bookRepo := repository.NewMysqlBookRepo(db)
	usecase := usecase.NewUserUsecase(userRepo, bookRepo)

	http.BindHttpHandler(router, usecase)

	httpDsn := fmt.Sprintf(httpFormat, httpHost, httpPort)
	router.Run(httpDsn)
}
