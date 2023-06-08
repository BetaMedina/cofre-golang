package services

import (
	"secrets-golang/internal/infra"
	userDTO "secrets-golang/internal/user/dto"
)

func CryptPass(user *userDTO.UserPayloadDto) error {
	password, err := infra.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = password
	return nil
}
