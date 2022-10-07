package http

import (
	"log"
	"sql-test/internal/domain/model"
	"sql-test/internal/domain/model/http"

	"github.com/gin-gonic/gin"
)

func (delivery *httpDelivery) UserRegister(context *gin.Context) {
	var request http.UserRegisterRequest
	var response http.UserRegisterResponse

	if err := context.ShouldBindJSON(&request); err != nil {
		log.Print("/register api bind json error:", err)
		response.Status = -1
		response.ErrorMessage = "/register api bind json error:"
		context.JSON(400, response)
		return
	}

	userData := model.User{
		UserId: request.UserId,
		Passwd: request.Passwd,
		Name:   request.Name,
		Email:  request.Email,
	}

	if err := delivery.UserUsecase.UserRegister(&userData); err != nil {
		log.Print("/register api bind json error:", err)
		response.Status = -1
		response.ErrorMessage = err.Error()
		context.JSON(400, response)
		return
	}

	response.Status = 0
	context.JSON(200, response)
}
