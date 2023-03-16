package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/philipokiokio/book-ms-api/pkg/routes"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterBookStoreRouter(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:7070", r))
}
