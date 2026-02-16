package wolfyclient

import "fmt"

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

// GetAccountDetails retrieves the detailed private profile for the currently authenticated user.
// This includes sensitive information such as email, currency, and account settings.
func (c *Client) GetAccountDetails() (*UserAccountInfo, error) {
	req, err := c.newRequest("GET", "/user", nil)
	if err != nil {
		return nil, err
	}

	var accountInfo UserAccountInfo
	if err := c.do(req, &accountInfo); err != nil {
		return nil, err
	}
	return &accountInfo, nil
}

// UpdateUsername changes the authenticated user's username.
func (c *Client) ChangeUsername(newUsername string) (*MessageResponse, error) {
	payload := ChangeUsernameRequest{
		Username: newUsername,
	}

	var resp MessageResponse
	err := c.doPostForm("/settings/username", payload, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateEmail changes the authenticated user's email address.
func (c *Client) ChangeEmail(newEmail string) (*MessageResponse, error) {
	payload := ChangeEmailRequest{
		Email: newEmail,
	}

	var resp MessageResponse
	err := c.doPostForm("/settings/email", payload, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdatePassword changes the authenticated user's password.
// It requires both the old and the new password.
func (c *Client) ChangePassword(oldPassword, newPassword string) (*MessageResponse, error) {
	payload := ChangePasswordRequest{
		OldPassword: oldPassword,
		NewPassword: newPassword,
	}

	var resp MessageResponse
	err := c.doPostForm("/settings/password", payload, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateSkinSlot changes the equipped cosmetic items for a specific skin slot.
// The 'updates' map should contain the skin parts to change, e.g., "top": SkinPart{ID:"002", Color:5}.
func (c *Client) UpdateSkinSlot(slotID string, updates map[string]SkinPart) (*UpdateSkinSlotResponse, error) {
	path := fmt.Sprintf("/slot/%s", slotID)

	var resp UpdateSkinSlotResponse
	err := c.doPutJSON(path, updates, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
