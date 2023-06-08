package password

import (
	factory "secrets-golang/internal/auth/factories"
	config "secrets-golang/internal/infra"
	password_handlers "secrets-golang/internal/password/handlers"
	password_repository "secrets-golang/internal/password/repositories"
	user_repository "secrets-golang/internal/user/repositories"

	"github.com/gin-gonic/gin"
)

var savePassHandler password_handlers.SavePassword
var listPassHandler password_handlers.ListPassword
var readPassHandler password_handlers.ReadPassword

func init() {
	password_collection := password_repository.NewPasswordRepository(config.GetConnection("passwords"))
	user_collection := user_repository.NewUserRepository(config.GetConnection("users"))
	savePassHandler = password_handlers.NewSavePassword(password_collection, user_collection)
	listPassHandler = password_handlers.NewListPassword(password_collection, user_collection)
	readPassHandler = password_handlers.NewReadPassword(password_collection, user_collection)
}

func PasswordRoutes(app *gin.Engine) {
	app.POST("/api/passwords", factory.AuthorizedMiddlewareFactory.AuthorizedMidldeware, savePassHandler.SavePassword)
	app.GET("/api/passwords/", factory.AuthorizedMiddlewareFactory.AuthorizedMidldeware, listPassHandler.ListPassword)
	app.GET("/api/passwords/:passwordId", factory.AuthorizedMiddlewareFactory.AuthorizedMidldeware, readPassHandler.ReadPassword)
}
