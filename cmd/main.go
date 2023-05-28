package main

import (
	"bookapi/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	fmt.Println("📌 listen and server on port [::]8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
