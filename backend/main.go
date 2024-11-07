package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
	"job-scheduler-backend/api"
	"job-scheduler-backend/db"
)

func main() {
	app := fiber.New()

	// Enable CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Set up routes
	app.Post("/jobs", api.CreateJob)
	app.Get("/jobs", api.ListJobs)
	app.Get("/ws", websocket.New(api.HandleWebSocket))

	// Start job scheduler
	db.RunScheduler()

	// Start the server
	if err := app.Listen(":2027"); err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}
}
