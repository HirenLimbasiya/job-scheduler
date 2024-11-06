package main

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
)

type Job struct {
    ID            int           `json:"id"`
    Name          string        `json:"name"`
    Duration      time.Duration `json:"duration"` // Duration in seconds
    Status        string        `json:"status"`   // "pending", "running", "completed"
    RemainingTime  time.Duration `json:"remaining_time"` // Remaining time in seconds
}


var (
    jobQueue []Job
    jobMutex sync.Mutex
    clients  = make(map[*websocket.Conn]bool)
    wsMutex  sync.Mutex
)

func main() {
    app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE, PATCH",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization", // Include Authorization header
	}))

    app.Post("/jobs", createJob)
    app.Get("/jobs", listJobs)

    app.Get("/ws", websocket.New(handleWebSocket))

    go runJobScheduler()

    // app.Listen(":5000")
if err := app.Listen(":2027"); err != nil {
    fmt.Println("Failed to start server:", err)
    return
}

}

func createJob(c *fiber.Ctx) error {
    var job Job
    if err := c.BodyParser(&job); err != nil {
        return c.Status(400).SendString("Invalid input")
    }

    jobMutex.Lock()
    job.ID = len(jobQueue) + 1
    job.Status = "pending"
    jobQueue = append(jobQueue, job)
    jobMutex.Unlock()

    scheduleJobs()
    broadcastJobStatus()

    return c.Status(201).JSON(job)
}

func listJobs(c *fiber.Ctx) error {
    jobMutex.Lock()
    defer jobMutex.Unlock()
    return c.JSON(jobQueue)
}

func handleWebSocket(c *websocket.Conn) {
    wsMutex.Lock()
    clients[c] = true
    wsMutex.Unlock()

    defer func() {
        wsMutex.Lock()
        delete(clients, c)
        wsMutex.Unlock()
        c.Close()
    }()

    for {
        _, _, err := c.ReadMessage()
        if err != nil {
            break
        }
    }
}

func broadcastJobStatus() {
    wsMutex.Lock()
    defer wsMutex.Unlock()

    for client := range clients {
        client.WriteJSON(jobQueue)
    }
}

func scheduleJobs() {
    jobMutex.Lock()
    defer jobMutex.Unlock()

    // Sort jobs by shortest duration first (SJF) but keep completed jobs at the end
    sort.SliceStable(jobQueue, func(i, j int) bool {
        if jobQueue[i].Status == "completed" {
            return false
        }
        if jobQueue[j].Status == "completed" {
            return true
        }
        return jobQueue[i].Duration < jobQueue[j].Duration
    })
}

// func runJobScheduler() {
//     for {
//         time.Sleep(1 * time.Second)

//         jobMutex.Lock()
//         // Look for the first "pending" job to process
//         var currentJob *Job
//         for i := range jobQueue {
//             if jobQueue[i].Status == "pending" {
//                 currentJob = &jobQueue[i]
//                 break
//             }
//         }
//         jobMutex.Unlock()

//         // If no pending job found, continue to next iteration
//         if currentJob == nil {
//             continue
//         }

//         // Set the job as running
//         jobMutex.Lock()
//         currentJob.Status = "running"
//         jobMutex.Unlock()
//         broadcastJobStatus()

//         // Simulate job processing
//         time.Sleep(currentJob.Duration * time.Second)

//         // Mark job as completed
//         jobMutex.Lock()
//         currentJob.Status = "completed"
//         jobMutex.Unlock()
//         broadcastJobStatus()
//     }
// }

func runJobScheduler() {
    for {
        time.Sleep(1 * time.Second)

        var currentJob *Job

        jobMutex.Lock()
        // Look for the first "pending" job to process
        for i := range jobQueue {
            if jobQueue[i].Status == "pending" {
                currentJob = &jobQueue[i]
                break
            }
        }
        jobMutex.Unlock()

        // If no pending job found, continue to next iteration
        if currentJob == nil {
            continue
        }

        // Set the job as running and initialize remaining time
        jobMutex.Lock()
        currentJob.Status = "running"
        currentJob.RemainingTime = currentJob.Duration // Set initial remaining time
        jobMutex.Unlock()

        // Broadcast status before starting job processing
        broadcastJobStatus()

        // Simulate job processing
        for remaining := currentJob.Duration; remaining > 0; remaining-- {
            jobMutex.Lock()
            currentJob.RemainingTime = remaining
            jobMutex.Unlock()

            broadcastJobStatus() // Broadcast remaining time updates every second
            time.Sleep(1 * time.Second) // Simulate time passing
        }

        // Mark job as completed
        jobMutex.Lock()
        currentJob.Status = "completed"
        currentJob.RemainingTime = 0 // Reset remaining time
        jobMutex.Unlock()
        broadcastJobStatus()
    }
}
