package routes

import (
	"SNFOK/api/controllers"

	"github.com/gorilla/mux"
)

func AuditCheckRoutes(router *mux.Router) {
	router.HandleFunc("/audit/check/success", controllers.AuditCheckSuccess)
	router.HandleFunc("/audit/get/all", controllers.GetAllAudits)
}
