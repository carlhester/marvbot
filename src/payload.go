package main

type Payload struct {
	T  interface{}            `json:"t"`
	S  interface{}            `json:"s"`
	Op int                    `json:"op"`
	D  map[string]interface{} `json:"d"`
}
