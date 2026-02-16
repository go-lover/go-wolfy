package wolfyclient

import (
	"fmt"
	"io"
	"net/http"
)

// Skin format constants to use with GetUserSkin
const (
	SkinFormatPNG = "png"
	SkinFormatSVG = "svg"
)

// Skin profile constants to use with GetUserSkin
const (
	SkinProfileFull   = "full"   // Renders the entire user skin.
	SkinProfileCenter = "center" // Renders a centered profile (face).
	SkinProfileRight  = "right"  // Renders a right-facing profile (face).
)

// Skin size constants to use with GetUserSkin (PNG only)
const (
	SkinSizeDefault = ""      // The default size.
	SkinSizeLarge   = "large" // A larger rendering.
	SkinSizeSmall   = "small" // A smaller rendering.
)

// GetUserSkin fetches the rendered skin image for a given user ID.
// It returns the raw image data as a byte slice.
// The 'size' parameter is only applied if the format is PNG.
func (c *Client) GetUserSkin(userID, format, profile, size string) ([]byte, error) {
	// 1. Build the base URL
	skinURL := fmt.Sprintf("https://wolfy.net/api/skin/render/user.%s?id=%s", format, userID)

	// 2. Add the profile parameter if it's not the default "full"
	if profile == SkinProfileCenter || profile == SkinProfileRight {
		skinURL += fmt.Sprintf("&profile=%s", profile)
	}

	// 3. Add the size parameter ONLY if the format is PNG and a size is specified
	if format == SkinFormatPNG && size != SkinSizeDefault {
		skinURL += fmt.Sprintf("&size=%s", size)
	}

	// 4. Create and execute the unauthenticated request
	req, err := http.NewRequest("GET", skinURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.defaultHeaders["User-Agent"])

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get skin, server responded with status: %s", resp.Status)
	}

	imageData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return imageData, nil
}

// GetPlayerInfo retrieves the detailed profile for a given player by their username or ID.
func (c *Client) GetPlayerInfo(usernameOrID string) (*PlayerInfoResponse, error) {
	path := fmt.Sprintf("/leaderboard/player/%s", usernameOrID)
	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var playerInfo PlayerInfoResponse
	if err := c.do(req, &playerInfo); err != nil {
		return nil, err
	}
	return &playerInfo, nil
}

// FindUserID finds a user by their exact username and returns their unique ID.
// This uses the player leaderboard endpoint for a more direct lookup.
func (c *Client) GetUserID(username string) (string, error) {
	// Use the GetPlayerInfo function which is designed for this lookup.
	playerInfo, err := c.GetPlayerInfo(username)
	if err != nil {
		// If there was an error (e.g., user not found, network issue), pass it on.
		return "", fmt.Errorf("could not find user '%s': %w", username, err)
	}

	// Success! Return the ID from the player data.
	return playerInfo.User.ID, nil
}
