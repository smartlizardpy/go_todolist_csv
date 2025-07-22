# Go Todo List

A simple command-line todo list application written in Go.

## Description

This application allows you to manage a list of tasks from the command line. You can add new tasks, list existing tasks, and mark tasks as complete. The tasks are stored in a `tasks.csv` file in the same directory as the executable.

## Features

- Add a new task to the list.
- List all tasks in a formatted table.
- Mark a task as "Done".

## Usage

The application is controlled via command-line arguments.

### List Tasks

To see all the tasks in your todo list, use the `--list` command:

```bash
go run main.go --list
```

### Add a Task

To add a new task, use the `--add` command followed by the task description in quotes:

```bash
go run main.go --add "My new task"
```

### Mark a Task as Done

To mark a task as complete, use the `--mark` command followed by the task's ID number. The ID is the number shown in the first column when you list the tasks.

```bash
go run main.go --mark 0
```

## Dependencies

This project uses the following external Go module:

- `github.com/jedib0t/go-pretty/v6`: For creating and displaying the todo list in a visually appealing table format.

## Installation

To run this program, you need to have Go installed on your system.

1.  Clone the repository or download the source code.
2.  Open a terminal in the project directory.
3.  Install the dependencies:

    ```bash
    go mod tidy
    ```

4.  Run the application using the `go run` command with one of the supported flags.

## CSV File Format

The tasks are stored in a `tasks.csv` file with two columns: `Task` and `Status`.

- **Task**: The description of the task.
- **Status**: The current status of the task (e.g., "Not Done", "Done").

A new `tasks.csv` file will be created automatically if one does not exist when you add your first task.

---

Inspired by: https://roadmap.sh/projects/task-tracker
