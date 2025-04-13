package routes

import (
	"SNFOK/api/controllers"

	"github.com/gorilla/mux"
)

func ClientBeatRoutes(router *mux.Router) {
	router.HandleFunc("/client/beat/info", controllers.ClientCheckBeatWithInfo)
}
