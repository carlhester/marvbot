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
	c := Config{
		token:   os.Getenv("DISCORD_BOT_TOKEN"),
		baseURL: os.Getenv("DISCORD_API"),
		client:  http.Client{Timeout: 10 * time.Second},
	}

	// init
	bot := Bot{
		uid:    c.getBotUser().ID,
		guilds: c.getBotGuilds(),
	}

	knownHooks := getWebhookMap(bot, c)

	color.Green.Println(knownHooks)

	HandleWS(c)
}

func getWebhookMap(b Bot, c Config) map[string]string {
	var knownHooks = make(map[string]string)
	for _, g := range b.guilds {
		hooks := c.getWebhooksForGuild(g)
		for _, hook := range hooks {
			knownHooks[string(hook.ChannelID)] = string(hook.Token)
		}
	}
	return knownHooks
}

func (c Config) getWebhooksForGuild(g Guild) []Webhook {
	// Build the appropriate request
	fullURL := c.baseURL + "/guilds/" + string(g.Id) + "/webhooks"
	request, _ := http.NewRequest("GET", fullURL, nil)
	request.Header.Add("Authorization", c.token)

	// Make the request
	resp, err := c.client.Do(request)
	if err != nil {
		log.Printf("Error making request %+v", err)
	}
	defer resp.Body.Close()

	// Parse the response into the data we return
	var wh []Webhook
	bodyJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("%+v", err)
	}
	err = json.Unmarshal([]byte(bodyJson), &wh)
	if err != nil {
		fmt.Errorf("%+v", err)
	}
	return wh

}

func (c Config) getBotUser() User {
	// Build the appropriate request
	fullURL := c.baseURL + "/users/@me"
	request, _ := http.NewRequest("GET", fullURL, nil)
	request.Header.Add("Authorization", c.token)

	// Make the request
	resp, err := c.client.Do(request)
	if err != nil {
		log.Printf("Error making request %+v", err)
	}
	defer resp.Body.Close()

	// Parse the response into the data we return
	var u User
	bodyJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("%+v", err)
	}
	err = json.Unmarshal([]byte(bodyJson), &u)
	if err != nil {
		fmt.Errorf("%+v", err)
	}
	return u
}

func (c Config) getGateway() string {
	// Build the appropriate request
	fullURL := c.baseURL + "/gateway"
	request, _ := http.NewRequest("GET", fullURL, nil)
	//request.Header.Add("Authorization", c.token)

	// Make the request
	resp, err := c.client.Do(request)
	if err != nil {
		log.Printf("Error making request %+v", err)
	}
	defer resp.Body.Close()

	// Parse the response into the data we return
	var gw Gateway
	bodyJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("%+v", err)
	}
	err = json.Unmarshal([]byte(bodyJson), &gw)
	if err != nil {
		fmt.Errorf("%+v", err)
	}
	fmt.Println(string(gw.Url))
	return string(gw.Url)
}

func (c Config) getBotGuilds() []Guild {
	// Build the appropriate request
	fullURL := c.baseURL + "/users/@me/guilds"
	request, _ := http.NewRequest("GET", fullURL, nil)
	request.Header.Add("Authorization", c.token)

	// Make the request
	resp, err := c.client.Do(request)
	if err != nil {
		fmt.Errorf("%+v", err)
	}
	defer resp.Body.Close()

	// Parse the response into the data we return
	var g []Guild
	bodyJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("%+v", err)
	}
	err = json.Unmarshal([]byte(bodyJson), &g)
	if err != nil {
		fmt.Errorf("%+v", err)
	}
	return g

}
