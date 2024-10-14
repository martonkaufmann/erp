package customer

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/martonkaufmann/erp/http/response"
	"github.com/martonkaufmann/erp/model"
	"github.com/martonkaufmann/erp/provider"
	"gorm.io/gorm"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(provider.DatabaseKey).(*gorm.DB)
	l := r.Context().Value(provider.LogKey).(*slog.Logger)
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		l.Error("Failed to parse id", "error", err)

		w.WriteHeader(http.StatusNotFound)
		return
	}

	result := db.Where("id", id).Delete(&model.Customer{})

	if result.Error != nil {
		l.Error("Failed to delete customer", "error", result.Error)

        response.JSON(w, response.Error{Message: "Failed to delete customer"}, http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
