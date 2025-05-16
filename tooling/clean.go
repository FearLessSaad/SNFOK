package tooling

import (
	"github.com/FearLessSaad/SNFOK/db"
	"github.com/FearLessSaad/SNFOK/tooling/logger"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CleanupPanic(app *fiber.App, reason string) {
	logger.Log("error", fmt.Sprintf("Application crashed due to panic: %s", reason))
	db.Close()
	db.CloseAll()
	if err := app.Shutdown(); err != nil {
		logger.Log("error", "Failed to shutdown Fiber app after panic", logger.Field{Key: "error", Value: err.Error()})
	}
}

func CleanupShutdown(app *fiber.App, reason string) {
	logger.Log("", fmt.Sprintf("Application is shutting down gracefully: %s", reason))
	db.CloseAll()
	db.Close()
	if err := app.Shutdown(); err != nil {
		logger.Log("error", "Failed to shutdown Fiber app during graceful shutdown", logger.Field{Key: "error", Value: err.Error()})
	}
}
