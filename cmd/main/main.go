package main

import (
	"log"
	"net/http"

	"github.com/apiorno/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":9010", router))
}
