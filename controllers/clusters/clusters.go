package clusters

import "github.com/gofiber/fiber/v2"

func ClusterController(router fiber.Router) {
	ClusterInfo(router)
}
