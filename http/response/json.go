package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, data any, statusCode int) error {
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(data)
}
