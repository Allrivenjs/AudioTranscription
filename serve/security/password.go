package security

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	fmt.Println(string(hashed))
	return string(hashed), nil
}

func VerifyPassword(hashed, password string) error {
	fmt.Println("hashed: ", hashed)
	fmt.Println("password: ", password)
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}
