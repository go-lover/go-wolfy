package wolfyclient

import (
	"fmt"
)

// Logout invalidates the current user's session on the server.
func (c *Client) Logout() (*MessageResponse, error) {
	var resp MessageResponse
	err := c.doPostForm("/auth/logout", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSelfInfo retrieves the detailed profile for the currently authenticated user.
func (c *Client) GetSelfInfo() (*PlayerInfoResponse, error) {
	req, err := c.newRequest("GET", "/leaderboard/player/self", nil)
	if err != nil {
		return nil, err
	}

	var playerInfo PlayerInfoResponse
	if err := c.do(req, &playerInfo); err != nil {
		return nil, err
	}
	return &playerInfo, nil
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

// GetFriends retrieves the friend list of the authenticated user.
func (c *Client) GetFriendList() ([]string, error) {
	req, err := c.newRequest("GET", "/social/friends", nil)
	if err != nil {
		return nil, err
	}

	// We expect a direct slice of strings, not a struct.
	var friendIDs []string
	if err := c.do(req, &friendIDs); err != nil {
		return nil, err
	}
	return friendIDs, nil
}

// AddFriend sends a friend request to the specified user ID.
// This function adds specific headers required for this endpoint.
func (c *Client) AddFriend(userID string) (*MessageResponse, error) {
	path := fmt.Sprintf("/social/add/%s", userID)

	// Create a new POST request with an empty body (nil).
	req, err := c.newRequest("POST", path, nil)
	if err != nil {
		return nil, err
	}

	// Add the special headers required for this action.
	// This will override the default "Referer" for this one request.
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", "https://wolfy.net/fr/play")

	// Execute the request and decode the JSON response into our struct.
	var resp MessageResponse
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveFriend sends a request to remove the specified user from the friend list.
func (c *Client) RemoveFriend(userID string) (*MessageResponse, error) {
	path := fmt.Sprintf("/social/remove/%s", userID)

	// Create a new POST request with an empty body.
	req, err := c.newRequest("POST", path, nil)
	if err != nil {
		return nil, err
	}

	// Add the specific headers required for this action, based on the browser's request.
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Referer", "https://wolfy.net/fr/play")

	// Execute the request and decode the JSON response.
	var resp MessageResponse
	if err := c.do(req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
