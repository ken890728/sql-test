package http

import (
	"sql-test/internal/domain/usecase"

	"github.com/gin-gonic/gin"
)

type httpDelivery struct {
	UserUsecase usecase.UserUsecase
}

func BindHttpHandler(ginHandler *gin.Engine, userUsecase usecase.UserUsecase) {
	delivery := httpDelivery{UserUsecase: userUsecase}

	v1 := ginHandler.Group("/v1")
	{
		userHandler := v1.Group("/users")
		{
			userHandler.POST("/register", delivery.UserRegister)
			// userHandler.POST("/login")
		}
	}
}
