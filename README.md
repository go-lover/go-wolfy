# Wolfy.net Go API Client [![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](https://golang.org)
*An unofficial Go client library for interacting with the `wolfy.net` web API.*

This library handles session management via an authentication token and provides simple, typed methods for most common API actions, including fetching user data, managing friends, updating account settings, and rendering user skins.

## Features

-   **Token Validation:** Automatically validates the session token upon client creation.
-   **Typed Structs:** Clear Go structs for all JSON API responses.
-   **Advanced Skin Rendering:** A powerful function to download user skins with multiple formats (PNG, SVG), profiles (full, center, right), and sizes (small, large).
-   **Rich Game History:** Access strongly-typed data for a player's recent games, including roles, outcomes, kills, and full game settings.

### Authentication (Required)

This library **does not** handle username/password login because the website is protected by a captcha. Instead, it authenticates by using a session token that you must retrieve from your browser.

**How to get your session token:**

1.  Open your web browser (e.g., Chrome, Firefox) and log in to `wolfy.net` normally.
2.  After logging in, open the **Developer Tools** (usually by pressing **F12**).
3.  Go to the **Application** tab (in Chrome) or the **Storage** tab (in Firefox).
4.  On the left-side menu, find the "Cookies" section and click on `https://wolfy.net`.
5.  Find the cookie with the name `wolfy`.
6.  Copy the long string of text from the **"Value"** column. This is your session token.

You will provide this token when creating a new client instance. The library will immediately check if the token is valid.

## Usage Example

Here is a complete example of how to import the library, create a client, and use some of its key features.

```go
package main

import (
	"fmt"
	"log"
	"os"

	// Import the client library using your project's module path
	"github.com/go-lover/go-wolfy"
)

func main() {
	// 1. Paste the session token you copied from your browser. Should look like this: "s%3A1ae7777d-8501-4bb1-858f-11fb9ecf66a8.xPQ4zq69F59DPLBfZuhZBoDpU0Cgmozw"
	mySessionToken := "YOUR_TOKEN"

	// 2. Create a new client. The library automatically validates the token.
	client, err := wolfyclient.NewClient(mySessionToken)
	if err != nil {
		log.Fatalf("Failed to create client (is your token valid?): %v", err)
	}
	fmt.Println("Client created and token validated successfully!")

	// 3. Find a user's ID.
	username := "SOMEONE"
	userID, err := client.FindUserID(username)
	if err != nil {
		log.Fatalf("Could not find user %s: %v", username, err)
	}
	fmt.Printf("Found user %s with ID: %s\n", username, userID)
	
	// 4. Download a specific version of their skin.
	fmt.Println("\nDownloading large, right-facing PNG profile...")
	imageData, err := client.GetUserSkin(
		userID,
		wolfyclient.SkinFormatPNG,
		wolfyclient.SkinProfileRight,
		wolfyclient.SkinSizeLarge,
	)
	if err != nil {
		log.Fatalf("Failed to download skin: %v", err)
	}

	// 5. Save the downloaded image to a file.
	filename := "user_skin.png"
	err = os.WriteFile(filename, imageData, 0644)
	if err != nil {
		log.Fatalf("Failed to save image to file: %v", err)
	}
	fmt.Printf("Success! Skin saved to %s\n", filename)
}
```

## API Reference

### Client Initialization
---

#### `func NewClient(authToken string) (*Client, error)`
Creates and configures a new API client. It immediately validates the `authToken` by making a test API call. Returns an error if the token is invalid or expired.

#### `func (c *Client) SetSessionCookie(token string)`
Updates the session cookie on an existing client instance.

### Account & Social Methods
---

#### `func (c *Client) GetAccountDetails() (*UserAccountInfo, error)`
Retrieves the detailed private profile for the currently authenticated user, including email, currency balances, owned slots, and account settings.

#### `func (c *Client) GetSelfInfo() (*PlayerInfoResponse, error)`
Retrieves the detailed public profile (leaderboard version) for the currently authenticated user.

#### `func (c *Client) GetPlayerInfo(usernameOrID string) (*PlayerInfoResponse, error)`
Retrieves the detailed public profile for any player by their username or ID, including their game history.

#### `func (c *Client) GetUserID(username string) (string, error)`
Finds a user by their exact username and returns their unique ID.

#### `func (c *Client) GetFriendList() ([]string, error)`
Retrieves a slice of user IDs representing the authenticated user's friend list.

#### `func (c *Client) AddFriend(userID string) (*MessageResponse, error)`
Sends a friend request to the specified user ID.

#### `func (c *Client) RemoveFriend(userID string) (*MessageResponse, error)`
Removes the specified user from the authenticated user's friend list.

#### `func (c *Client) GetFriendLeaderboard() ([]LeaderboardEntry, error)`
Retrieves the leaderboard of the authenticated user's friends, returning a slice of users with their rank and summary information.

#### `func (c *Client) Logout() (*MessageResponse, error)`
Invalidates the current user's session on the server.

### Settings & Actions
---

#### `func (c *Client) ChangeUsername(newUsername string) (*MessageResponse, error)`
Changes the authenticated user's username.

#### `func (c *Client) ChangeEmail(newEmail string) (*MessageResponse, error)`
Changes the authenticated user's email address.

#### `func (c *Client) ChangePassword(oldPassword, newPassword string) (*MessageResponse, error)`
Changes the authenticated user's password.

### Skin Management
---

#### `func (c *Client) GetUserSkin(userID, format, profile, size string) ([]byte, error)`
Fetches the rendered skin image for a given user ID. This is an unauthenticated call. It returns the raw image data as a byte slice (`[]byte`), which can be saved to a file. See the constants table for available options.

#### `func (c *Client) UpdateSkinSlot(slotID string, updates map[string]SkinPart) (*UpdateSkinSlotResponse, error)`
Changes the equipped cosmetic items for a specific skin slot. The `updates` map should contain the skin parts to change, e.g., `"top": SkinPart{ID:"002", Color:5}`.

### Game & Shop Data
---

#### `func (c *Client) GetSkinCatalog() ([]SkinElement, error)`
Retrieves the master catalog of all available cosmetic items in the game.

#### `func (c *Client) GetCurrentDrop() (*CurrentDrop, error)`
Retrieves details about the current featured item drop, including available cosmetic packs.

#### `func (c *Client) GetDailyShopOffers() ([]DailyOfferSet, error)`
Retrievelist of daily offers from the shop, including the free item and rotating skins.

#### `func (c *Client) GetSubscriptionOffers() ([]SubscriptionOffer, error)`
Retrieves the available Alpha subscription plans.

#### `func (c *Client) GetMoonOffers() ([]MoonOffer, error)`
Retrieves the available Moon currency purchase options.

#### `func (c *Client) CollectDailyItem() (string, error)`
Attempts to claim the free daily item from the shop. Returns a plain text response from the API.## Data Structures

---

The main data structure returned by the API for player lookups.

```go
// The full response for a player lookup.
type PlayerInfoResponse struct {
	User       PlayerUser         `json:"user"`
	Statistics PlayerStatistics   `json:"statistics"`
	History    []GameHistoryEntry `json:"history"`
}

// Detailed information about a player.
type PlayerUser struct {
	ID          string        `json:"id"`
	Username    string        `json:"username"`
	CreatedAt   string        `json:"createdAt"`
	Rank        int           `json:"rank"`
	XP          int           `json:"xp"`
	SkinVersion string        `json:"skinVersion,string"`
	Elo         int           `json:"elo"`
	Ranking     PlayerRanking `json:"ranking"`
}

// Other sub-structs are defined in wolfyclient/models.go
```    

## Disclaimer

This is an unofficial library and is not affiliated with, sponsored by, or endorsed by Wolfy.net. Please use this tool responsibly and be mindful of the service's terms of use. Do not abuse the API.

```
   __                         __                
  / /  __ __   ___ ____  ____/ /__ _  _____ ____
 / _ \/ // /  / _ `/ _ \/___/ / _ \ |/ / -_) __/
/_.__/\_, /   \_, /\___/   /_/\___/___/\__/_/   
     /___/   /___/                                      
