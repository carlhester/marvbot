package main

type Identify struct {
	Op int          `json:"op"`
	D  IdentifyData `json:"d"`
}

type IdentifyData struct {
	Token      string            `json:"token"`
	Properties map[string]string `json:"properties"`
}
