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

	router.Get("/namespaces/all", func(c *fiber.Ctx) error {

		fmt.Println("Get All NameSpaces")
		// Get the singleton Kubernetes clientset
		clientset, err := k8sclient.GetClientset()
		if err != nil {
			log.Fatalf("Failed to get Kubernetes clientset: %v", err)
		}

		nodes, err := features.GetAllNamespaces(clientset)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting worker nodes: %v\n", err)
			os.Exit(1)
		}

		return c.Status(fiber.StatusOK).JSON(nodes)
	})

	router.Get("/count/pods", func(c *fiber.Ctx) error {

		fmt.Println("Get All Running Pods")
		// Get the singleton Kubernetes clientset
		clientset, err := k8sclient.GetClientset()
		if err != nil {
			log.Fatalf("Failed to get Kubernetes clientset: %v", err)
		}

		nodes, err := features.CountAllRunningPods(clientset)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting all running pods: %v\n", err)
			os.Exit(1)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"running_pods": nodes,
		})
	})

	router.Get("/namespaces/:namespace/resources", func(c *fiber.Ctx) error {

		// Get the singleton Kubernetes clientset
		clientset, err := k8sclient.GetClientset()
		if err != nil {
			log.Fatalf("Failed to get Kubernetes clientset: %v", err)
		}

		nodes, err := features.GetNamespaceResources(clientset, c.AllParams()["namespace"])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting worker nodes: %v\n", err)
			os.Exit(1)
		}

		return c.Status(fiber.StatusOK).JSON(nodes)
	})

	router.Get("/pod/describe/:namespace/:pd", func(c *fiber.Ctx) error {

		// Get the singleton Kubernetes clientset
		clientset, err := k8sclient.GetClientset()
		if err != nil {
			log.Fatalf("Failed to get Kubernetes clientset: %v", err)
		}

		nodes, err := features.GetNamespaceResources(clientset, c.AllParams()["namespace"])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting worker nodes: %v\n", err)
			os.Exit(1)
		}

		return c.Status(fiber.StatusOK).JSON(nodes)
	})
}
