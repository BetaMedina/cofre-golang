package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addHeader(ctx *gin.Context) {
	ctx.Header("content-type", "application/json")
}

func HttpResponse(ctx *gin.Context, statusCode int, message map[string]interface{}) {
	addHeader(ctx)
	ctx.JSON(statusCode, message)
	return
}

func InternalError(ctx *gin.Context) {
	addHeader(ctx)
	ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "failed", "message": "Internal server error"})
	return
}

func Ok(ctx *gin.Context, message map[string]interface{}) {
	addHeader(ctx)
	ctx.JSON(http.StatusOK, message)
	return
}

func BadRequest(ctx *gin.Context, message map[string]interface{}) {
	addHeader(ctx)
	ctx.JSON(http.StatusBadRequest, message)
	return
}

func NotFound(ctx *gin.Context, message map[string]interface{}) {
	addHeader(ctx)
	ctx.JSON(http.StatusNotFound, message)
	return
}

func Forbidden(ctx *gin.Context, message map[string]interface{}) {
	addHeader(ctx)
	ctx.JSON(http.StatusForbidden, message)
	return
}
