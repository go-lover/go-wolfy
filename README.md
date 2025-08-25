# Wolfy.net Go API Client [![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](https://golang.org)
*An unofficial Go client library for interacting with the `wolfy.net` web API.*

> This library handles session management via an authentication token and provides simple, typed methods to interact with nearly every aspect of the Wolfy.net API. It allows you to fetch detailed user profiles and game histories, manage friends, update account settings, and both render and programmatically change user skins. Furthermore, it offers complete access to the in-game economy, letting you retrieve the full skin catalog, daily shop offers, featured drops, and currency packs.

## Features

*   **Comprehensive Typed Structs:** Clean, predictable Go structs for every API response, from user profiles to the entire shop catalog, enabling safe and easy data access.

*   **Automatic Token Validation:** The client automatically validates the session token upon creation, providing immediate feedback and preventing runtime errors.

*   **Full Account Management:** Fetch detailed public profiles, private account data (like currency and email), and programmatically update settings like username, password, or email.

*   **Advanced Skin Customization:**
    *   **Render:** Download user skins with multiple formats (PNG, SVG), profiles, and sizes.
    *   **Edit:** Programmatically change the cosmetic items equipped in a user's active skin slot.

*   **Complete Social Integration:** Manage the entire friend lifecycle (add/remove/list) and retrieve the friend-specific leaderboard.

*   **Full Shop & Game Data Access:**
    *   Download the **entire skin catalog** of all in-game items.
    *   Get the current **featured item drop**.
    *   Retrieve all **daily shop offers**, including the free item.
    *   List all available **subscriptions** and **currency packs**.

*   **Rich Game History:** Access strongly-typed data for a player's recent games, including roles, outcomes, kills, and full game settings.

---

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

Here is a complete example of how to import the library, create a client, and download someone's skin from the username.

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

## Disclaimer

This is an unofficial library and is not affiliated with, sponsored by, or endorsed by Wolfy.net. Please use this tool responsibly and be mindful of the service's terms of use. Do not abuse the API.

```
   __                         __                
  / /  __ __   ___ ____  ____/ /__ _  _____ ____
 / _ \/ // /  / _ `/ _ \/___/ / _ \ |/ / -_) __/
/_.__/\_, /   \_, /\___/   /_/\___/___/\__/_/   
     /___/   /___/                                      
