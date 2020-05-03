package main

type Heartbeat struct {
	Op int         `json:"op"`
	D  interface{} `json:"d"`
}
