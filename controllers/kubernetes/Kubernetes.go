package kubernetes

import "github.com/gofiber/fiber/v2"

func KubernetesController(router fiber.Router) {
	KubernetesInfo(router)
}
