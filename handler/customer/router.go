package customer

import (
	"net/http"
)

func RegisterRoutes(r *http.ServeMux) {
    r.HandleFunc("GET /customers", List)
    r.HandleFunc("POST /customers", Create)
    r.HandleFunc("DELETE /customers/{id}", Delete)
    r.HandleFunc("POST /customers/{id}/restore", Restore)
}
