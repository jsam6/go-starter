package main

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jsam6/go-orders-api/modules/user"
	"github.com/jsam6/go-orders-api/config"
)

func main() {
	config.Connect()

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/player", user.List )
	router.Get("/player/{id}", user.Get )

	http.ListenAndServe(":3000", router)
}
