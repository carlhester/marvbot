package main

import (
	"net/http"
	"time"
)

type Channel struct {
	ID                   string        `json:"id"`
	LastMessageID        string        `json:"last_message_id"`
	LastPinTimestamp     time.Time     `json:"last_pin_timestamp"`
	Type                 int           `json:"type"`
	Name                 string        `json:"name"`
	Position             int           `json:"position"`
	ParentID             string        `json:"parent_id"`
	Topic                string        `json:"topic"`
	GuildID              string        `json:"guild_id"`
	PermissionOverwrites []interface{} `json:"permission_overwrites"`
	Nsfw                 bool          `json:"nsfw"`
	RateLimitPerUser     int           `json:"rate_limit_per_user"`
}

type Guild struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Icon        string   `json:"icon"`
	Owner       string   `json:"owner"`
	Permissions int      `json:"permissions"`
	Features    []string `json:"features"`
}

type User struct {
	ID            string      `json:"id"`
	Username      string      `json:"username"`
	Avatar        interface{} `json:"avatar"`
	Discriminator string      `json:"discriminator"`
	PublicFlags   int         `json:"public_flags"`
	Flags         int         `json:"flags"`
	Bot           bool        `json:"bot"`
	Email         interface{} `json:"email"`
	Verified      bool        `json:"verified"`
	Locale        string      `json:"locale"`
	MfaEnabled    bool        `json:"mfa_enabled"`
}

type wsConfig struct {
	Token string
}

type Payload struct {
	T  interface{}            `json:"t"`
	S  interface{}            `json:"s"`
	Op int                    `json:"op"`
	D  map[string]interface{} `json:"d"`
}

type Heartbeat struct {
	Op int         `json:"op"`
	D  interface{} `json:"d"`
}

type Identify struct {
	Op int          `json:"op"`
	D  IdentifyData `json:"d"`
}

type IdentifyData struct {
	Token      string            `json:"token"`
	Properties map[string]string `json:"properties"`
}
type Config struct {
	token   string
	baseURL string
	client  http.Client
}

type MessageAuthor struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Avatar        string `json:"avatar"`
	Discriminator string `json:"discriminator"`
	PublicFlags   int    `json:"public_flags"`
}

type Message struct {
	ID          string        `json:"id"`
	Type        int           `json:"type"`
	Content     string        `json:"content"`
	ChannelID   string        `json:"channel_id"`
	Author      MessageAuthor `json:"author"`
	Attachments []struct {
		ID       string `json:"id"`
		Filename string `json:"filename"`
		Size     int    `json:"size"`
		URL      string `json:"url"`
		ProxyURL string `json:"proxy_url"`
		Width    int    `json:"width"`
		Height   int    `json:"height"`
	} `json:"attachments"`
	Embeds          []interface{} `json:"embeds"`
	Mentions        []interface{} `json:"mentions"`
	MentionRoles    []interface{} `json:"mention_roles"`
	Pinned          bool          `json:"pinned"`
	MentionEveryone bool          `json:"mention_everyone"`
	Tts             bool          `json:"tts"`
	Timestamp       time.Time     `json:"timestamp"`
	EditedTimestamp interface{}   `json:"edited_timestamp"`
	Flags           int           `json:"flags"`
}

type Bot struct {
	uid    string
	guilds []Guild
}

type Gateway struct {
	Url string `json:"url"`
}

type Webhook struct {
	Name      string      `json:"name"`
	Type      int         `json:"type"`
	ChannelID string      `json:"channel_id"`
	Token     string      `json:"token"`
	Avatar    interface{} `json:"avatar"`
	GuildID   string      `json:"guild_id"`
	ID        string      `json:"id"`
	User      struct {
		Username      string `json:"username"`
		Discriminator string `json:"discriminator"`
		ID            string `json:"id"`
		Avatar        string `json:"avatar"`
	} `json:"user"`
}
