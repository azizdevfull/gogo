package resources

import "myapp/models"

type UserResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewUserResponse(user models.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Age:       user.Age,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func NewUserListResponse(users []models.User) []UserResponse {
	var response []UserResponse
	for _, user := range users {
		response = append(response, NewUserResponse(user))
	}
	return response
}
