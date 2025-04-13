package main

import (
	"SNFOK/client/routes"
	"SNFOK/client/tooling"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	tooling.SendInitialBeatWithInfo()

	routes.HealthCheckRoutes(router)
	http.ListenAndServe(":45667", router)
}
