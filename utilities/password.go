package utilities

import (
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
)

func HashPassword(password string) (string, error) {
	saltStr := os.Getenv("BCRYPT_SALT")
	salt, err := strconv.Atoi(saltStr)
	if err != nil {
		panic(err)
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(bytes), err
}

func CheckPassword(hashPassword string, plainPassword string) bool {
	hashPw := []byte(hashPassword)
	plainPw := []byte(plainPassword)
	if err := bcrypt.CompareHashAndPassword(hashPw, plainPw); err != nil {
		return false
	}
	return true
}
