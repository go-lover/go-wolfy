package wolfyclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	baseURL = "https://wolfy.net/api/"
)

// Client is the main API client for the Wolfy.net API.
type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
	// ADD THIS: A map to hold headers that will be sent with every request.
	defaultHeaders map[string]string
}

// NewClient creates and new, authenticated API client.
// It immediately checks if the provided authToken is valid by making a test API call.
func NewClient(authToken string) (*Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	apiBaseURL, _ := url.Parse(baseURL)

	client := &Client{
		baseURL: apiBaseURL,
		httpClient: &http.Client{
			Jar: jar,
		},
		defaultHeaders: map[string]string{
			"User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:142.0) Gecko/20100101 Firefox/142.0",
			"Referer":    "https://wolfy.net/fr/shop",
		},
	}

	client.SetSessionCookie(authToken)

	// --- NEW VALIDATION STEP ---
	// We test the token by making a lightweight, authenticated API call.
	// We use the blank identifier '_' because we don't need the player data here,
	// we only care if the call produces an error.
	_, err = client.GetSelfInfo()
	if err != nil {
		// If the call fails, it's highly likely the token is invalid or expired.
		// We wrap the original error to provide more context.
		return nil, fmt.Errorf("invalid token: authentication check failed: %w", err)
	}
	// --- END VALIDATION STEP ---

	// If we reach here, the token is valid and the client is ready to use.
	return client, nil
}

func (c *Client) SetSessionCookie(token string) {
	cookie := &http.Cookie{
		Name:  "wolfy",
		Value: token,
	}
	c.httpClient.Jar.SetCookies(c.baseURL, []*http.Cookie{cookie})
}

// --- Internal Helper Methods ---

func (c *Client) newRequest(method, path string, body io.Reader) (*http.Request, error) {
	rel, err := url.Parse(strings.TrimPrefix(path, "/"))
	if err != nil {
		return nil, err
	}
	fullURL := c.baseURL.ResolveReference(rel)

	req, err := http.NewRequest(method, fullURL.String(), body)
	if err != nil {
		return nil, err
	}

	// NEW LOOP: This is the magic. It applies all default headers to the request.
	for key, value := range c.defaultHeaders {
		req.Header.Set(key, value)
	}

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// --- NEW ROBUSTNESS CHECK ---
	// First, check if the request was successful.
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		// If the status is not OK, read the body as plain text for a better error message.
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("api request failed with status %s: %s", resp.Status, string(bodyBytes))
	}

	// If a struct was provided to decode into...
	if v != nil {
		// ...only attempt to decode if the server explicitly says it's sending JSON.
		contentType := resp.Header.Get("Content-Type")
		if strings.Contains(contentType, "application/json") {
			return json.NewDecoder(resp.Body).Decode(v)
		}
		// If it's not JSON, we don't attempt to decode, preventing a crash.
	}

	return nil
}

func (c *Client) doPostForm(path string, payload, v interface{}) error {
	var bodyReader io.Reader
	if payload != nil {
		formValues, err := query.Values(payload)
		if err != nil {
			return err
		}
		bodyReader = strings.NewReader(formValues.Encode())
	}

	req, err := c.newRequest("POST", path, bodyReader)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return c.do(req, v)
}

// doPutJSON is a helper for making PUT requests with a JSON-encoded body.
func (c *Client) doPutJSON(path string, payload, v interface{}) error {
	var bodyReader io.Reader
	if payload != nil {
		// Marshal the payload struct into a JSON byte slice
		jsonBytes, err := json.Marshal(payload)
		if err != nil {
			return err
		}
		// Create a reader from the byte slice
		bodyReader = bytes.NewBuffer(jsonBytes)
	}

	req, err := c.newRequest("PUT", path, bodyReader)
	if err != nil {
		return err
	}

	// Set the correct content type for JSON data
	req.Header.Set("Content-Type", "application/json")

	return c.do(req, v)
}
