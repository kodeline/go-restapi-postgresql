package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kodeline/go-restapi-postgresql/routes"
)

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", route)
}
