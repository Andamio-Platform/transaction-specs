# AI Agent Guide

This project is a transaction classifier for the Andamio Platform, built in Go. It identifies and extracts data from Cardano transactions based on specific patterns.

## Project Structure

- `cmd/classifier/main.go`: The entry point of the application. It iterates through a list of transaction hashes, fetches the transaction data, and runs it through various classifiers.
- `internal/handlers`: Contains the classifier logic, organized by domain:
    - `admincourse`: Handlers for admin-related course actions (create course, update teachers).
    - `studentcourse`: Handlers for student-related actions (enroll, submit assignment, etc.).
    - `teachercourse`: Handlers for teacher-related actions (manage modules, assess assignments).
    - `useraccesstoken`: Handlers for minting user access tokens.
- `internal/models`: Defines the data structures (structs) for the extracted transaction data.
- `internal/utils`: Shared utility functions (e.g., fetching transactions).

## How to Extend

1.  **Define a Model**: If you are adding a new transaction type, define a struct in `internal/models/models.go` to represent the data you want to extract.
2.  **Create a Handler**: Create a new function in the appropriate `internal/handlers` package.
    - The function should accept `*cardano.Tx` and any necessary policy IDs.
    - It should return `(*models.YourModel, bool)`.
    - If the transaction matches, populate the model and return `(model, true)`.
    - If not, return `(nil, false)`.
3.  **Register in Main**: Update `cmd/classifier/main.go` to call your new handler and handle the returned data.

## Running the Project

```bash
go run cmd/classifier/main.go
```

## Testing

```bash
go test ./...
```
