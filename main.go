package main

import "fmt"

type AuthError struct {
	Message string
}

func (e AuthError) Error() string {
	return e.Message
}

func login(email string, password string) (string, error) {
	myEmail := "azizdev.full@gmail.com"
	myPassword := "123456"
	if myEmail != email || myPassword != password {
		return "", AuthError{Message: "Email yoki parol xato"}
	} else {
		return "Siz muvaffaqiyatli ro'yxatdan o'tdingiz", nil
	}
}

func main() {
	user, err := login("azizdev.full@gmail.com", "123456")
	if err != nil {
		fmt.Println("Xato: ", err)
	} else {
		fmt.Println("Foydalanuvchi: ", user)
	}
}
