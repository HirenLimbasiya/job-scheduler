# Job Scheduler Application

This project is a **Job Scheduler** application that allows users to submit, view, and track the status of jobs. The backend is built with **Go**, while the frontend is built with **React**. The application uses WebSockets for real-time updates and provides job status tracking for each job's progress.

## Description

The Job Scheduler consists of two main components:

- **Backend**: Built with **Go** and **Fiber** for API routes and WebSocket communication.
- **Frontend**: Built with **React** for a user-friendly interface to submit jobs and view their statuses in real-time.

Key features:
- Submit a job with a name and duration.
- View all jobs and their current status (pending, running, completed).
- Real-time updates of job statuses via WebSocket.

## Technologies Used

### Backend
- **Go**: A statically typed, compiled programming language for building the backend.
- **Fiber**: A fast web framework for Go.
- **WebSocket**: For real-time updates of job statuses.

### Frontend
- **React**: JavaScript library for building user interfaces.

## Installation

```bash
git clone https://github.com/HirenLimbasiya/job-scheduler.git
```

## Backend Setup

### Step 1: Navigate to the backend folder:

```bash
cd backend
```
### Step 2: Install required Go packages:

```bash
go mod tidy
```
### Step 3: Run the Go backend:

```bash
go run main.go
```
## Frontend Setup

### Step 1: Navigate to the frontend folder:

```bash
cd frontend
```
### Step 2: Install dependencies:

```bash
npm install
```
### Step 3: Run the frontend:

```bash
npm start

```
## How to Use

1. **Create a Job**:
   - Navigate to the form on the homepage.
   - Fill out the job name and duration, then submit the form.
   - The job will be added to the job list, and its status will initially be **pending**.

2. **View Job List**:
   - The job list will display all jobs with their status (pending, running, or completed).
   - Job status will be updated in real-time via WebSockets.

3. **Real-time Updates**:
   - As jobs progress, the status will be updated on the frontend without the need to refresh the page.

4. **Job Status**:
   - **Pending**: Jobs that are waiting to be processed.
   - **Running**: Jobs that are currently being processed.
   - **Completed**: Jobs that have finished processing.
## Design Choices

1. **Backend (Go)**:
   - The backend is built with **Go** to provide a high-performance, concurrent server capable of managing real-time job status updates via WebSockets.
   - **Fiber** is used for the web framework, chosen for its lightweight and fast performance.
   - **WebSocket** is used to push job status updates to the frontend in real-time, allowing users to view job progress without refreshing the page.
   - The **SJF (Shortest Job First)** algorithm is implemented to handle job scheduling, ensuring that the shortest jobs are prioritized and completed first.

2. **Frontend (React)**:
   - The frontend is built using **React** for component-based architecture, making the UI modular and easy to maintain.
   - **CSS** and **Flexbox** are used for responsive layout design to ensure the app is usable on both large and small screens.
   - Real-time updates are handled using **WebSocket** to listen for changes to job statuses and update the UI instantly.
   - **State management** is handled using React's `useState` and `useEffect` hooks to manage and update job data in the UI.

3. **Job Scheduler**:
   - Jobs are processed in a queue, with their statuses updated as they progress through various stages: **pending**, **running**, and **completed**.
   - The system is designed to handle multiple jobs concurrently, using Go's goroutines and channels for job management.


## Author

This project is created and maintained by **Hiren Limbasiya**.
You can explore more of my work on my [Portfolio](https://www.hirenlimbasiya.com/).

