package models

import (
	"github.com/go-playground/validator/v10"

	internalValidator "github.com/Snehashish1609/validator-api/internal/validators"
)

type User struct {
	Name   string `json:"name" validate:"required"`
	PAN    string `json:"pan" validate:"required,pan"`
	Mobile string `json:"mobile" validate:"required,mobile"`
	Email  string `json:"email" validate:"required,email"`
}

type UserHandler struct {
	Validator *validator.Validate
}

func NewUserHandler() *UserHandler {
	v := validator.New()
	internalValidator.RegisterCustomValidators(v)
	return &UserHandler{Validator: v}
}
