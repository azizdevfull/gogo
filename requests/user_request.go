package requests

import "github.com/go-playground/validator/v10"

type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
	Age  int    `json:"age" validate:"required,min=18"`
}

type UpdateUserRequest struct {
	Name string `json:"name" validate:"min=3,max=100"`
	Age  int    `json:"age" validate:"min=18"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateCreateUserRequest(req CreateUserRequest) error {
	return validate.Struct(req)
}

func ValidateUpdateUserRequest(req UpdateUserRequest) error {
	return validate.Struct(req)
}
