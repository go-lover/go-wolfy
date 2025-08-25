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
	User       PlayerUser         `json:"user"`
	Statistics PlayerStatistics   `json:"statistics"`
	History    []GameHistoryEntry `json:"history"`
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

// GameSettings holds the specific rules and role composition for a game.
type GameSettings struct {
	Slots     int            `json:"slots"`
	Mayor     bool           `json:"mayor"`
	Roles     map[string]int `json:"roles"`
	Balancing int            `json:"balancing"`
}

// Game holds detailed information about a specific match instance.
type Game struct {
	ID          string       `json:"id"`
	InstanceID  string       `json:"instanceId"`
	Status      int          `json:"status"`
	PlayerCount int          `json:"playerCount"`
	Settings    GameSettings `json:"settings"`
	Private     bool         `json:"private"`
	Voice       bool         `json:"voice"`
	Serious     bool         `json:"serious"`
	Platform    string       `json:"platform"`
	Lang        string       `json:"lang"`
	CreatedAt   string       `json:"createdAt"`
	UpdatedAt   string       `json:"updatedAt"`
	NextID      interface{}  `json:"nextId"` // Can be null
	AdminID     string       `json:"adminId"`
}

// DeathReason provides details on how a player died in a game.
// Fields are optional as they depend on the type of death.
type DeathReason struct {
	Type      string   `json:"type"`
	DayNumber int      `json:"dayNumber"`
	VotersIDs []string `json:"votersIds,omitempty"`
	HunterID  string   `json:"hunterId,omitempty"`
	MayorID   string   `json:"mayorId,omitempty"`
	LoverID   string   `json:"loverId,omitempty"`
}

// GameHistoryEntry represents a single game played by the user in their history.
type GameHistoryEntry struct {
	Role        string       `json:"role"`
	Winner      bool         `json:"winner"`
	DeathReason *DeathReason `json:"deathReason"` // Pointer to handle null
	WordCount   int          `json:"wordCount"`
	KillCount   int          `json:"killCount"`
	XP          int          `json:"xp"`
	Elo         int          `json:"elo"`
	Lovers      bool         `json:"lovers"`
	Infected    bool         `json:"infected"`
	UserID      string       `json:"userId"`
	GameID      string       `json:"gameId"`
	CreatedAt   string       `json:"createdAt"`
	UpdatedAt   string       `json:"updatedAt"`
	Game        Game         `json:"game"`
}
