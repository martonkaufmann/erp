package middleware

import (
	"net/http"
	"strings"
)

func JSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        contentType := strings.ToLower(r.Header.Get("Content-Type"))

        if contentType != "application/json" {
            http.Error(w, "Unsupported Media Type", http.StatusUnsupportedMediaType)
            return
        }

		next.ServeHTTP(w, r)
	})
}
