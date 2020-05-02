package main

import (
	"encoding/json"
	"fmt"
	color "gopkg.in/gookit/color.v1"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Get the config details and init the bot
	config := Config{
		token:   os.Getenv("DISCORD_BOT_TOKEN"),
		baseURL: os.Getenv("DISCORD_API"),
		client:  http.Client{Timeout: 10 * time.Second},
	}

	bot := Bot{
		uid:    config.getBotUser().ID,
		guilds: config.getGuilds(),
	}

	knownHooks := config.getWebhookMap(bot)
	color.Green.Println(knownHooks)

	// Do websocket handling
	HandleWS(config)
}

// getBotUser returns the userID for the Bot
func (conf Config) getBotUser() User {
	// Build the appropriate request
	fullURL := conf.baseURL + "/users/@me"
	request, _ := http.NewRequest("GET", fullURL, nil)
	request.Header.Add("Authorization", conf.token)

	// Make the request
	resp, err := conf.client.Do(request)
	if err != nil {
		log.Printf("Error making request %+v", err)
	}
	defer resp.Body.Close()

	// Parse the response into the data we return
	var user User
	bodyJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("%+v", err)
	}
	err = json.Unmarshal([]byte(bodyJson), &user)
	if err != nil {
		fmt.Errorf("%+v", err)
	}
	return user
}

// getGateway returns the current API gateway published by discord
func (conf Config) getGateway() string {
	// Build the appropriate request
	fullURL := conf.baseURL + "/gateway"
	request, _ := http.NewRequest("GET", fullURL, nil)

	// Make the request
	resp, err := conf.client.Do(request)
	if err != nil {
		log.Printf("Error making request %+v", err)
	}
	defer resp.Body.Close()

	// Parse the response into the data we return
	var gateway Gateway
	bodyJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("%+v", err)
	}
	err = json.Unmarshal([]byte(bodyJson), &gateway)
	if err != nil {
		fmt.Errorf("%+v", err)
	}
	return string(gateway.Url)
}

// getGuilds returns the guilds that our tokens provide access to
func (conf Config) getGuilds() []Guild {
	// Build the appropriate request
	fullURL := conf.baseURL + "/users/@me/guilds"
	request, _ := http.NewRequest("GET", fullURL, nil)
	request.Header.Add("Authorization", conf.token)

	// Make the request
	resp, err := conf.client.Do(request)
	if err != nil {
		fmt.Errorf("%+v", err)
	}
	defer resp.Body.Close()

	// Parse the response into the data we return
	var guild []Guild
	bodyJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("%+v", err)
	}
	err = json.Unmarshal([]byte(bodyJson), &guild)
	if err != nil {
		fmt.Errorf("%+v", err)
	}
	return guild

}
