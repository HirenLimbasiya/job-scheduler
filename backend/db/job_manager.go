package db

import (
	"fmt"
	"sync"
	"time"
	"job-scheduler-backend/types"
	"job-scheduler-backend/broadcast"
)

var (
	jobQueue []types.Job
	jobMutex sync.Mutex
)

// AddJob adds a job to the queue and returns the job with a unique ID.
func AddJob(job types.Job) types.Job {
	jobMutex.Lock()
	defer jobMutex.Unlock()

	job.ID = len(jobQueue) + 1
	job.Status = "pending"
	job.Duration = job.Duration * time.Second
	jobQueue = append(jobQueue, job)

	return job
}

// ListJobs returns a copy of the job queue.
func ListJobs() []types.Job {
	jobMutex.Lock()
	defer jobMutex.Unlock()

	queueCopy := make([]types.Job, len(jobQueue))
	copy(queueCopy, jobQueue)

	return queueCopy
}

// RunScheduler starts the job scheduler in the background.
func RunScheduler() {
	go func() {
		for {
			time.Sleep(1 * time.Second)
			jobMutex.Lock()

			var runningJob *types.Job
			for i := range jobQueue {
				if jobQueue[i].Status == "running" {
					runningJob = &jobQueue[i]
					break
				}
			}

			if runningJob != nil {
				runningJob.RemainingTime -= time.Second
				if runningJob.RemainingTime <= 0 {
					runningJob.Status = "completed"
					runningJob.RemainingTime = 0
					fmt.Printf("Job '%s' has completed!\n", runningJob.Name)
				}
				broadcast.UpdateClients(jobQueue)
			} else {
				var shortestJob *types.Job
				for i := range jobQueue {
					if jobQueue[i].Status == "pending" {
						if shortestJob == nil || jobQueue[i].Duration < shortestJob.Duration {
							shortestJob = &jobQueue[i]
						}
					}
				}
				if shortestJob != nil {
					shortestJob.Status = "running"
					shortestJob.RemainingTime = shortestJob.Duration
					broadcast.UpdateClients(jobQueue)
				}
			}

			jobMutex.Unlock()
		}
	}()
}
