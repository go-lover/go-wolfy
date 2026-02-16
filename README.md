# [go-wolfy](https://github.com/go-lover/go-wolfy/wiki) [![Go Reference](https://pkg.go.dev/badge/github.com/go-lover/go-wolfy.svg)](https://pkg.go.dev/github.com/go-lover/go-wolfy) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

An unofficial, robust, and strictly-typed Go client library for the **[Wolfy.net](https://wolfy.net)** API.

`go-wolfy` allows you to programmatically interact with user profiles, manage friend lists, browse the item shop, and even customize your character's skin directly via Go code.

## Features

- **Type-Safe:** Complete Go structs for every API response (Profiles, Stats, History, Shop).
- **Session Management:** Simple authentication via session tokens with automatic validation.
- **Skin Customization:** Render user skins (PNG/SVG) and programmatically change equipped items/colors.
- **Social Integration:** Full support for friend lists, adding/removing users, and friend leaderboards.
- **Economy & Shop:** Access the full skin catalog, daily rotating offers, and featured drops.
- **Detailed Statistics:** Deep access to player win rates, role-specific stats, and game history.

## Installation

```bash
go get github.com/go-lover/go-wolfy
```

## Authentication

Because Wolfy.net uses Captcha for login, this library uses your **Session Token** (Cookie) to authenticate.

1.  Log in to [Wolfy.net](https://wolfy.net) in your browser.
2.  Open Developer Tools (`F12`) -> **Application/Storage** tab.
3.  Go to **Cookies** -> `https://wolfy.net`.
4.  Copy the value of the `wolfy` cookie (starts with `s%3A...`).

**Treat this token like a password.** Do not share it or commit it to public repositories.

**Note:** In the code example below, the session token is set as an environment variable named `WOLFY_TOKEN`. Ensure you set this environment variable before running the code.

## Usage Example

```go
package main

import (
	"fmt"
	"log"
	"os"

	wolfyclient "github.com/go-lover/go-wolfy"
)

func main() {
	// 1. Initialize the client with your token
	token := os.Getenv("WOLFY_TOKEN")
	client, err := wolfyclient.NewClient(token)
	if err != nil {
		log.Fatalf("Authentication failed: %v", err)
	}

	// 2. Fetch your own public profile
	me, err := client.GetSelfInfo()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfully logged in as %s!\n", me.User.Username)
	fmt.Printf("Rank: %d | XP: %d | Elo: %d\n", me.User.Rank, me.User.XP, me.User.Elo)
}
```

## Documentation

For detailed guides and full API references, please visit our **[GitHub Wiki](https://github.com/go-lover/go-wolfy/wiki)**.

- [Authentication Guide](https://github.com/go-lover/go-wolfy/wiki/Authentication)
- [Account & Settings](https://github.com/go-lover/go-wolfy/wiki/Account-Settings)
- [Players & Social](https://github.com/go-lover/go-wolfy/wiki/Players-Social)
- [Skins & Cosmetics](https://github.com/go-lover/go-wolfy/wiki/Skins-Cosmetics)
- [Shop & Economy](https://github.com/go-lover/go-wolfy/wiki/Shop-Economy)

## Disclaimer

This is an unofficial library and is not affiliated with, sponsored by, or endorsed by Wolfy.net. 
- Use this tool responsibly. 
- High-frequency requests may lead to rate limiting or account restrictions.
- **Use at your own risk.**

## License

Distributed under the **MIT License**. See `LICENSE` for more information.


