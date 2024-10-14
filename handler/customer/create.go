package customer

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/martonkaufmann/erp/http/response"
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

	matchingRecords := int64(0)
	db := r.Context().Value(provider.DatabaseKey).(*gorm.DB)
	l := r.Context().Value(provider.LogKey).(*slog.Logger)
	v := r.Context().Value(provider.ValidateKey).(*validator.Validate)

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		l.Error("Failed to decode payload", "error", err)

        response.JSON(w, response.Error{Message: "Request failed"}, http.StatusBadRequest)
		return
	}

	if err := v.Struct(request); err != nil {
		l.Debug("Failed to validate payload", "error", err)

        response.JSON(w, response.Error{Message: "Request validation failed"}, http.StatusBadRequest)
		return
	}

	db.Model(&model.Customer{}).Where("email = ?", request.Email).Count(&matchingRecords)

	if matchingRecords != 0 {
        response.JSON(w, response.Error{Message: "Email already exists"}, http.StatusBadRequest)
		return
	}

	user := model.Customer{FirstName: request.FirstName, LastName: request.LastName, Email: request.Email}
	result := db.Create(&user)

	if result.Error != nil {
		l.Error("Failed to create customer", "error", result.Error)

        response.JSON(w, response.Error{Message: "Failed to create customer"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
