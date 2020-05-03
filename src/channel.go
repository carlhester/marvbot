package main

import (
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
