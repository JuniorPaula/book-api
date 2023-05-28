package routes

import (
	"bookapi/pkg/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", handlers.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/book/", handlers.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/book/{bookId}", handlers.GetBookById).Methods(http.MethodGet)
	router.HandleFunc("/book/{bookId}", handlers.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/book/{bookId}", handlers.DeleteBook).Methods(http.MethodDelete)
}
