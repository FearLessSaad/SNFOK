package main

import (
	"encoding/json"

	"github.com/FearLessSaad/SNFOK/agent/controllers/health"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func main() {

	app := fiber.New(fiber.Config{
		AppName:      "HashX SNFOK AGENT API",
		ServerHeader: "HashX SNFOK AGENT",
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})

	// Use Helmet Middleare
	app.Use(helmet.New())

	// CORS Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3989",
		AllowMethods: "*",
		AllowHeaders: "*",
		MaxAge:       86400,
	}))

	api := app.Group("/api")
	health.HealthController(api.Group("/health"))

	app.Listen("0.0.0.0:8990")
}
