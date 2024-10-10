package middleware

import (
	"log/slog"
	"net/http"

	"github.com/martonkaufmann/erp/provider"
)

func RequestLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := r.Context().Value(provider.LogKey).(*slog.Logger)

		l.Info("Request", "method", r.Method, "path", r.URL.Path)

		next.ServeHTTP(w, r)
	})
}
