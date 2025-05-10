package main

import "fmt"

type AuthError struct {
	Message string
}

func (e AuthError) Error() string {
	return e.Message
}

type User struct {
	Name     string
	Email    string
	Password string
}

// In memory userlar bazasi
var users = []User{
	{Name: "Azizdev", Email: "azizdev.full@gmail.com", Password: "123456"},
	{Name: "YoungDev", Email: "young@dev.com", Password: "abcdef"},
}

func login(email, password string) (User, error) {
	for _, user := range users {
		if user.Email == email && user.Password == password {
			return user, nil
		}
	}
	return User{}, AuthError{Message: "Email yoki parol noto‘g‘ri"}
}

func main() {
	user, err := login("azizdev.full@gmail.com", "123456")
	if err != nil {
		fmt.Println("Xato:", err)
		return
	}
	fmt.Println("Xush kelibsiz", user.Name)
}
