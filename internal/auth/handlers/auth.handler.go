package auth

import (
	authDTO "secrets-golang/internal/auth/dto"
	adapters "secrets-golang/internal/infra"
	infra "secrets-golang/internal/infra"
	repository "secrets-golang/internal/user/repositories"

	"secrets-golang/internal/utils"

	"github.com/gin-gonic/gin"
)

type Auth interface {
	Auth(ctx *gin.Context)
}

type auth struct {
	repository repository.UserRepository
	token      infra.TokenInfra
}

func (u auth) Auth(ctx *gin.Context) {
	var payload authDTO.AuthPayloadDto
	ctx.BindJSON(&payload)
	user := u.repository.FindOne(payload.Email)
	if user == nil {
		utils.NotFound(ctx, map[string]interface{}{"status": "failed", "message": "User not found"})
		return
	}
	validatePassword := adapters.CheckPasswordHash(payload.Password, user.Password)
	if validatePassword == false {
		utils.Forbidden(ctx, map[string]interface{}{"status": "failed", "message": "Password is invalid"})
		return
	}
	signedToken, err := u.token.GenerateToken(user.Id.Hex())
	if err != nil {
		utils.InternalError(ctx)
		return
	}
	utils.Ok(ctx, map[string]interface{}{"status": "success", "token": signedToken})
	return
}
func NewAuth(repository repository.UserRepository, token adapters.TokenInfra) Auth {
	return &auth{repository, token}
}
