package wolfyclient

// GetSkinCatalog retrieves the master catalog of all available cosmetic items in the game.
// This is an authenticated call and returns a slice of all skin elements.
func (c *Client) GetSkinCatalog() ([]SkinElement, error) {
	req, err := c.newRequest("GET", "/skin/elements", nil)
	if err != nil {
		return nil, err
	}

	var skinCatalog []SkinElement
	if err := c.do(req, &skinCatalog); err != nil {
		return nil, err
	}
	return skinCatalog, nil
}
