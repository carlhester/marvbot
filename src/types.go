package main

import (
	"net/http"
)

type wsConfig struct {
	Token string
}

type Config struct {
	token   string
	baseURL string
	client  http.Client
}
