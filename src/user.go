package main

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
