package customer

import (
	"fmt"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Called")
}
