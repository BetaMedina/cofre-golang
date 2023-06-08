package passwords

import (
	password_repository "secrets-golang/internal/password/repositories"
	user_repository "secrets-golang/internal/user/repositories"
	"secrets-golang/internal/utils"

	"github.com/gin-gonic/gin"
)

type ListPassword interface {
	ListPassword(ctx *gin.Context)
}

type listPassword struct {
	password_repository password_repository.PasswordRepository
	user_repository     user_repository.UserRepository
}

func (u listPassword) ListPassword(ctx *gin.Context) {
	userId, _ := ctx.Get("id")
	user := u.user_repository.FindById(userId.(string))
	if user == nil {
		utils.NotFound(ctx, map[string]interface{}{"status": "failed", "message": "User not found"})
		return
	}
	rows, err := u.password_repository.List(userId.(string))
	if err != nil {
		utils.InternalError(ctx)
		return
	}
	if len(*rows) < 1 {
		utils.NotFound(ctx, map[string]interface{}{"status": "failed", "message": "Passwords not found"})
		return
	}
	utils.Ok(ctx, map[string]interface{}{"status": "success", "passwords": rows})
	return
}
func NewListPassword(repository password_repository.PasswordRepository, user_repository user_repository.UserRepository) ListPassword {
	return &listPassword{
		repository,
		user_repository,
	}
}
