package main

import (
	"SNFOK/api/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.AuditCheckRoutes(router)

	http.ListenAndServe(":6444", router)
}
