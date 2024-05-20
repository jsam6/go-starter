package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jsam6/go-orders-api/config"
	"github.com/jsam6/go-orders-api/modules/user"
)

func main() {
	
	helloWorld := "helloworld"
	var pointerToHelloWorld *string
	pointerToHelloWorld = &helloWorld
	fmt.Println(helloWorld)
	fmt.Println(&helloWorld)
	fmt.Println()
	fmt.Println(pointerToHelloWorld)
	fmt.Println(&pointerToHelloWorld)
	fmt.Println(*pointerToHelloWorld)
	fmt.Println()
	fmt.Println(*&pointerToHelloWorld)
	fmt.Println(&*pointerToHelloWorld)

	config.Connect()

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Route("/player", user.PlayerRoutes)

	http.ListenAndServe(":3000", router)
}
