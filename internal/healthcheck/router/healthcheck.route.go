package healthcheck_router

import (
	healthcheck_handlers "secrets-golang/internal/healthcheck/handlers"

	"github.com/gin-gonic/gin"
)

func HealthCheckRoutes(app *gin.Engine) {
	app.GET("/", healthcheck_handlers.HealthCheckHandler)
}
