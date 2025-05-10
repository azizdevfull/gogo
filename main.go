package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	password := "123456"

	hash, err := hashPassword(password)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Hashed:", hash)

	match := checkPasswordHash(password, hash)
	fmt.Println("Tog'ri parol:", match)

	match = checkPasswordHash("1234567", hash)
	fmt.Println("Xato parol:", match)
}
