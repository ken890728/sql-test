package http

import (
	"sql-test/internal/domain/model/jwt"
	"sql-test/internal/domain/usecase"
	"strings"

	"github.com/gin-gonic/gin"
)

type httpDelivery struct {
	UserUsecase usecase.UserUsecase
}

type errorResponse struct {
	Status       int    `json:"status"`
	ErrorMessage string `json:"error_message,omitempty"`
}

func BindHttpHandler(ginHandler *gin.Engine, userUsecase usecase.UserUsecase) {
	delivery := httpDelivery{UserUsecase: userUsecase}

	v1 := ginHandler.Group("/v1")
	{
		userHandler := v1.Group("/users")
		{
			userHandler.POST("/register", delivery.UserRegister)
			userHandler.POST("/login", delivery.UserLogin)
			userHandler.GET("/data", JwtCheckMiddleware(), delivery.UserData)
			userHandler.POST("/book", JwtCheckMiddleware(), delivery.UserAddBook)
			userHandler.GET("/book", JwtCheckMiddleware(), delivery.UserGetBooks)
		}
	}
}

func JwtCheckMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.Request.Header.Get("Authorization")
		var response errorResponse
		if authHeader == "" {
			response.Status = -1
			response.ErrorMessage = "authorization not be null"
			context.JSON(401, response)
			context.Abort()
			return
		}

		bearers := strings.Split(authHeader, " ")
		if len(bearers) != 2 || bearers[0] != "Bearer" {
			response.Status = -1
			response.ErrorMessage = "authorization format error"
			context.JSON(401, response)
			context.Abort()
			return
		}

		authClaims, err := jwt.VerifyToken(bearers[1])
		if err != nil {
			response.Status = -1
			response.ErrorMessage = "token invalid"
			context.JSON(401, response)
			context.Abort()
			return
		}

		context.Set("user_id", authClaims.UserId)
		context.Next()
	}
}
