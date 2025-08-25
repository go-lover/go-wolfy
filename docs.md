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

This function returns a `UserAccountInfo` struct. The full definition, including all available fields and sub-structures like `Slot` and `Skin`, can be found in the `types.go` file.

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

This function returns a `PlayerInfoResponse` struct. This is a complex structure that includes the user's game history and detailed statistics. For a complete list of all available fields, please refer to the `types.go` file.

---

### `GetPlayerInfo`

Retrieves the detailed public profile for any player by their username or unique ID. This is the ideal function for looking up other players to see their stats, rank, and game history.

The data returned is the same structure as `GetSelfInfo`.

**Function Signature**
```go
func (c *Client) GetPlayerInfo(usernameOrID string) (*PlayerInfoResponse, error)
```

**Parameters**
*   `usernameOrID (string)`: The exact username or the unique user ID of the player you want to look up.

**Return Values**
*   `(*PlayerInfoResponse, nil)`: On success, returns a pointer to a `PlayerInfoResponse` struct containing the player's public profile data.
*   `(nil, error)`: Returns an error if the user is not found, or if the API call fails for other reasons.

**Usage Example**

This example looks up the user "SOMEONE" and displays their rank, total kill count, and the outcome of their most recent game.

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

	// The username of the player we want to find.
	usernameToFind := "SOMEONE"

	// Fetch the public profile information for the specified user.
	playerInfo, err := client.GetPlayerInfo(usernameToFind)
	if err != nil {
		log.Fatalf("Failed to get info for player '%s': %v", usernameToFind, err)
	}

	// Display some of the retrieved data.
	fmt.Printf("\n--- Public Profile for %s ---\n", playerInfo.User.Username)
	fmt.Printf("  Rank: %d\n", playerInfo.User.Rank)
	fmt.Printf("  Total Kills: %d\n", playerInfo.Statistics.Individual.KillCount)

	// Display details from their most recent game.
	if len(playerInfo.History) > 0 {
		mostRecentGame := playerInfo.History[0]
		outcome := "Loss"
		if mostRecentGame.Winner {
			outcome = "Win"
		}
		fmt.Printf("  Most Recent Game: %s as role '%s'\n", outcome, mostRecentGame.Role)
	}
	fmt.Println("---------------------------------")
}
```

**Response Data Structure**

This function returns a `PlayerInfoResponse` struct. For a complete list of all available fields and sub-structures, please refer to the `types.go` file.

---

### `GetUserID`

Finds a user by their exact username and returns their unique user ID as a string. This is a convenient helper function to use when you have a username but need the ID to perform other actions, such as `AddFriend` or `RemoveFriend`.

**Function Signature**
```go
func (c *Client) GetUserID(username string) (string, error)
```

**Parameters**
*   `username (string)`: The exact username of the player to find.

**Return Values**
*   `(string, nil)`: On success, returns the user's unique ID.
*   `("", error)`: Returns an empty string and an error if the user is not found or if the API call fails.

**Usage Example**

This example demonstrates how to find the user ID for the username "SOMEONE".

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

	usernameToFind := "SOMEONE"
	fmt.Printf("\nLooking up user ID for: %s\n", usernameToFind)

	// Use the GetUserID helper function.
	userID, err := client.GetUserID(usernameToFind)
	if err != nil {
		log.Fatalf("Could not find user '%s': %v", usernameToFind, err)
	}

	fmt.Printf("Successfully found ID: %s\n", userID)
}
```

**Notes**

This function works by calling `GetPlayerInfo` internally and extracting the `ID` from the response.

---

### `GetFriendList`

Retrieves the friend list of the currently authenticated user. The function returns a slice of strings, where each string is the unique user ID of a friend.

**Function Signature**
```go
func (c *Client) GetFriendList() ([]string, error)
```

**Parameters**
*   None.

**Return Values**
*   `([]string, nil)`: On success, returns a slice of user ID strings.
*   `(nil, error)`: Returns an error if the API call fails.

**Usage Example**

This example fetches the authenticated user's friend list and prints the total number of friends, along with the first 5 user IDs in the list.

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

	// Fetch the user ID list for all friends.
	friendIDs, err := client.GetFriendList()
	if err != nil {
		log.Fatalf("Failed to get friend list: %v", err)
	}

	fmt.Printf("\nSuccessfully fetched friend list. You have %d friends.\n", len(friendIDs))

	// Print the first 5 friend IDs for demonstration.
	if len(friendIDs) > 0 {
		fmt.Println("--- First 5 Friend IDs ---")
		for i, id := range friendIDs {
			if i >= 5 {
				break
			}
			fmt.Printf("  - %s\n", id)
		}
		fmt.Println("--------------------------")
	}
}
```

**Notes**

This function returns a list of IDs. If you need the usernames or other details for these friends, you would need to loop through the returned slice and call `GetPlayerInfo` for each ID.

---

### `AddFriend`

Sends a friend request to another user. This action requires the unique user ID of the person you want to add.

**Function Signature**
```go
func (c *Client) AddFriend(userID string) (*MessageResponse, error)
```

**Parameters**
*   `userID (string)`: The unique user ID of the player to send a friend request to.

**Return Values**
*   `(*MessageResponse, nil)`: On success, returns a pointer to a `MessageResponse` struct containing the API's confirmation message (e.g., "Friend request sent.").
*   `(nil, error)`: Returns an error if the user ID is invalid or if the API call fails.

**Usage Example**

This example first finds the user ID for the username "SOMEONE" and then sends a friend request to that user.

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

	// Change this to a real username that you want to add.
	usernameToAdd := "SOMEONE"
	fmt.Printf("\nAttempting to send a friend request to '%s'...\n", usernameToAdd)

	// Step 1: Get the user's ID from their username.
	userID, err := client.GetUserID(usernameToAdd)
	if err != nil {
		log.Fatalf("Could not find user '%s': %v", usernameToAdd, err)
	}
	fmt.Printf("Found user ID: %s\n", userID)

	// Step 2: Send the friend request using their ID.
	response, err := client.AddFriend(userID)
	if err != nil {
		log.Fatalf("Failed to send friend request: %v", err)
	}

	// The API returns a simple confirmation message.
	fmt.Printf("API Response: %s\n", response.Message)
}
```

**Response Data Structure**

This function returns a `MessageResponse` struct, which contains a single field:
```go
type MessageResponse struct {
	Message string `json:"message"`
}
```

---


### `RemoveFriend`

Removes a user from your friend list. This action requires the unique user ID of the person you want to remove.

**Function Signature**
```go
func (c *Client) RemoveFriend(userID string) (*MessageResponse, error)
```

**Parameters**
*   `userID (string)`: The unique user ID of the player to remove from your friend list.

**Return Values**
*   `(*MessageResponse, nil)`: On success, returns a pointer to a `MessageResponse` struct containing the API's confirmation message.
*   `(nil, error)`: Returns an error if the user ID is invalid, if the user is not on your friend list, or if the API call fails.

**Usage Example**

This example first finds the user ID for the username "SOMEONE" and then removes that user from the friend list.

```go
package main

import (
	"fmt"
	"log"
	"os"

	wolfyclient "github.com/go-lover/go-wolfy"
)

func main() {
	mySessionToken := os.Getenv("WOLF Y_TOKEN")
	if mySessionToken == "" {
		log.Fatal("WOLFY_TOKEN environment variable not set.")
	}

	client, err := wolfyclient.NewClient(mySessionToken)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	fmt.Println("Client created successfully!")

	// Change this to a real username that is currently on your friend list.
	usernameToRemove := "SOMEONE"
	fmt.Printf("\nAttempting to remove '%s' from your friend list...\n", usernameToRemove)

	// Step 1: Get the user's ID from their username.
	userID, err := client.GetUserID(usernameToRemove)
	if err != nil {
		log.Fatalf("Could not find user '%s': %v", usernameToRemove, err)
	}
	fmt.Printf("Found user ID: %s\n", userID)

	// Step 2: Remove the friend using their ID.
	response, err := client.RemoveFriend(userID)
	if err != nil {
		log.Fatalf("Failed to remove friend: %v", err)
	}

	// The API returns a simple confirmation message.
	fmt.Printf("API Response: %s\n", response.Message)
}
```

**Response Data Structure**

This function returns a `MessageResponse` struct, which contains a single field:
```go
type MessageResponse struct {
	Message string `json:"message"`
}
```

---

### `GetFriendLeaderboard`

Retrieves the leaderboard composed exclusively of the authenticated user's friends. This is useful for comparing ranks, XP, and Elo with people you know.

**Function Signature**
```go
func (c *Client) GetFriendLeaderboard() ([]LeaderboardEntry, error)
```

**Parameters**
*   None.

**Return Values**
*   `([]LeaderboardEntry, nil)`: On success, returns a slice of `LeaderboardEntry` structs, where each entry represents a friend.
*   `(nil, error)`: Returns an error if the API call fails.

**Usage Example**

This example fetches the friend leaderboard and prints the username, rank, and XP for the top 5 friends.

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

	// Fetch the leaderboard for your friends.
	friendLeaderboard, err := client.GetFriendLeaderboard()
	if err != nil {
		log.Fatalf("Failed to get friend leaderboard: %v", err)
	}

	fmt.Printf("\n--- Friend Leaderboard (Top %d) ---\n", len(friendLeaderboard))
	
	// Loop through the leaderboard and display each friend's info.
	for i, friend := range friendLeaderboard {
		if i >= 5 { // Limit to 5 for this example
			break
		}
		fmt.Printf("  %d. %-20s | Rank: %-5d | XP: %d\n", i+1, friend.Username, friend.Rank, friend.XP)
	}
	fmt.Println("------------------------------------")
}
```

**Response Data Structure**

This function returns a slice of `LeaderboardEntry` structs. For a complete list of all available fields, please refer to the `types.go` file.

---

### `Logout`

Invalidates the current session token on the Wolfy.net server. After this call is successful, the `authToken` used to create the client will no longer be valid for any future API requests.

**Function Signature**
```go
func (c *Client) Logout() (*MessageResponse, error)
```

**Parameters**
*   None.

**Return Values**
*   `(*MessageResponse, nil)`: On success, returns a pointer to a `MessageResponse` struct containing the API's confirmation message.
*   `(nil, error)`: Returns an error if the API call fails.

**Usage Example**

This example shows how to properly log out, invalidating the session.

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

	fmt.Println("\nAttempting to log out and invalidate the session token...")

	// Call the Logout function.
	response, err := client.Logout()
	if err != nil {
		log.Fatalf("Failed to log out: %v", err)
	}

	fmt.Printf("API Response: %s\n", response.Message)
	fmt.Println("The session token is now invalid.")

	// Any subsequent calls with this client will likely fail.
	// For example, the following call should now return an authentication error.
	_, err = client.GetSelfInfo()
	if err != nil {
		fmt.Printf("\nAs expected, a new API call failed with an error: %v\n", err)
	}
}
```

**Response Data Structure**

This function returns a `MessageResponse` struct, which contains a single field:
```go
type MessageResponse struct {
	Message string `json:"message"`
}
```

---

## Settings & Actions

### `ChangeUsername`

Changes the username for the currently authenticated user.

**Function Signature**
```go
func (c *Client) ChangeUsername(newUsername string) (*MessageResponse, error)
```

**Parameters**
*   `newUsername (string)`: The new username to set for the account.

**Return Values**
*   `(*MessageResponse, nil)`: On success, returns a pointer to a `MessageResponse` struct containing the API's confirmation message.
*   `(nil, error)`: Returns an error if the username is already taken, is invalid, or if the API call fails.

**Usage Example**

This example attempts to change the user's username to "NewWolfyName".

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

	newUsername := "NewWolfyName"
	fmt.Printf("\nAttempting to change username to '%s'...\n", newUsername)

	// Call the ChangeUsername function.
	response, err := client.ChangeUsername(newUsername)
	if err != nil {
		// This will trigger if the name is taken, invalid, etc.
		log.Fatalf("Failed to change username: %v", err)
	}

	fmt.Printf("API Response: %s\n", response.Message)
}
```

**Response Data Structure**

This function returns a `MessageResponse` struct, which contains a single field:
```go
type MessageResponse struct {
	Message string `json:"message"`
}
```

---

### `ChangeEmail`

Changes the email address for the currently authenticated user.

**Function Signature**
```go
func (c *Client) ChangeEmail(newEmail string) (*MessageResponse, error)
```

**Parameters**
*   `newEmail (string)`: The new email address to associate with the account.

**Return Values**
*   `(*MessageResponse, nil)`: On success, returns a pointer to a `MessageResponse` struct containing the API's confirmation message.
*   `(nil, error)`: Returns an error if the email address is invalid, already in use, or if the API call fails.

**Usage Example**

This example attempts to change the user's email address.

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

	newEmail := "new-email@example.com"
	fmt.Printf("\nAttempting to change email to '%s'...\n", newEmail)

	// Call the ChangeEmail function.
	response, err := client.ChangeEmail(newEmail)
	if err != nil {
		log.Fatalf("Failed to change email: %v", err)
	}

	fmt.Printf("API Response: %s\n", response.Message)
}
```

**Response Data Structure**

This function returns a `MessageResponse` struct, which contains a single field:
```go
type MessageResponse struct {
	Message string `json:"message"`
}
```

---

### `ChangePassword`

Changes the password for the currently authenticated user. This action requires both the user's current (old) password and the new password they wish to set.

**Function Signature**
```go
func (c *Client) ChangePassword(oldPassword, newPassword string) (*MessageResponse, error)
```

**Parameters**
*   `oldPassword (string)`: The user's current password.
*   `newPassword (string)`: The new password to set.

**Return Values**
*   `(*MessageResponse, nil)`: On success, returns a pointer to a `MessageResponse` struct containing the API's confirmation message.
*   `(nil, error)`: Returns an error if the old password is incorrect, if the new password is invalid, or if the API call fails.

**Usage Example**

This example demonstrates how to change the user's password.

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

	// NOTE: For security, it is highly recommended to get passwords from a secure
	// source (like environment variables or a prompt) rather than hardcoding them.
	oldPassword := os.Getenv("WOLFY_OLD_PASS")
	newPassword := os.Getenv("WOLFY_NEW_PASS")
	if oldPassword == "" || newPassword == "" {
		log.Fatal("Please set WOLFY_OLD_PASS and WOLFY_NEW_PASS environment variables.")
	}

	client, err := wolfyclient.NewClient(mySessionToken)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	fmt.Println("Client created successfully!")

	fmt.Println("\nAttempting to change password...")

	// Call the ChangePassword function.
	response, err := client.ChangePassword(oldPassword, newPassword)
	if err != nil {
		log.Fatalf("Failed to change password: %v", err)
	}

	fmt.Printf("API Response: %s\n", response.Message)
}
```

**Response Data Structure**

This function returns a `MessageResponse` struct, which contains a single field:
```go
type MessageResponse struct {
	Message string `json:"message"`
}
```

---

## Skin Management

### `GetUserSkin`

Fetches the rendered skin image for a given user ID. This is an unauthenticated call that does not require a valid session token to be used on the client, however, a token is still required to create the client instance itself.

The function returns the raw image data as a byte slice (`[]byte`), which can be saved directly to a file (e.g., `user_skin.png`).

**Function Signature**
```go
func (c *Client) GetUserSkin(userID, format, profile, size string) ([]byte, error)
```

**Parameters**
*   `userID (string)`: The unique user ID of the player whose skin you want to render.
*   `format (string)`: The desired image format. It is highly recommended to use the exported constants: `wolfyclient.SkinFormatPNG` or `wolfyclient.SkinFormatSVG`.
*   `profile (string)`: The type of render. Use constants: `wolfyclient.SkinProfileFull` (full body), `wolfyclient.SkinProfileCenter` (face), or `wolfyclient.SkinProfileRight` (face, right-facing).
*   `size (string)`: The desired image dimensions. This only applies to the PNG format. Use constants: `wolfyclient.SkinSizeDefault`, `wolfyclient.SkinSizeLarge`, or `wolfyclient.SkinSizeSmall`.

**Return Values**
*   `([]byte, nil)`: On success, returns the raw image data.
*   `(nil, error)`: Returns an error if the user ID is invalid or if the server fails to render the image.

**Usage Example**

This example first finds the user ID for "SOMEONE", then downloads a large, full-profile PNG of their skin and saves it to `SOMEONE_skin.png`.

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

	// Step 1: Find the user ID.
	username := "SOMEONE"
	userID, err := client.GetUserID(username)
	if err != nil {
		log.Fatalf("Could not find user %s: %v", username, err)
	}
	fmt.Printf("Found user %s with ID: %s\n", username, userID)

	// Step 2: Download the skin using the ID and desired format options.
	fmt.Println("\nDownloading large, full-profile PNG of the skin...")
	imageData, err := client.GetUserSkin(
		userID,
		wolfyclient.SkinFormatPNG,
		wolfyclient.SkinProfileFull,
		wolfyclient.SkinSizeLarge,
	)
	if err != nil {
		log.Fatalf("Failed to download skin: %v", err)
	}

	// Step 3: Save the downloaded image data to a file.
	filename := "SOMEONE_skin.png"
	err = os.WriteFile(filename, imageData, 0644)
	if err != nil {
		log.Fatalf("Failed to save image to file: %v", err)
	}
	fmt.Printf("Success! Skin for %s saved to %s\n", username, filename)
}
```

---

### `UpdateSkinSlot`

Changes the equipped cosmetic items for a specific, unlocked skin slot owned by the authenticated user. This is a powerful function that allows you to programmatically change a user's appearance.

**Function Signature**
```go
func (c *Client) UpdateSkinSlot(slotID string, updates map[string]SkinPart) (*UpdateSkinSlotResponse, error)
```

**Parameters**
*   `slotID (string)`: The unique ID of the skin slot to modify. You can get this from the `Slots` slice in the `UserAccountInfo` struct returned by `GetAccountDetails`.
*   `updates (map[string]SkinPart)`: A map specifying which skin parts to change. The key is the part type (e.g., "top", "shoes"), and the value is a `SkinPart` struct defining the new item ID and color index.

**Return Values**
*   `(*UpdateSkinSlotResponse, nil)`: On success, returns a pointer to an `UpdateSkinSlotResponse` struct, which contains the user's updated list of slots and their new skin configuration.
*   `(nil, error)`: Returns an error if the slot ID is invalid, if the user does not own one of the specified cosmetic items, or if the API call fails.

**Usage Example**

This example first fetches the user's account details to find their active slot ID, and then uses that ID to change their equipped shoes and top.

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

	// Step 1: Get the active slot ID programmatically.
	fmt.Println("\nFetching account details to find the active slot ID...")
	accountInfo, err := client.GetAccountDetails()
	if err != nil {
		log.Fatalf("Could not get account details: %v", err)
	}
	var activeSlotID string
	for _, slot := range accountInfo.Slots {
		if slot.Equiped && slot.Unlocked {
			activeSlotID = slot.ID
			break
		}
	}
	if activeSlotID == "" {
		log.Fatal("Could not find an active, unlocked slot for this user.")
	}
	fmt.Printf("Found active slot ID: %s\n", activeSlotID)

	// Step 2: Define the skin parts to update.
	// This map will change the shoes to item "S1" and the top to item "001".
	skinUpdates := map[string]wolfyclient.SkinPart{
		"shoes": {ID: "S1", Color: 0},
		"top":   {ID: "001", Color: 0},
	}
	fmt.Println("\nAttempting to update skin slot...")

	// Step 3: Call the update function.
	updateResponse, err := client.UpdateSkinSlot(activeSlotID, skinUpdates)
	if err != nil {
		log.Fatalf("Failed to update skin slot: %v", err)
	}

	fmt.Println("Success! Skin slot updated.")
	fmt.Printf("New skin version is: %s\n", updateResponse.Version)
}
```

**Response Data Structure**

This function returns an `UpdateSkinSlotResponse` struct. For a complete list of all available fields, please refer to the `types.go` file.

---

## Game & Shop Data

### `GetSkinCatalog`

Retrieves the master catalog of all available cosmetic items in the game. This function is incredibly powerful for building tools that need to understand what items exist, their properties, and their names. The response is a large array containing every skin, hat, pet, etc.

**Function Signature**
```go
func (c *Client) GetSkinCatalog() ([]SkinElement, error)
```

**Parameters**
*   None.

**Return Values**
*   `([]SkinElement, nil)`: On success, returns a slice of `SkinElement` structs, representing the entire cosmetic catalog.
*   `(nil, error)`: Returns an error if the API call fails.

**Usage Example**

This example fetches the entire skin catalog and then performs a simple analysis to count how many items of type "hat" exist.

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

	// Fetch the entire cosmetic item catalog.
	fmt.Println("\nFetching skin catalog... (This may take a moment)")
	skinCatalog, err := client.GetSkinCatalog()
	if err != nil {
		log.Fatalf("Failed to get skin catalog: %v", err)
	}
	fmt.Printf("Successfully fetched %d total items from the catalog.\n", len(skinCatalog))

	// Perform an analysis: count how many items are hats.
	hatCount := 0
	for _, item := range skinCatalog {
		if item.Type == "hat" {
			hatCount++
		}
	}

	fmt.Printf("Analysis complete: Found %d hats in the catalog.\n", hatCount)
}
```

**Response Data Structure**

This function returns a slice of `SkinElement` structs. This is a complex structure containing details like item rarity, price, and colors. For a complete list of all available fields, please refer to the `types.go` file.

---

### `GetCurrentDrop`

Retrieves detailed information about the current featured item drop. "Drops" are typically monthly themed releases containing exclusive, limited-time cosmetic packs.

**Function Signature**
```go
func (c *Client) GetCurrentDrop() (*CurrentDrop, error)
```

**Parameters**
*   None.

**Return Values**
*   `(*CurrentDrop, nil)`: On success, returns a pointer to a `CurrentDrop` struct, which contains the drop's name, duration, and a list of all cosmetic packs included in it.
*   `(nil, error)`: Returns an error if the API call fails.

**Usage Example**

This example fetches the current drop and displays its name, its end date, and the name and price of the first pack available in the drop.

```go
package main

import (
	"fmt"
	"log"
	"os"
	"time"

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

	// Fetch the current featured item drop.
	fmt.Println("\nFetching current item drop...")
	currentDrop, err := client.GetCurrentDrop()
	if err != nil {
		log.Fatalf("Failed to get current drop: %v", err)
	}

	fmt.Printf("\n--- Current Drop: %s ---\n", currentDrop.Name)

	// Parse and display the end date of the drop.
	endDate, err := time.Parse(time.RFC3339, currentDrop.End)
	if err == nil {
		fmt.Printf("  This drop ends on: %s\n", endDate.Format("January 2, 2006"))
	}

	// Display details about the first pack in the drop.
	if len(currentDrop.Packs) > 0 {
		firstPack := currentDrop.Packs[0]
		fmt.Printf("\n--- Featured Pack ---\n")
		fmt.Printf("  Pack Name: %s\n", firstPack.Name)
		fmt.Printf("  Price: %d %s\n", firstPack.Price, firstPack.Currency)
		fmt.Printf("  Rarity: %s\n", firstPack.Rarity)
		fmt.Printf("  Number of items in pack: %d\n", len(firstPack.SkinElements))
		fmt.Println("---------------------")
	}
}
```

**Response Data Structure**

This function returns a `CurrentDrop` struct, which is a complex structure containing nested `DropPack` and `DropSkinElement` objects. For a complete list of all available fields, please refer to the `types.go` file.

---

### `GetDailyShopOffers`

Retrieves the list of daily offers from the in-game shop. This includes the free daily item, as well as a rotating selection of individual skins and cosmetic bundles available for purchase with Coins and Moons.

**Function Signature**
```go
func (c *Client) GetDailyShopOffers() ([]DailyOfferSet, error)
```

**Parameters**
*   None.

**Return Values**
*   `([]DailyOfferSet, nil)`: On success, returns a slice of `DailyOfferSet` structs. The first element (`[0]`) in the slice represents today's offers.
*   `(nil, error)`: Returns an error if the API call fails.

**Usage Example**

This example fetches the daily shop offers and prints the details of the "free" offer and the "coinsHigh" offer (a high-value item for Coins).

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

	// Fetch the daily shop offers.
	fmt.Println("\nFetching daily shop offers...")
	dailyOffers, err := client.GetDailyShopOffers()
	if err != nil {
		log.Fatalf("Failed to get daily shop offers: %v", err)
	}

	if len(dailyOffers) == 0 {
		log.Fatal("No daily offers were returned by the API.")
	}

	// Today's offers are the first element in the response slice.
	todaysOffers := dailyOffers[0].Elements

	// Example 1: Check the free daily item.
	freeOffer := todaysOffers.Free
	fmt.Println("\n--- Free Daily Item ---")
	if freeOffer.Collected {
		fmt.Println("  You have already collected today's free item.")
	} else {
		fmt.Printf("  Today's free item is: %d Coins.\n", freeOffer.Coins)
		fmt.Println("  You can claim this using the CollectDailyItem() function.")
	}
	fmt.Println("-----------------------")

	// Example 2: Check the high-value item for sale for Coins.
	coinsHighOffer := todaysOffers.CoinsHigh
	fmt.Println("\n--- High-Value Coin Offer ---")
	if coinsHighOffer.Skin != nil {
		fmt.Printf("  Item for sale: %s\n", coinsHighOffer.Skin.Name)
		fmt.Printf("  Price: %d %s\n", coinsHighOffer.Skin.Price, coinsHighOffer.Skin.Currency)
	} else {
		fmt.Println("  No individual skin is available in this slot today.")
	}
	fmt.Println("-----------------------------")
}
```

**Response Data Structure**

This function returns a slice of `DailyOfferSet` structs. This is a highly complex structure containing many nested objects. For a complete definition of all fields, please refer to the `types.go` file.

---

### `GetSubscriptionOffers`

Retrieves the list of available Alpha subscription plans. This includes details like price, duration (monthly, yearly, etc.), currency, and any associated discounts.

**Function Signature**
```go
func (c *Client) GetSubscriptionOffers() ([]SubscriptionOffer, error)
```

**Parameters**
*   None.

**Return Values**
*   `([]SubscriptionOffer, nil)`: On success, returns a slice of `SubscriptionOffer` structs, where each entry represents a different subscription plan.
*   `(nil, error)`: Returns an error if the API call fails.

**Usage Example**

This example fetches all available Alpha subscription offers and prints their details.

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

	// Fetch the available subscription plans.
	fmt.Println("\nFetching Alpha subscription offers...")
	subOffers, err := client.GetSubscriptionOffers()
	if err != nil {
		log.Fatalf("Failed to get subscription offers: %v", err)
	}

	fmt.Println("\n--- Available Alpha Subscriptions ---")
	for _, offer := range subOffers {
		fmt.Printf("  - Plan: %s\n", offer.ID)
		fmt.Printf("    Price: %.2f %s\n", offer.Price, offer.Currency)
		fmt.Printf("    Duration: %d %s(s)\n", offer.IntervalCount, offer.Interval)
		if offer.MostPopular {
			fmt.Println("    (Most Popular!)")
		}
		fmt.Println() // Blank line for spacing
	}
	fmt.Println("-------------------------------------")
}
```

**Response Data Structure**

This function returns a slice of `SubscriptionOffer` structs. For a complete list of all available fields, please refer to the `types.go` file.

---

### `GetMoonOffers`

Retrieves the list of available Moon currency packs. Moons are a premium currency in Wolfy, and this function provides all the details for the different purchase options, including price, currency, and bonus amounts.

**Function Signature**
```go
func (c *Client) GetMoonOffers() ([]MoonOffer, error)
```

**Parameters**
*   None.

**Return Values**
*   `([]MoonOffer, nil)`: On success, returns a slice of `MoonOffer` structs, where each entry represents a different currency pack.
*   `(nil, error)`: Returns an error if the API call fails.

**Usage Example**

This example fetches all available Moon currency packs and prints their details.

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

	// Fetch the available Moon currency packs.
	fmt.Println("\nFetching Moon currency offers...")
	moonOffers, err := client.GetMoonOffers()
	if err != nil {
		log.Fatalf("Failed to get Moon offers: %v", err)
	}

	fmt.Println("\n--- Available Moon Packs ---")
	for _, offer := range moonOffers {
		totalMoons := offer.Moons + offer.Bonus
		fmt.Printf("  - ID: %s\n", offer.ID)
		fmt.Printf("    Get %d Moons (includes %d bonus) for %.2f %s\n", totalMoons, offer.Bonus, offer.Price, offer.Currency)
		if offer.Tag != "" {
			fmt.Printf("    (%s deal!)\n", offer.Tag)
		}
		fmt.Println() // Blank line for spacing
	}
	fmt.Println("----------------------------")
}
```

**Response Data Structure**

This function returns a slice of `MoonOffer` structs. For a complete list of all available fields, please refer to the `types.go` file.

---

### `CollectDailyItem`

Attempts to claim the free daily item from the shop for the authenticated user. This function is an action and does not return complex data.

**Function Signature**
```go
func (c *Client) CollectDailyItem() (string, error)
```

**Parameters**
*   None.

**Return Values**
*   `(string, nil)`: On success, returns the raw text response from the API. This is typically a simple string like `"OK"` if the collection was successful, or a descriptive message if the item has already been collected (e.g., `"already_collected"`).
*   `("", error)`: Returns an empty string and an error if the API call fails.

**Usage Example**

This example first checks the daily offers to see if the free item has been collected. If it has not, it calls `CollectDailyItem` to claim it.

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

	// Step 1: Check if the daily item has already been collected.
	fmt.Println("\nChecking status of daily free item...")
	dailyOffers, err := client.GetDailyShopOffers()
	if err != nil {
		log.Fatalf("Could not check daily offers: %v", err)
	}

	if len(dailyOffers) == 0 {
		log.Fatal("No daily offers found.")
	}

	if dailyOffers[0].Elements.Free.Collected {
		fmt.Println("You have already collected today's free item.")
	} else {
		// Step 2: If not collected, attempt to claim it.
		fmt.Println("Free item has not been collected yet. Attempting to claim it now...")
		
		response, err := client.CollectDailyItem()
		if err != nil {
			log.Fatalf("Failed to collect daily item: %v", err)
		}
		
		fmt.Printf("Successfully sent collection request. API Response: %s\n", response)
	}
}
```
