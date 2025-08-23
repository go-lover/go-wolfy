package wolfyclient

// --- Request Structs (for application/x-www-form-urlencoded encoding) ---

// ChangeUsernameRequest is the request payload for changing a user's username.
type ChangeUsernameRequest struct {
	Username string `url:"username"`
}

// ChangeEmailRequest is the request payload for changing a user's email.
type ChangeEmailRequest struct {
	Email string `url:"email"`
}

// ChangePasswordRequest is the request payload for changing a user's password.
type ChangePasswordRequest struct {
	OldPassword string `url:"oldPass"`
	NewPassword string `url:"newPass"`
}

// --- Response Structs (for JSON decoding) ---

// MessageResponse is a generic response containing a single message string.
// It is used for many POST actions like logout, settings changes, etc.
type MessageResponse struct {
	Message string `json:"message"`
}

// AutocompleteUser represents a single user in the autocomplete search results.
type AutocompleteUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

// FriendsResponse contains a list of friend usernames.
type FriendsResponse struct {
	Friends []string `json:"friends"`
}

// PlayerInfoResponse is the top-level response from the /leaderboard/player endpoint.
type PlayerInfoResponse struct {
	User       PlayerUser       `json:"user"`
	Statistics PlayerStatistics `json:"statistics"`
	History    []interface{}    `json:"history"` // Kept as interface{} as the structure is unknown
}

// PlayerUser contains detailed information about a specific player.
type PlayerUser struct {
	ID          string        `json:"id"`
	Username    string        `json:"username"`
	CreatedAt   string        `json:"createdAt"` // Consider parsing to time.Time if needed
	Rank        int           `json:"rank"`
	XP          int           `json:"xp"`
	SkinVersion string        `json:"skinVersion"`
	Elo         int           `json:"elo"`
	Ranking     PlayerRanking `json:"ranking"`
}

// PlayerRanking holds the Elo ranking details for a user.
type PlayerRanking struct {
	Value   int     `json:"value"`
	Percent float64 `json:"percent"`
}

// PlayerStatistics holds the gameplay statistics for a user.
type PlayerStatistics struct {
	Individual PlayerIndividualStats `json:"individual"`
	Laurels    map[string]int        `json:"laurels"`
}

// PlayerIndividualStats contains specific win, kill, and word counts.
type PlayerIndividualStats struct {
	WinCount  int     `json:"winCount"`
	KillCount int     `json:"killCount"`
	WordAvg   float64 `json:"wordAvg"`
}
