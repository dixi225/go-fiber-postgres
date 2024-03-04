package routes

import (
	"github.com/dixi225/go-fiber-postgres/pkg/controllers"
	"github.com/gorilla/mux"
)

var BookRouter = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.createBook).Methods("POST")
	router.HandleFunc("/book", controllers.getBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.getBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.updateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.deleteBook).Methods("DELETE")
}
