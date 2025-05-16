package middlewares

import (
	"github.com/FearLessSaad/SNFOK/tooling/logger"

	"github.com/gofiber/fiber/v2"
)

func LoggingMiddleware(c *fiber.Ctx) error {
	// Process request
	err := c.Next()

	// Get request details
	status := c.Response().StatusCode()
	method := c.Method()
	path := c.Path()
	requestID := c.GetRespHeader("X-Request-ID")

	// Log request details
	logger.Log(logger.INFO, "REQUEST",
		logger.Field{Key: "status", Value: status},
		logger.Field{Key: "method", Value: method},
		logger.Field{Key: "path", Value: path},
		logger.Field{Key: "request_id", Value: requestID},
	)

	return err
}
