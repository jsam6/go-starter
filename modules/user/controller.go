package user

import (
	"net/http"
    "encoding/json"
)

func List(w http.ResponseWriter, r *http.Request) {
    emptyList := []string{}
	json.NewEncoder(w).Encode(emptyList)
}
