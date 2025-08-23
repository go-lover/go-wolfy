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

// CollectDailyItem attempts to claim the free daily item from the shop.
func (c *Client) CollectDailyItem() (string, error) {
	req, err := c.newRequest("POST", "/shop/collect/free", nil)
	if err != nil {
		return "", err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
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
