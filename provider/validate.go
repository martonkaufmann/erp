package provider

import (
	"context"

	validator "github.com/go-playground/validator/v10"
)

const ValidateKey = "validation"

func WithValidate(ctx context.Context) context.Context {
	validate := validator.New()

	return context.WithValue(ctx, ValidateKey, validate)
}
