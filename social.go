package wolfyclient

import "fmt"

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

// GetFriendLeaderboard retrieves the leaderboard of the authenticated user's friends,
// returning a slice of users with their rank and summary information.
func (c *Client) GetFriendLeaderboard() ([]LeaderboardEntry, error) {
	req, err := c.newRequest("GET", "/leaderboard", nil)
	if err != nil {
		return nil, err
	}

	var friendLeaderboard []LeaderboardEntry
	if err := c.do(req, &friendLeaderboard); err != nil {
		return nil, err
	}
	return friendLeaderboard, nil
}
