package user

import (
	"encoding/json"
	"net/http"
    "github.com/go-chi/chi/v5"
)

func List(w http.ResponseWriter, r *http.Request) {
	newMembers := GetAllUsers()
	res, _ := json.Marshal(newMembers)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userDetail, _ := GetUserById(id)
	res, _ := json.Marshal(userDetail)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}