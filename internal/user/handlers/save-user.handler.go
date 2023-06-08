package user

import (
	userDTO "secrets-golang/internal/user/dto"
	repository "secrets-golang/internal/user/repositories"
	"secrets-golang/internal/user/services"
	"secrets-golang/internal/utils"

	"github.com/gin-gonic/gin"
)

type SaveUser interface {
	SaveUser(ctx *gin.Context)
}

type saveUser struct {
	repository repository.UserRepository
}

func (u saveUser) SaveUser(ctx *gin.Context) {
	var payload userDTO.UserPayloadDto
	ctx.BindJSON(&payload)
	if user := u.repository.FindOne(payload.Email); user != nil {
		utils.BadRequest(ctx, map[string]interface{}{"status": "failed", "message": "User alredy exists"})
		return
	}
	err := services.CryptPass(&payload)
	if err != nil {
		utils.InternalError(ctx)
		return
	}
	result, err := u.repository.Save(&payload)
	if err != nil {
		utils.InternalError(ctx)
		return
	}
	utils.Ok(ctx, map[string]interface{}{"status": "success", "userId": result.InsertedID})
	return
}
func NewUser(repository repository.UserRepository) SaveUser {
	return &saveUser{repository}
}
