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

	guilds := config.getGuilds()

	bot := Bot{
		Uid:    config.getBotUser().ID,
		Guilds: guilds,
		Hooks:  config.getWebhookMap(guilds),
	}

	color.Blue.Printf("Uid: %+v\n", bot.Uid)
	color.Blue.Printf("Guilds: %+v\n", bot.Guilds)
	color.Blue.Printf("Hooks: %+v\n", bot.Hooks)
	//knownHooks := config.getWebhookMap(bot)
	//color.Green.Printf("%+v\n", knownHooks)

	// Do websocket handling
	//HandleWebSocket(config)
	channelID := "xxx"
	_ = config.sendMessage(bot, channelID, "I'm alive")
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
