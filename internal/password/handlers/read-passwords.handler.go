package passwords

import (
	password_repository "secrets-golang/internal/password/repositories"
	"secrets-golang/internal/password/services"
	user_repository "secrets-golang/internal/user/repositories"
	"secrets-golang/internal/utils"

	"github.com/gin-gonic/gin"
)

type ReadPassword interface {
	ReadPassword(ctx *gin.Context)
}

type readPassword struct {
	password_repository password_repository.PasswordRepository
	user_repository     user_repository.UserRepository
}

func (u readPassword) ReadPassword(ctx *gin.Context) {
	userId, _ := ctx.Get("id")
	id := ctx.Param("passwordId")
	user := u.user_repository.FindById(userId.(string))
	if user == nil {
		utils.NotFound(ctx, map[string]interface{}{"status": "failed", "message": "User not found"})
		return
	}
	row, err := u.password_repository.Read(id, userId.(string))
	if err != nil {
		utils.InternalError(ctx)
		return
	}
	decryptedPass, err := services.DecryptPass([]byte(row.HashedKey), row.Pass)
	if err != nil {
		utils.InternalError(ctx)
		return
	}
	row.Pass = decryptedPass
	utils.Ok(ctx, map[string]interface{}{"status": "success", "passwords": row})
	return
}
func NewReadPassword(repository password_repository.PasswordRepository, user_repository user_repository.UserRepository) ReadPassword {
	return &readPassword{
		repository,
		user_repository,
	}
}
