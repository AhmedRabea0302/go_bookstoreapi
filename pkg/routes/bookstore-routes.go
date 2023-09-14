package routes

import (
	"github.com/AhmedRabea0302/go_bookstoreapi/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.updateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controllers.deleteBook).Methods("DELETE")
}
