package passwords

import (
	passwordDTO "secrets-golang/internal/password/dto"
	password_repository "secrets-golang/internal/password/repositories"
	"secrets-golang/internal/password/services"
	user_repository "secrets-golang/internal/user/repositories"
	"secrets-golang/internal/utils"

	"github.com/gin-gonic/gin"
)

type SavePassword interface {
	SavePassword(ctx *gin.Context)
}

type savePassword struct {
	password_repository password_repository.PasswordRepository
	user_repository     user_repository.UserRepository
}

func (u savePassword) SavePassword(ctx *gin.Context) {
	var payload passwordDTO.SavePasswordPayloadDto
	ctx.BindJSON(&payload)
	id, _ := ctx.Get("id")
	user := u.user_repository.FindById(id.(string))
	if user == nil {
		utils.NotFound(ctx, map[string]interface{}{"status": "failed", "message": "User not found"})
		return
	}
	randomPass := services.GenerateRandomKey(16)
	payload.Pass, _ = services.EncryptPass(payload.Pass, []byte(randomPass))
	insertedId, err := u.password_repository.Save(&passwordDTO.SavePasswordRepositoryDto{
		Description: payload.Description,
		UserId:      id.(string),
		Name:        payload.Name,
		Platform:    payload.Platform,
		Pass:        payload.Pass,
		HashedKey:   randomPass,
	})
	if err != nil {
		utils.InternalError(ctx)
		return
	}
	utils.Ok(ctx, map[string]interface{}{"status": "success", "passwordId": insertedId.InsertedID})
	return
}
func NewSavePassword(repository password_repository.PasswordRepository, user_repository user_repository.UserRepository) SavePassword {
	return &savePassword{
		repository,
		user_repository,
	}
}
