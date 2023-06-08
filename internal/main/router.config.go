package config

import (
	auth_router "secrets-golang/internal/auth/router"
	healthcheck_router "secrets-golang/internal/healthcheck/router"
	password_router "secrets-golang/internal/password/router"
	user_router "secrets-golang/internal/user/router"

	"github.com/gin-gonic/gin"
)

func initRoutes(app *gin.Engine) {
	healthcheck_router.HealthCheckRoutes(app)
	user_router.UserRoutes(app)
	password_router.PasswordRoutes(app)
	auth_router.AuthRoutes(app)
}

func InitServer() {
	app := gin.Default()
	initRoutes(app)
	app.Run()
}
