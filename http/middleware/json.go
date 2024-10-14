package middleware

import (
	"net/http"
	"slices"
	"strings"
)

func JSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := strings.ToLower(r.Header.Get("Content-Type"))
		accept := strings.ToLower(r.Header.Get("Accept"))
		methods := []string{"POST", "PUT", "PATCH", "DELETE"}

		if (contentType != "application/json" || accept != "application/json") && slices.Contains(methods, r.Method) {
            w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
