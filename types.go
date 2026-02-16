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

// PlayerInfoResponse is the top-level response from the /leaderboard/player endpoint.
type PlayerInfoResponse struct {
	User       PlayerUser         `json:"user"`
	IsFriend   bool               `json:"isFriend"`
	Statistics PlayerStatistics   `json:"statistics"`
	History    []GameHistoryEntry `json:"history"`
}

// PlayerUser contains detailed information about a specific player.
type PlayerUser struct {
	ProfilePicture   string        `json:"profilePicture"`
	ID               string        `json:"id"`
	Username         string        `json:"username"`
	Rank             int           `json:"rank"`
	XP               int           `json:"xp"`
	SlotID           string        `json:"slotId"`
	SkinVersion      string        `json:"skinVersion"`
	Elo              int           `json:"elo"`
	CreatedAt        string        `json:"createdAt"`
	MonthsSubscribed int           `json:"monthsSubscribed"`
	GamePlayed       int           `json:"gamePlayed"`
	Ranking          PlayerRanking `json:"ranking"`
}

// PlayerRanking holds the Elo ranking details for a user.
type PlayerRanking struct {
	Value   int     `json:"value"`
	Percent float64 `json:"percent"`
}

// RoleStats contains win rate and advanced statistics for a specific role.
type RoleStats struct {
	ID            string             `json:"id"`
	WinRate       float64            `json:"winRate"`
	AdvancedStats map[string]float64 `json:"advancedStats"`
}

// GameTypeAdvancedStats contains overall advanced stats for a game alignment (innocent/threat).
type GameTypeAdvancedStats struct {
	Inactivity     float64 `json:"inactivity"`
	DaysAlive      float64 `json:"daysAlive"`
	Mayor          float64 `json:"mayor"`
	GoodVote       float64 `json:"goodVote,omitempty"`
	InnocentKilled float64 `json:"innocentKilled,omitempty"`
}

// GameTypeStats contains the overall win rate and stats for an alignment.
type GameTypeStats struct {
	ID            string                `json:"id"`
	WinRate       float64               `json:"winRate"`
	AdvancedStats GameTypeAdvancedStats `json:"advancedStats"`
}

// OverallGameStats holds the statistics broken down by alignment (innocent vs. threat).
type OverallGameStats struct {
	Innocent GameTypeStats `json:"innocent"`
	Threat   GameTypeStats `json:"threat"`
}

// PlayerStatistics holds the gameplay statistics for a user.
type PlayerStatistics struct {
	Laurels    map[string]int        `json:"laurels"`
	Individual PlayerIndividualStats `json:"individual"`
	Roles      []RoleStats           `json:"roles"`
	Game       OverallGameStats      `json:"game"`
}

// PlayerIndividualStats contains specific win, kill, and word counts.
type PlayerIndividualStats struct {
	Moonpass  int     `json:"moonpass"`
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

// SkinPart defines a single component of a skin, like eyes or hair.
type SkinPart struct {
	ID    string `json:"id"`
	Color int    `json:"color"`
}

// Skin provides a detailed breakdown of all components of a user's skin.
type Skin struct {
	Eyes      SkinPart `json:"eyes"`
	Face      SkinPart `json:"face"`
	Hair      SkinPart `json:"hair"`
	Nose      SkinPart `json:"nose"`
	Top       SkinPart `json:"top"`
	Bottom    SkinPart `json:"bottom"`
	Shoes     SkinPart `json:"shoes"`
	Tombstone SkinPart `json:"tombstone"`
	Glasses   SkinPart `json:"glasses"`
}

// Slot represents a single skin slot, which can be locked or unlocked.
type Slot struct {
	Unlocked    bool   `json:"unlocked"`
	ID          string `json:"id"`
	OfferID     string `json:"offerId,omitempty"` // Only in unlocked slots
	SkinVersion string `json:"skinVersion,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
	UserID      string `json:"userId,omitempty"`
	Skin        *Skin  `json:"skin,omitempty"` // Pointer to handle null
	Equiped     bool   `json:"equiped,omitempty"`
	Price       int    `json:"price,omitempty"` // Only in locked slots
	Currency    string `json:"currency,omitempty"`
	Alpha       bool   `json:"alpha,omitempty"`
}

// TokenInfo contains details about the user's current session token.
type TokenInfo struct {
	ID        string      `json:"id"`
	TwoFactor interface{} `json:"twoFactor"`
}

// UserAccountInfo is the top-level response from the /user endpoint,
// containing detailed private information for the authenticated user.
type UserAccountInfo struct {
	ID                  string      `json:"id"`
	Username            string      `json:"username"`
	Email               string      `json:"email"`
	TwitterID           interface{} `json:"twitterId"`
	FacebookID          interface{} `json:"facebookId"`
	GoogleID            interface{} `json:"googleId"`
	DiscordID           interface{} `json:"discordId"`
	AppleID             interface{} `json:"appleId"`
	ProfilePicture      string      `json:"profilePicture"`
	XP                  int         `json:"xp"`
	Elo                 int         `json:"elo"`
	Coins               int         `json:"coins"`
	Moons               int         `json:"moons"`
	Rank                int         `json:"rank"`
	SkinVersion         string      `json:"skinVersion"`
	SkinIndex           int         `json:"skinIndex"`
	AnonymousSkinIndex  int         `json:"anonymousSkinIndex"`
	SlotID              string      `json:"slotId"`
	AnonymousSlotID     interface{} `json:"anonymousSlotId"`
	AllowFriendRequests bool        `json:"allowFriendRequests"`
	AllowGroupRequests  bool        `json:"allowGroupRequests"`
	AllowNewsletter     bool        `json:"allowNewsletter"`
	Nickname            interface{} `json:"nickname"`
	Confirmed           bool        `json:"confirmed"`
	DiscountEndAt       interface{} `json:"discountEndAt"`
	TwoFactorSecret     bool        `json:"twoFactorSecret"`
	Lang                string      `json:"lang"`
	BanEnd              interface{} `json:"ban_end"`
	ReasonBan           interface{} `json:"reason_ban"`
	NeedRename          bool        `json:"needRename"`
	Banned              bool        `json:"banned"`
	FriendsVisibility   string      `json:"friendsVisibility"`
	AlphaLegacy         bool        `json:"alphaLegacy"`
	Password            bool        `json:"password"`
	Token               TokenInfo   `json:"token"`
	Slots               []Slot      `json:"slots"`
	Skin                Skin        `json:"skin"`
	Features            []string    `json:"features"`
	Subscription        interface{} `json:"subscription"`
}

// LeaderboardEntry represents a single user's summary on the main leaderboard.
type LeaderboardEntry struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	XP          int    `json:"xp"`
	SlotID      string `json:"slotId"`
	SkinVersion string `json:"skinVersion"`
	Rank        int    `json:"rank"`
	Elo         int    `json:"elo"`
	GamePlayed  int    `json:"gamePlayed"`
	IsFriend    bool   `json:"isFriend"`
}

// Disposition defines the positioning and scale of a skin element.
// It is a pointer in the parent struct because it can be null.
type Disposition struct {
	Y     float64 `json:"y"`
	Scale float64 `json:"scale"`
	X     float64 `json:"x,omitempty"` // omitempty because it's not always present
}

// SkinLayer represents a single graphical layer for a skin element.
type SkinLayer struct {
	ID int `json:"id"`
}

// SkinElement is the top-level response for a single item from the /skin/elements endpoint.
// It represents one available cosmetic item in the game's master catalog.
type SkinElement struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Type        string       `json:"type"`
	Access      string       `json:"access"`
	Rarity      string       `json:"rarity"`
	Level       int          `json:"level"`
	Price       int          `json:"price"`
	Colors      [][]string   `json:"colors"`
	New         bool         `json:"new"`
	Disposition *Disposition `json:"disposition"` // Pointer to handle null value
	Currency    string       `json:"currency"`
	SmallPet    interface{}  `json:"smallPet"` // Type is unknown, can be null
	SkinLayers  []SkinLayer  `json:"skinLayers"`
	Bought      bool         `json:"bought"`
}

// PackSkinElementLink contains metadata linking a skin element to a drop pack.
type PackSkinElementLink struct {
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
	SkinPackID    int    `json:"skinPackId"`
	SkinElementID string `json:"skinElementId"`
}

// DropSkinElement represents a skin element as part of a drop pack.
// It's similar to SkinElement but with some different fields.
type DropSkinElement struct {
	ID              string              `json:"id"`
	Name            string              `json:"name"`
	Type            string              `json:"type"`
	Access          string              `json:"access"`
	Rarity          string              `json:"rarity"`
	Level           int                 `json:"level"`
	Currency        string              `json:"currency"`
	Price           int                 `json:"price"`
	Colors          [][]string          `json:"colors"`
	New             bool                `json:"new"`
	Disposition     *Disposition        `json:"disposition"` // Re-using from SkinElement
	SmallPet        bool                `json:"smallPet"`
	CreatedAt       string              `json:"createdAt"`
	UpdatedAt       string              `json:"updatedAt"`
	PackSkinElement PackSkinElementLink `json:"PackSkinElement"`
}

// PreviewSkinElement represents a skin item used for previewing a pack.
type PreviewSkinElement struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Type        string       `json:"type"`
	Access      string       `json:"access"`
	Rarity      string       `json:"rarity"`
	Level       int          `json:"level"`
	Currency    string       `json:"currency"`
	Price       int          `json:"price"`
	New         bool         `json:"new"`
	Disposition *Disposition `json:"disposition"`
	SmallPet    interface{}  `json:"smallPet"` // Can be null
	CreatedAt   string       `json:"createdAt"`
	UpdatedAt   string       `json:"updatedAt"`
}

// DropPack represents a single bundle or pack within the current drop.
type DropPack struct {
	ID              int                      `json:"id"`
	Name            string                   `json:"name"`
	Colors          []map[string]interface{} `json:"colors"` // Flexible for varied keys
	Price           int                      `json:"price"`
	Rarity          string                   `json:"rarity"`
	Currency        string                   `json:"currency"`
	SkinElements    []DropSkinElement        `json:"SkinElements"`
	PreviewElements []PreviewSkinElement     `json:"previewElements"`
	Collected       bool                     `json:"collected"`
}

// CurrentDrop is the top-level response from the /drop endpoint.
type CurrentDrop struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Start     string     `json:"start"`
	End       string     `json:"end"`
	CreatedAt string     `json:"createdAt"`
	UpdatedAt string     `json:"updatedAt"`
	Packs     []DropPack `json:"packs"`
}

// OfferPack represents a bundle of items within a daily offer.
type OfferPack struct {
	Price        int           `json:"price"`
	Currency     string        `json:"currency"`
	ID           int           `json:"id"`
	Rarity       string        `json:"rarity"`
	SkinElements []SkinElement `json:"SkinElements"` // Reusing the existing SkinElement
}

// OfferElement represents a single slot in the daily offers, which can contain a skin or a pack.
type OfferElement struct {
	Category  string       `json:"category"`
	Moons     int          `json:"moons"`
	Coins     int          `json:"coins"`
	Pack      *OfferPack   `json:"pack"` // Pointer to handle null
	Skin      *SkinElement `json:"skin"` // Pointer to handle null
	Collected bool         `json:"collected,omitempty"`
}

// OfferElements is a structured representation of all the different offer categories for a day.
type OfferElements struct {
	MoonsUltraHigh OfferElement `json:"moonsUltraHigh"`
	CollectionHigh OfferElement `json:"collectionHigh"`
	MoonsLow       OfferElement `json:"moonsLow"`
	CoinsLow       OfferElement `json:"coinsLow"`
	CoinsHigh      OfferElement `json:"coinsHigh"`
	MoonsHigh      OfferElement `json:"moonsHigh"`
	MoonsMedium    OfferElement `json:"moonsMedium"`
	Premium        OfferElement `json:"premium"`
	CollectionLow  OfferElement `json:"collectionLow"`
	Free           OfferElement `json:"free"`
}

// DailyOfferSet represents the complete set of offers available for a single day.
type DailyOfferSet struct {
	ID       int           `json:"id"`
	End      string        `json:"end"`
	Elements OfferElements `json:"elements"`
}

// SubscriptionOffer represents a single purchasable Alpha subscription plan.
type SubscriptionOffer struct {
	ID            string  `json:"id"`
	Stripe        string  `json:"stripe"`
	Price         float64 `json:"price"`
	Currency      string  `json:"currency"`
	Interval      string  `json:"interval"`
	IntervalCount int     `json:"intervalCount"`
	SavedLabel    int     `json:"savedLabel,omitempty"`
	Badge         string  `json:"badge"`
	MostPopular   bool    `json:"mostPopular,omitempty"`
}

// MoonOffer represents a single purchasable pack of Moon currency.
type MoonOffer struct {
	ID                  string  `json:"id"`
	Moons               int     `json:"moons"`
	Bonus               int     `json:"bonus"`
	Tier                int     `json:"tier"`
	Discount            float64 `json:"discount"`
	Img                 string  `json:"img"`
	Stripe              string  `json:"stripe"`
	Price               float64 `json:"price"`
	Currency            string  `json:"currency"`
	IDNewPlayerDiscount string  `json:"idNewPlayerDiscount,omitempty"`
	NewPlayerDiscount   float64 `json:"newPlayerDiscount,omitempty"`
	Tag                 string  `json:"tag,omitempty"`
}

// UpdateSkinSlotResponse is the response received after successfully updating a skin slot.
type UpdateSkinSlotResponse struct {
	Slots   []Slot `json:"slots"`
	Skin    Skin   `json:"skin"`
	SlotID  string `json:"slotId"`
	Version string `json:"version"`
	Coins   int    `json:"coins"`
	Moons   int    `json:"moons"`
}
