package healthcheck_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{"running": true})
}
