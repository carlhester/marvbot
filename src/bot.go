package main

type Bot struct {
	Uid    string
	Guilds []Guild
	Hooks  map[string]Webhook
}
