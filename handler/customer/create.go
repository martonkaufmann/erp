package customer

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/martonkaufmann/erp/model"
	"github.com/martonkaufmann/erp/provider"
	"gorm.io/gorm"
)

func Create(w http.ResponseWriter, r *http.Request) {
	request := struct {
        FirstName string `json:"first_name" validate:"required,min=2,max=64"`
		LastName  string `json:"last_name" validate:"required,min=2,max=64"`
        Email     string `json:"email" validate:"required,email"`
    }{}

	db := r.Context().Value(provider.DatabaseKey).(*gorm.DB)
	l := r.Context().Value(provider.LogKey).(*slog.Logger)
    v := r.Context().Value(provider.ValidateKey).(*validator.Validate)

    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        l.Error("Failed to decode payload", err)

		http.Error(w, "Request failed", http.StatusBadRequest)
		return
    }

    if err := v.Struct(request); err != nil {
		http.Error(w, "Request validation failed", http.StatusBadRequest)
		return
    }

    user := model.Customer{FirstName: request.FirstName, LastName: request.LastName}
	result := db.Create(&user)

	if result.Error != nil {
		l.Error("Failed to create customer", result.Error)

		http.Error(w, "Request failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	if 	_, err := fmt.Fprintln(w); err != nil {
		l.Error("Failed to send response", err)
	}
}
