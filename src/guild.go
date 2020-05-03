package main

type Guild struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Icon        string   `json:"icon"`
	Owner       string   `json:"owner"`
	Permissions int      `json:"permissions"`
	Features    []string `json:"features"`
}
