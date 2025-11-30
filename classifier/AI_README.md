# AI Agent Guide

This project is a transaction classifier for the Andamio Platform, built in Go. It identifies and extracts data from Cardano transactions based on specific patterns.

## Project Structure

- `cmd/classifier/main.go`: The entry point of the application. It iterates through a list of transaction hashes, fetches the transaction data, and runs it through various classifiers.
- `handlers`: Contains the classifier logic, organized by domain:
    - `admincourse`: Handlers for admin-related course actions (create course, update teachers).
    - `studentcourse`: Handlers for student-related actions (enroll, submit assignment, etc.).
    - `teachercourse`: Handlers for teacher-related actions (manage modules, assess assignments).
    - `useraccesstoken`: Handlers for minting user access tokens.
- `models`: Defines the data structures (structs) for the extracted transaction data.
- `utils`: Shared utility functions (e.g., fetching transactions).

## How to Extend

1.  **Define a Model**: If you are adding a new transaction type, define a struct in `models/models.go` to represent the data you want to extract.
2.  **Create a Handler**: Create a new function in the appropriate `handlers` package.
    - The function should accept `*cardano.Tx`.
    - It should return `(*models.YourModel, bool)`.
    - Use `config.Get()` to access necessary configuration values (e.g., policy IDs).
    - If the transaction matches, populate the model and return `(model, true)`.
    - If not, return `(nil, false)`.
3.  **Register in Main**: Update `cmd/classifier/main.go` to call your new handler and handle the returned data. Ensure `config.Init` is called at the start of `main`.

## Running the Project

```bash
go run cmd/classifier/main.go
```

```bash
go test ./...
```

## Reusability

You can use the classifier handlers in other projects by importing the packages. However, you must initialize the configuration first.

### Example

```go
package main

import (
	"fmt"

	"github.com/Salvionied/apollo/constants"
	"github.com/andamio-platform/transaction-specs/classifier/config"
	"github.com/andamio-platform/transaction-specs/classifier/handlers/studentcourse"
	"github.com/andamio-platform/transaction-specs/classifier/utils"
)

func main() {
	// 1. Initialize Config
	config.Init(constants.PREPROD)
	
	// 2. Set Course State Policy IDs (if needed by the handler)
	config.SetCourseStatePolicyIds([]string{"policy_id_1", "policy_id_2"})

	// 3. Get Transaction
	tx := utils.GetCardanoTx("tx_hash")

	// 4. Call Handler
	if model, ok := studentcourse.Enroll(tx); ok {
		fmt.Printf("Enrollment found: %+v\n", model)
	}
}
```
