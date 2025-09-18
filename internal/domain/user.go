package domain

import (
	"time"
)

type User struct {
	Activated            bool
	Avatar               string
	DisplayName          string
	Email                string
	IsDeleted            bool
	LastOnline           time.Time
	MessageBoardLastRead time.Time
	PasswordHash         string
	Rating               int
	Settings             string
	Skin                 string
	Achievements         []UserAchievement
}

type UserAchievement struct {
	User        *User
	Achievement *Achievement
}

type UserActivity struct {
	User    *User
	Ip      string
	Context string
	Status  string
	Device  string
}

type UserDiscordUser struct {
	User          *User
	DiscordUserId string
	Username      string
	Discriminator int
	Email         string
	RefreshToken  string
	ExpiresAt     time.Time
}

type UserFeedback struct {
	SourceId int
	TargetId int
	Endorsed bool
}

type UserIdentity struct {
	User        *User
	Fingerprint string
	UserAgent   string
	Count       int
}

type UserOrigin struct {
	User      *User
	Dominion  *Dominion
	IPAddress string
	Count     int
}

type UserOriginLookup struct {
	IPAddress    string
	ISP          string
	Organization string
	Country      string
	Region       string
	City         string
	VPN          bool
	Score        float64
	Data         string
}
