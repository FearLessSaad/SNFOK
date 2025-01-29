package routes

import (
	"SNFOK/api/controllers"

	"github.com/gorilla/mux"
)

func InventoryRoutes(router *mux.Router) {
	router.HandleFunc("/inventory/hosts/all", controllers.GetAllHosts).Methods("GET")
	router.HandleFunc("/inventory/hosts/live/all", controllers.GetAllLiveHosts).Methods("GET")
}
