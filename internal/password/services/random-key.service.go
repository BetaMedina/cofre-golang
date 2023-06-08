package services

import (
	"math/rand"
	"secrets-golang/internal/infra"
)

var letters = []rune(infra.GetEnvs("HASHED_CHARACTERS"))

func GenerateRandomKey(length int) string {
	mountedString := make([]rune, length)
	for i := range mountedString {
		mountedString[i] = letters[rand.Intn(len(letters))]
	}
	return string(mountedString)
}
