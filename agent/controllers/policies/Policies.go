package policies

import (
	"github.com/FearLessSaad/SNFOK/agent/controllers/policies/routes"
	"github.com/gofiber/fiber/v2"
)

func PoliciesController(router fiber.Router) {
	routes.DeployPolicy(router)
	routes.PodIsolation(router)
}
