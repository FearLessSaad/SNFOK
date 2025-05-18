package kubernetes

import (
	"github.com/FearLessSaad/SNFOK/agent/controllers/kubernetes/routes"
	"github.com/gofiber/fiber/v2"
)

func KubernetesController(router fiber.Router) {
	routes.GetAllWorkerNodes(router)
}
