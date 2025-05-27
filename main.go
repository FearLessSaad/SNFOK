package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"

	"github.com/FearLessSaad/SNFOK/controllers/auth"
	"github.com/FearLessSaad/SNFOK/controllers/clusters"
	"github.com/FearLessSaad/SNFOK/controllers/kubernetes"
	"github.com/FearLessSaad/SNFOK/controllers/policies"
	"github.com/FearLessSaad/SNFOK/db/initializer"
	"github.com/FearLessSaad/SNFOK/middlewares"
	"github.com/FearLessSaad/SNFOK/tooling"
	"github.com/FearLessSaad/SNFOK/tooling/logger"
)

func main() {
	logger.Log("", "Application is starting.")

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		AppName:      "HashX SNFOK API",
		ServerHeader: "HashX SNFOK",
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})

	// Do All Other Application Related Code Below
	initializer.InitializeDatabase()

	// Encrypt Cookies
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: "eqnVqTihpmg5ico1TCccc2JrvHyWbbpHiuVlOi/5Gp4=",
	}))

	// Use Helmet Middleare
	app.Use(helmet.New())

	// CORS Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Content-Type, Authorization, Accept, Origin",
		AllowCredentials: true, // Enable credentials for cookies
		ExposeHeaders:    "Content-Type, Authorization",
		MaxAge:           86400,
	}))

	// Add Request ID  Middleware
	app.Use(requestid.New(requestid.Config{
		Header: "X-Request-ID",
		Generator: func() string {
			return uuid.New().String()
		},
	}))

	// Custom logger middleware
	app.Use(middlewares.LoggingMiddleware)

	api := "/api/v1"
	auth.AuthController(app.Group(api + "/auth"))

	app.Use(middlewares.AuthMiddleware)
	clusters.ClusterController(app.Group(api + "/clusters"))
	kubernetes.KubernetesController(app.Group(api + "/kubernetes"))
	policies.PoliciesController(app.Group(api + "/policies"))
	// -----------------------------------------------

	// Channel to receive OS signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Handle panics with defer
	defer func() {
		if r := recover(); r != nil {
			tooling.CleanupPanic(app, fmt.Sprintf("Panic: %v", r))
			app.Shutdown()
			//os.Exit(1)
		}
	}()
	// Goroutine to handle signals
	go func() {
		sig := <-sigChan
		tooling.CleanupShutdown(app, fmt.Sprintf("Received signal %s", sig))
		app.Shutdown()
		//os.Exit(0)
	}()

	// Start Fiber server
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "4000" // Default port if not set
	}
	listenAddr := fmt.Sprintf("0.0.0.0:%s", port)

	if err := app.Listen(listenAddr); err != nil {
		tooling.CleanupShutdown(app, fmt.Sprintf("server error: %v", err))
		os.Exit(1)
	}
}
