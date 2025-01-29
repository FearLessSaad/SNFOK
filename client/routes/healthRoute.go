package routes

import (
	"SNFOK/client/controllers"

	"github.com/gorilla/mux"
)

func HealthCheckRoutes(router *mux.Router) {
	router.HandleFunc("/health/check", controllers.HealthCheckController)
	router.HandleFunc("/health/check/info", controllers.HealthCheckAndInfoCOntroller)
}
