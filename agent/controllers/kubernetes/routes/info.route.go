package routes

import (
	"fmt"
	"log"
	"os"

	"github.com/FearLessSaad/SNFOK/agent/controllers/kubernetes/features"
	"github.com/FearLessSaad/SNFOK/agent/tooling/k8sclient"
	"github.com/gofiber/fiber/v2"
)

func GetAllWorkerNodes(router fiber.Router) {

	router.Get("/workers/nodes/all", func(c *fiber.Ctx) error {

		// Get the singleton Kubernetes clientset
		clientset, err := k8sclient.GetClientset()
		if err != nil {
			log.Fatalf("Failed to get Kubernetes clientset: %v", err)
		}

		nodes, err := features.GetWorkerNodes(clientset)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting worker nodes: %v\n", err)
			os.Exit(1)
		}

		return c.Status(fiber.StatusOK).JSON(nodes)
	})
}
