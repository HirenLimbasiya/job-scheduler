package api

import (
	"github.com/gofiber/fiber/v2"
	"job-scheduler-backend/db"
	"job-scheduler-backend/types"
	"job-scheduler-backend/broadcast"
)

// CreateJob handles job creation.
func CreateJob(c *fiber.Ctx) error {
	var job types.Job
	if err := c.BodyParser(&job); err != nil {
		return c.Status(400).SendString("Invalid input")
	}

	job = db.AddJob(job)
	broadcast.UpdateClients(db.ListJobs())

	return c.Status(201).JSON(job)
}

// ListJobs handles listing all jobs.
func ListJobs(c *fiber.Ctx) error {
	return c.JSON(db.ListJobs())
}
