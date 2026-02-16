package wolfyclient

import "io"

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

// GetCurrentDrop retrieves details about the current featured item drop,
// including available packs and the cosmetic items within them.
func (c *Client) GetCurrentDrop() (*CurrentDrop, error) {
	req, err := c.newRequest("GET", "/drop", nil)
	if err != nil {
		return nil, err
	}

	var dropInfo CurrentDrop
	if err := c.do(req, &dropInfo); err != nil {
		return nil, err
	}
	return &dropInfo, nil
}

// GetDailyShopOffers retrieves the list of daily offers from the shop.
// This includes free items, and rotating skins and packs for coins and moons.
func (c *Client) GetDailyShopOffers() ([]DailyOfferSet, error) {
	req, err := c.newRequest("GET", "/shop/dailyOffers", nil)
	if err != nil {
		return nil, err
	}

	var dailyOffers []DailyOfferSet
	if err := c.do(req, &dailyOffers); err != nil {
		return nil, err
	}
	return dailyOffers, nil
}

// GetSubscriptionOffers retrieves the available Alpha subscription plans.
func (c *Client) GetSubscriptionOffers() ([]SubscriptionOffer, error) {
	req, err := c.newRequest("GET", "/shop/subscriptions/offers", nil)
	if err != nil {
		return nil, err
	}

	var subOffers []SubscriptionOffer
	if err := c.do(req, &subOffers); err != nil {
		return nil, err
	}
	return subOffers, nil
}

// GetMoonOffers retrieves the available Moon currency purchase options.
func (c *Client) GetMoonOffers() ([]MoonOffer, error) {
	req, err := c.newRequest("GET", "/shop/offers", nil)
	if err != nil {
		return nil, err
	}

	var moonOffers []MoonOffer
	if err := c.do(req, &moonOffers); err != nil {
		return nil, err
	}
	return moonOffers, nil
}
