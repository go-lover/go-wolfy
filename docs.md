# Full API Documentation

Welcome to the detailed documentation for the `go-wolfy` API Client. This document provides in-depth explanations, parameter details, and usage examples for every function available in the library.

For a general overview and quick-start guide, please see the main [`README.md`](https://github.com/go-lover/go-wolfy/blob/main/README.md) file.

## Getting Started: Client Initialization

Before you can use any of the library's features, you must first install it and create an authenticated client instance.

### 1. Installation

Add the library to your project using `go get`:
```bash
go get github.com/go-lover/go-wolfy
```

### 2. Authentication

This library requires a session token to interact with authenticated endpoints. Please follow the guide in the [main README](https://github.com/go-lover/go-wolfy#authentication-required) to retrieve your token from your browser.

### 3. Creating a Client

All interactions with the API are performed through a `Client` object. The `NewClient` function is the entry point to the library. It takes your session token and immediately validates it to ensure it's active.

It is highly recommended to provide your token via an environment variable rather than hardcoding it in your source code.

**Basic Client Creation Example:**
```go
package main

import (
	"fmt"
	"log"
	"os"

	wolfyclient "github.com/go-lover/go-wolfy"
)

func main() {
	// Get the session token from an environment variable for security.
	// In your terminal, run: export WOLFY_TOKEN="s%3Ayour-token..."
	mySessionToken := os.Getenv("WOLFY_TOKEN")
	if mySessionToken == "" {
		log.Fatal("WOLFY_TOKEN environment variable not set.")
	}

	// Create a new client. This call will fail if the token is invalid.
	client, err := wolfyclient.NewClient(mySessionToken)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	fmt.Println("Client created and authenticated successfully!")
	
	// You can now use the 'client' object to call other library functions.
}
```

## Account & Social Methods

### `GetAccountDetails`

Retrieves the detailed, private account information for the currently authenticated user. This is the most comprehensive data available for a user, including sensitive information like email, currency balances, owned items, and account settings.

**Function Signature**
```go
func (c *Client) GetAccountDetails() (*UserAccountInfo, error)
```

**Parameters**
*   None.

**Return Values**
*   `(*UserAccountInfo, nil)`: On success, returns a pointer to a `UserAccountInfo` struct containing the user's detailed account data.
*   `(nil, error)`: Returns an error if the API call fails due to an invalid token, network issues, or other server-side problems.

**Usage Example**

This example creates a client, fetches the authenticated user's account details, and prints their username, currency balances, and whether they have two-factor authentication enabled.

```go
package main

import (
	"fmt"
	"log"
	"os"

	wolfyclient "github.com/go-lover/go-wolfy"
)

func main() {
	mySessionToken := os.Getenv("WOLFY_TOKEN")
	if mySessionToken == "" {
		log.Fatal("WOLFY_TOKEN environment variable not set.")
	}

	client, err := wolfyclient.NewClient(mySessionToken)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	fmt.Println("Client created successfully!")

	// Fetch the detailed account information.
	accountInfo, err := client.GetAccountDetails()
	if err != nil {
		log.Fatalf("Failed to get account details: %v", err)
	}

	// Display some of the retrieved private data.
	fmt.Println("\n--- Your Private Account Details ---")
	fmt.Printf("  Username: %s\n", accountInfo.Username)
	fmt.Printf("  Email: %s\n", accountInfo.Email)
	fmt.Printf("  Coins: %d\n", accountInfo.Coins)
	fmt.Printf("  Moons: %d\n", accountInfo.Moons)
	fmt.Printf("  Two-Factor Enabled: %t\n", accountInfo.TwoFactorSecret)
	fmt.Println("----------------------------------")
}
```

**Response Data Structure**

This function returns a `UserAccountInfo` struct. The full definition, including all available fields and sub-structures like `Slot` and `Skin`, can be found in the `models.go` file.

---

### `GetSelfInfo`

Retrieves the detailed public profile for the currently authenticated user. This is the same data that is shown on the leaderboard, including detailed game history and role-by-role gameplay statistics.

**Function Signature**
```go
func (c *Client) GetSelfInfo() (*PlayerInfoResponse, error)
```

**Parameters**
*   None.

**Return Values**
*   `(*PlayerInfoResponse, nil)`: On success, returns a pointer to a `PlayerInfoResponse` struct containing the user's public profile data.
*   `(nil, error)`: Returns an error if the API call fails.

**Usage Example**

This example fetches the authenticated user's public profile and displays their username, total win count, and statistics for their most played role.

```go
package main

import (
	"fmt"
	"log"
	"os"

	wolfyclient "github.com/go-lover/go-wolfy"
)

func main() {
	mySessionToken := os.Getenv("WOLFY_TOKEN")
	if mySessionToken == "" {
		log.Fatal("WOLFY_TOKEN environment variable not set.")
	}

	client, err := wolfyclient.NewClient(mySessionToken)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	fmt.Println("Client created successfully!")

	// Fetch your own public profile information.
	selfInfo, err := client.GetSelfInfo()
	if err != nil {
		log.Fatalf("Failed to get self info: %v", err)
	}

	// Display some of the retrieved public data.
	fmt.Println("\n--- Your Public Profile & Statistics ---")
	fmt.Printf("  Username: %s\n", selfInfo.User.Username)
	fmt.Printf("  Total Wins: %d\n", selfInfo.Statistics.Individual.WinCount)
	
	// Display stats for the first role in the statistics list (usually the most played).
	if len(selfInfo.Statistics.Roles) > 0 {
		topRole := selfInfo.Statistics.Roles[0]
		fmt.Printf("  Top Role: %s\n", topRole.ID)
		fmt.Printf("    - Win Rate: %.2f%%\n", topRole.WinRate*100)
	}
	fmt.Println("----------------------------------------")
}
```

**Response Data Structure**

This function returns a `PlayerInfoResponse` struct. This is a complex structure that includes the user's game history and detailed statistics. For a complete list of all available fields, please refer to the `models.go` file.
