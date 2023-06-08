package user

import (
	repository "secrets-golang/internal/user/repositories"
	"secrets-golang/internal/utils"

	"github.com/gin-gonic/gin"
)

type contextCredential struct {
	Data string `json:"data"`
}

type GetUser interface {
	GetUser(ctx *gin.Context)
}

type getUser struct {
	repository repository.UserRepository
}

func (u getUser) GetUser(ctx *gin.Context) {
	id, _ := ctx.Get("id")
	user := u.repository.FindById(id.(string))
	if user == nil {
		utils.NotFound(ctx, map[string]interface{}{"status": "failed", "message": "User not found"})
		return
	}
	utils.Ok(ctx, map[string]interface{}{"status": "success", "user": user})
	return
}
func NewGetUser(repository repository.UserRepository) GetUser {
	return &getUser{repository}
}
