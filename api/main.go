package main

import (
	"SNFOK/api/database"
	"SNFOK/api/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	database.InitDB()
	defer database.DB.Close()

	routes.AuditCheckRoutes(router)
	routes.ClientBeatRoutes(router)
	routes.InventoryRoutes(router)

	http.ListenAndServe(":6444", router)
}
