package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jsam6/go-orders-api/utils"
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

func Create(w http.ResponseWriter, r *http.Request) {
	user := User{}
	fmt.Println(user)
	utils.ParseBody(r, user)
	m := user.CreateUser()
	res, _ := json.Marshal(m)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}