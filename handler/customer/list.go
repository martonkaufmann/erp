package customer

import (
	"github.com/go-playground/form/v4"
	"github.com/go-playground/validator/v10"
	"github.com/martonkaufmann/erp/http/response"
	"github.com/martonkaufmann/erp/model"
	"github.com/martonkaufmann/erp/provider"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	resp := struct {
		Customers []model.Customer `json:"customers"`
	}{}
	request := struct {
		Sort      string            `form:"sort" validate:"omitempty,oneof=first_name last_name email"`
		Direction string            `form:"direction" validate:"omitempty,oneof=asc desc"`
		Filter    map[string]string `form:"filter" validate:"omitempty,dive,keys,oneof=first_name last_name email,endkeys"`
	}{}
	db := r.Context().Value(provider.DatabaseKey).(*gorm.DB)
	l := r.Context().Value(provider.LogKey).(*slog.Logger)
	v := r.Context().Value(provider.ValidateKey).(*validator.Validate)

	if err := form.NewDecoder().Decode(&request, r.URL.Query()); err != nil {
		l.Error("Failed to decode payload", "error", err)

		response.JSON(w, response.Error{Message: "Request failed"}, http.StatusBadRequest)
		return
	}

	if err := v.Struct(request); err != nil {
		l.Debug("Failed to validate payload", "error", err)

        response.JSON(w, response.Error{Message: "Request validation failed"}, http.StatusBadRequest)
		return
	}

	if request.Direction == "" {
		request.Direction = "asc"
	}

	if request.Sort == "" {
		request.Sort = "id"
	}

	db = db.Order(request.Sort + " " + request.Direction)

	for key, value := range request.Filter {
		db = db.Where(key+" = ?", value)
	}

	result := db.Find(&resp.Customers)

	if result.Error != nil {
		l.Error("Error while fetching customers", "error", result.Error)

        w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.JSON(w, resp, http.StatusOK)
}
