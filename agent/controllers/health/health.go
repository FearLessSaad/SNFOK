package health

import (
	"context"
	"fmt"
	"log"

	"github.com/FearLessSaad/SNFOK/agent/tooling/k8sclient"
	"github.com/FearLessSaad/SNFOK/shared/agent_dto"
	"github.com/gofiber/fiber/v2"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func HealthController(router fiber.Router) {
	router.Get("/beat", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{})
	})

	router.Get("/get", func(c *fiber.Ctx) error {
		// Get the singleton Kubernetes clientset
		clientset, err := k8sclient.GetClientset()
		if err != nil {
			log.Fatalf("Failed to get Kubernetes clientset: %v", err)
		}

		// Create a context for API calls
		ctx := context.Background()

		// Check 1: Verify API server is responsive by listing namespaces
		fmt.Println("Checking Kubernetes API server...")
		_, err = clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
		if err != nil {
			log.Fatalf("Kubernetes API server is not responding: %v", err)
		}

		// Check 2: Verify node status
		fmt.Println("\nChecking cluster nodes...")
		nodes, err := clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
		if err != nil {
			log.Fatalf("Failed to list nodes: %v", err)
		}

		// Check each node's Ready condition
		healthyNodes := 0
		for _, node := range nodes.Items {
			for _, condition := range node.Status.Conditions {
				if condition.Type == corev1.NodeReady && condition.Status == corev1.ConditionTrue {
					healthyNodes++
					break
				}
			}
		}

		k8s_status := true
		if healthyNodes == 0 {
			k8s_status = false
		}

		c_name, err := k8sclient.GetClusterName("")

		if err != nil {
			c_name = "empty"
		}

		return c.JSON(agent_dto.HealthResponse{
			K8sInfo: agent_dto.Health_K8s{
				ClusterName:      c_name,
				KubernetesStatus: k8s_status,
				HealthyNodes:     healthyNodes,
			},
			SystemInfo: agent_dto.Health_System{
				Status:  true,
				Message: "SNFOK Agent is running correctly.",
			},
		})
	})
}
