package models

// Identity represents an identity data structure
type Identity struct {
	User  User    `json:"user"`
	Guild []Guild `json:"guilds"`
	Token string  `json:"token"`
}
type SyncData struct {
	Ver        int
	InSync     int
	TechLevels map[int]TechLevel
}

// TechLevel represents a tech level data structure
type TechLevel struct {
	Level int   `json:"level"`
	Ts    int64 `json:"ts"`
}

// Guild represents a guild data structure
type Guild struct {
	URL  string `json:"url"`
	ID   string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

// User represents a user data structure
type User struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
	AvatarURL     string `json:"avatarUrl"`
}

// CorpData represents corporation data structure
type CorpData struct {
	Members    []CorpMember `json:"members"`
	Roles      []CorpRole   `json:"roles"`
	FilterID   string       `json:"filterId"`
	FilterName string       `json:"filterName"`
}

// CorpMember represents a member of a corporation.
type CorpMember struct {
	Name         string        `json:"name"`
	UserID       string        `json:"userId"`
	ClientUserID string        `json:"clientUserId"`
	Avatar       string        `json:"avatar"`
	Tech         map[int][]int `json:"tech"`
	AvatarURL    string        `json:"avatarUrl"`
	TimeZone     string        `json:"timeZone"`
	LocalTime    string        `json:"localTime"`
	ZoneOffset   int           `json:"zoneOffset"`
	AfkFor       string        `json:"afkFor"`
	AfkWhen      int           `json:"afkWhen"`
}

// CorpRole represents a corporation role data structure
type CorpRole struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type StorageData struct {
	Ident        *Identity `json:"ident"`
	UserData     *SyncData `json:"userData"`
	Refresh      int64     `json:"refresh"`
	TokenRefresh int64     `json:"tokenRefresh"`
}
