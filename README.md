# TODO App CLI


# Overview
This is a CLI app for managing tasks through the CLI based on [cobra](https://github.com/spf13/cobra).

Todolist provides the following commands:

- `add` – Add a task to the list. Command argument should be the task name.
- `done` – Mark a task as done. Command argument should be task ID, starting from 1.
- `undone` – Mark task as not done. Command argument should be task ID, starting from 1.
- `list` – List the tasks that have not been done.
- `cleanup` – Remove from the store all tasks marked as done.

# Example usage of the app:
`$ todolist help`
```
Usage:
    todolist [command]
Available Commands:
    add Add task to the list
    cleanup Cleanup done tasks
    done Mark task as done
    help Help about any command
        list List all tasks still to do
    undone Mark task as not done
Flags:
    -h, --help help for todolist
Use "todolist [command] --help" for more information about a command.
```

# Code structure

code base consists of four layers:
- infra a.k.a. "frameworks and drivers" layer contains the tools required to connect to external services such as caching servers, databases, etc.
- delivery a.k.a. "controllers/presenters/gateways" contains the code responsible to to convert the received data in format that the use cases accept, send it to them, and return the response in a correct format.
- use case layer contains the application business rules logic. Here are located all services of the package.

# Application Logic

the CLI app eventually saves and deletes the data from *db.csv* which acts as the database for the tasks
the data is saved in each row in the following form: `id,description,active,deleted`

# Running the app

Run the following command before running `todolist` command.
- `go get`
- `go install`

# Testing the app

- `ginkgo -race -tags e2e ./...`