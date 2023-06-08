package user

import (
	factory "secrets-golang/internal/auth/factories"
	config "secrets-golang/internal/infra"
	user_handlers "secrets-golang/internal/user/handlers"
	user_repository "secrets-golang/internal/user/repositories"

	"github.com/gin-gonic/gin"
)

var saveUserHandler user_handlers.SaveUser
var getUserHandler user_handlers.GetUser

func init() {
	repository := user_repository.NewUserRepository(config.GetConnection("users"))
	saveUserHandler = user_handlers.NewUser(repository)
	getUserHandler = user_handlers.NewGetUser(repository)

}

func UserRoutes(app *gin.Engine) {
	app.POST("/api/users", saveUserHandler.SaveUser)
	app.GET("/api/users/profile", factory.AuthorizedMiddlewareFactory.AuthorizedMidldeware, getUserHandler.GetUser)
}
