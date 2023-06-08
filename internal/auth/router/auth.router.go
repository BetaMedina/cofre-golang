package user

import (
	auth_handlers "secrets-golang/internal/auth/handlers"
	config "secrets-golang/internal/infra"
	user_repository "secrets-golang/internal/user/repositories"

	"github.com/gin-gonic/gin"
)

var auth auth_handlers.Auth

func init() {
	repository := user_repository.NewUserRepository(config.GetConnection("users"))
	auth = auth_handlers.NewAuth(repository, config.NewTokenInfra())
}

func AuthRoutes(app *gin.Engine) {
	app.POST("/api/login", auth.Auth)
}
