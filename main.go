package main

import (
	"fmt"
	"log"
	"net/http"

	"webapi/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes.RouteHandler(r)

	fmt.Println("Starting Server")

	log.Fatal(http.ListenAndServe(":6969", r))
}
