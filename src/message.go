package main

import (
	"time"
)

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
