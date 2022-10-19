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

func (delivery *httpDelivery) UserLogin(context *gin.Context) {
	var request http.UserLoginRequest
	var response http.UserLoginResponse

	if err := context.ShouldBindJSON(&request); err != nil {
		log.Print("/login api bind json error:", err)
		response.Status = -1
		response.ErrorMessage = "/login api bind json error:"
		context.JSON(400, response)
		return
	}

	token, err := delivery.UserUsecase.UserLogin(request.UserId, request.Passwd)

	if err != nil {
		log.Print("/login api user_id or password error:", err)
		response.Status = -1
		response.ErrorMessage = "user_id or password error"
		context.JSON(400, response)
		return
	}

	response.Status = 0
	response.Token = string(token)
	context.JSON(200, response)
}

func (delivery *httpDelivery) UserData(context *gin.Context) {
	var response http.UserDataResponse

	userId := context.GetString("user_id")
	userData, err := delivery.UserUsecase.UserData(userId)

	if err != nil {
		log.Print("/data api bind json error:", err)
		response.Status = -1
		response.ErrorMessage = "token invalid or expired"
		context.JSON(400, response)
		return
	}

	response.Status = 0
	response.UserId = userData.UserId
	response.Email = userData.Email
	response.Name = userData.Name
	context.JSON(200, response)
}

func (delivery *httpDelivery) UserAddBook(context *gin.Context) {
	var request http.UserAddBookRequest
	var response http.UserAddBookResponse

	userId := context.GetString("user_id")

	if err := context.ShouldBindJSON(&request); err != nil {
		log.Print("/book api bind json error:", err)
		response.Status = -1
		response.ErrorMessage = "/book api bind json error"
		context.JSON(400, response)
		return
	}

	bookData := model.Book{
		Isbn:   request.Isbn,
		Title:  request.Title,
		Author: request.Author,
	}

	err := delivery.UserUsecase.UserAddBook(userId, &bookData)

	if err != nil {
		log.Print("/book add book error:", err)
		response.Status = -1
		response.ErrorMessage = "add book error"
		context.JSON(400, response)
		return
	}

	response.Status = 0
	context.JSON(200, response)
}

func (delivery *httpDelivery) UserGetBooks(context *gin.Context) {
	var request http.UserGetBooksRequest
	var response http.UserGetBooksResponse
	var booksResponse []http.BookResponse

	userId := context.GetString("user_id")

	if err := context.ShouldBindJSON(&request); err != nil {
		log.Print("/book api bind json error:", err)
		response.Status = -1
		response.ErrorMessage = "/book api bind json error"
		context.JSON(400, response)
		return
	}

	cmpBookData := model.Book{
		Isbn:   request.Isbn,
		Title:  request.Title,
		Author: request.Author,
	}
	booksData, err := delivery.UserUsecase.UserGetBooks(userId, &cmpBookData)

	if err != nil {
		log.Print("/book get books error:", err)
		response.Status = -1
		response.ErrorMessage = "get books error"
		context.JSON(400, response)
		return
	}

	for _, bookData := range booksData {
		booksResponse = append(booksResponse, http.BookResponse{
			Isbn:   bookData.Isbn,
			Title:  bookData.Title,
			Author: bookData.Author,
		})
	}

	response.Status = 0
	response.Books = booksResponse
	context.JSON(200, response)
}
