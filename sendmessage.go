package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// sendMessage sends text content to the specific channel
func (conf Config) sendMessage(bot Bot, channelID string, content string) error {
	hook, ok := bot.Hooks[channelID]
	if ok {
		fmt.Println(hook)
	}
	// Build the appropriate request
	fullURL := conf.baseURL + "/webhooks/" + string(hook.ID) + "/" + string(hook.Token)

	requestBody, err := json.Marshal(map[string]string{
		"content":       content,
		"Authorization": conf.token,
	})
	if err != nil {
		fmt.Errorf("Error Marshaling requestBody: %+v", err)
	}
	//request.Header.Add("Authorization", conf.token)

	resp, err := http.Post(fullURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Errorf("Error making Post: %+v", err)
		defer resp.Body.Close()
	}
	return nil
}
