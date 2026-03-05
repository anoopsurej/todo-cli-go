# tasks

A simple CLI task manager written in Go.

## Usage

```
tasks add <description>       # Create a new task
tasks list                    # List incomplete tasks
tasks list -a / --all         # List all tasks including completed
tasks complete <id>           # Mark a task as done
tasks delete <id>             # Delete a task
```

## Build

```bash
go build -o tasks ./cmd/tasks/
```

## Data

Tasks are stored locally in `tasks.json`.
