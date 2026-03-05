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

## Development

```bash
make build   # compile the binary
make test    # run tests
make check   # run go vet
make clean   # remove compiled binary
```

## Data

Tasks are stored locally in `tasks.json`.
