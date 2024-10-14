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

func Restore(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(provider.DatabaseKey).(*gorm.DB)
	l := r.Context().Value(provider.LogKey).(*slog.Logger)
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		l.Error("Failed to parse id", "error", err)

		w.WriteHeader(http.StatusNotFound)
		return
	}

	result := db.Model(&model.Customer{}).
		Unscoped().
		Where("id = ?", id).
		Update("deleted_at", nil)

	if result.Error != nil {
		l.Error("Failed to restore customer", "error", result.Error)

        response.JSON(w, response.Error{Message: "Request failed"}, http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
