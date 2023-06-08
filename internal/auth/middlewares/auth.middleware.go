package middlewares

import (
	"net/http"
	"secrets-golang/internal/infra"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthorizedMiddleware interface {
	AuthorizedMidldeware(c *gin.Context)
}
type authorizedMiddleware struct {
	token infra.TokenInfra
}

func (t authorizedMiddleware) AuthorizedMidldeware(c *gin.Context) {
	if c.Request.Header["Authorization"] == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "failed", "message": "Token is required"})
		c.Abort()
		return
	}
	token := c.Request.Header["Authorization"][0]
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Invalid token"})
		c.Abort()
		return
	}
	tokenPrefix := strings.Split(token, " ")
	if tokenPrefix[0] != "Bearer" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Invalid token type"})
		c.Abort()
		return
	}
	validateToken, err := t.token.VerifyToken(tokenPrefix[1])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Token invalid"})
		c.Abort()
		return
	}
	if validateToken.Valid != true {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Token is expired"})
		c.Abort()
		return
	}
	issuer, _ := validateToken.Claims.GetIssuer()
	c.Set("id", issuer)
	c.Next()
}

func NewAuthorizedMiddleware(token infra.TokenInfra) AuthorizedMiddleware {
	return &authorizedMiddleware{token}
}
