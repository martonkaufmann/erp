package customer

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/martonkaufmann/erp/model"
	"github.com/martonkaufmann/erp/provider"
	"gorm.io/gorm"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(provider.DatabaseKey).(*gorm.DB)
	l := r.Context().Value(provider.LogKey).(*slog.Logger)
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		l.Error("Failed to parse id", err)

		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	result := db.Delete(model.Customer{
		Model: model.Model{
			ID: model.ModelId(id),
		},
	})

	if result.Error != nil {
		l.Error("Failed to delete customer", result.Error)

		http.Error(w, "Request failed", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
