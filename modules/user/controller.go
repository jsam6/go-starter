package user

import (
	"encoding/json"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	newMembers := GetAllUsers()
	res, _ := json.Marshal(newMembers)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}