package types

import "time"

// Job represents a job in the queue.
type Job struct {
	ID            int           `json:"id"`
	Name          string        `json:"name"`
	Duration      time.Duration `json:"duration"`       // Duration in seconds
	Status        string        `json:"status"`         // "pending", "running", "completed"
	RemainingTime time.Duration `json:"remaining_time"` // Remaining time in seconds
}
