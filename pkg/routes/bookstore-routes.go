package routes

import (
	"github.com/apiorno/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterBookStoreRoutes(router *mux.Router) {
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/books/{id}", controllers.GetBooks).Methods("GET")
}
