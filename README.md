# Task Tracker CLI

This is a command-line task tracker application built in Go. It allows users to manage their tasks by adding, updating, deleting, and listing them.

## Features

*   **Add a new task:** `go run main.go add "Your task description"`
*   **Update a task:** `go run main.go update <task_id> "New task description"`
*   **Delete a task:** `go run main.go delete <task_id>`
*   **List all tasks:** `go run main.go list`
*   **Mark a task as "in progress":** `go run main.go in-progress <task_id>`
*   **Mark a task as "done":** `go run main.go done <task_id>`

## Getting Started

1.  Clone the repository.
2.  Run `go mod tidy` to install dependencies.
3.  Use the `go run main.go` command with the available options to manage your tasks.

## Project Reference

This project is based on the [Task Tracker CLI project from roadmap.sh](https.roadmap.sh/projects/task-tracker).
