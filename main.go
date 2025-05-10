package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// User struct
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// yaratilgan userlar
var users = []User{
	{ID: 1, Name: "Azizbek", Email: "azizdev.full@gmail.com"},
	{ID: 2, Name: "John Doe", Email: "john.doe@example.com"},
}

// userlarni olish (GET)
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// user yaratish (POST)
func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User

	contentType := r.Header.Get("Content-Type")

	if contentType == "application/json" {
		// Agar JSON bo'lsa
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			http.Error(w, "JSON noto'g'ri", http.StatusBadRequest)
			return
		}
	} else {
		// Boshqa formatlar uchun (form-data yoki x-www-form-urlencoded)
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			err = r.ParseForm() // fallback
			if err != nil {
				http.Error(w, "Formani o'qib bo'lmadi", http.StatusBadRequest)
				return
			}
		}

		newUser.Name = r.FormValue("name")
		newUser.Email = r.FormValue("email")
	}

	if newUser.Name == "" || newUser.Email == "" {
		http.Error(w, "Name va Email to'ldirilishi shart", http.StatusBadRequest)
		return
	}

	newUser.ID = len(users) + 1
	users = append(users, newUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

// Foydalanuvchi ma'lumotlarini olish (GET by ID)
func getUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, user := range users {
		if user.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.NotFound(w, r)
}

func main() {
	http.HandleFunc("/users", getUsers)     // Foydalanuvchilar ro'yxatini olish
	http.HandleFunc("/users/", getUserByID) // Foydalanuvchi ID bilan olish
	http.HandleFunc("/create", createUser)  // Yangi foydalanuvchi yaratish

	fmt.Println("Server ishga tushdi: http://localhost:8080")
	http.ListenAndServe(":8080", nil) // 8080 portda serverni ishga tushuramiz
}
