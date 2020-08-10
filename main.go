package main

import (
	"fmt"
	"log"
	"net/http"

	"webapi/routes"

	"github.com/gorilla/mux"
	// "github.com/rs/cors" // for testing
)

func main() {
	r := mux.NewRouter()

	routes.RouteHandler(r)

	fmt.Println("Starting Server")

	// handler := cors.Default().Handler(r)             // testing
	log.Fatal(http.ListenAndServe(":6969", r)) // log.Fatal(http.ListenAndServe(":6969", r))
}
