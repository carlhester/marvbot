package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	color "gopkg.in/gookit/color.v1"
)

func HandleWS(conf Config) {
	flag.Parse()
	log.SetFlags(0)

	token := os.Getenv("DISCORDTOKEN")
	config := wsConfig{Token: token}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	addr := conf.getGateway()
	u := fmt.Sprintf("%s/?v=6&encoding=json", addr)
	c, _, err := websocket.DefaultDialer.Dial(u, http.Header{"Authorization": []string{config.Token}})
	if err != nil {
		log.Fatal("dial error:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			color.Cyan.Println("reading...")
			_, message, err := c.ReadMessage()
			color.Cyan.Println("ReadMessage got a new message...")
			if err != nil {
				log.Println("read error: ", err)
				return
			}
			var p Payload
			err = json.Unmarshal(message, &p)
			if err != nil {
				log.Println("Unmarshal error: ", err)

				log.Fatal(err)
			}
			color.Cyan.Println("Payload: ")
			color.Green.Printf("%+v\n", p)
			color.Cyan.Println("Switching p.Op: ", p.Op, p.T, p.S)
			switch p.Op {
			case 0:
				_, ok := p.D["content"]
				if ok && p.T == "MESSAGE_CREATE" {
					color.Cyan.Println(p.Op, p.D["author"], p.D["content"])
					handlePayload(p)
				}
			case 7:
				return
			case 9:
				return
			case 10:
				go sendIdentify(config, c)
				go sendHeartbeat(p, c)
			default:
				color.Cyan.Println("No handler defined for p.Op: ", p.Op)
			}
		}
	}()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}

}

func handlePayload(p Payload) {
	color.Blue.Printf("%+v\n", p)
	authorData := p.D["author"].(map[string]interface{})
	channelID := p.D["channel_id"]
	content := p.D["content"]
	userName := (authorData["username"])
	color.Green.Printf("%s[%s]: %s\n", userName, channelID, content)
}

func sendIdentify(config wsConfig, c *websocket.Conn) {
	color.Green.Println("Sending Identify")

	properties := make(map[string]string)
	properties["$os"] = "linux"
	properties["$browser"] = "mybot"
	properties["$device"] = "mybot"

	identifyData := IdentifyData{
		Token:      config.Token,
		Properties: properties,
	}

	identify := Identify{
		Op: 2,
		D:  identifyData,
	}

	identifyJson, err := json.Marshal(identify)
	if err != nil {
		log.Println("error marshalling:", err)
	}
	err = c.WriteMessage(websocket.TextMessage, []byte(identifyJson))
	if err != nil {
		log.Println("error WriteMessage:", err)
	}
	color.Green.Println("Sent Identify", string(identifyJson))
}

func sendHeartbeat(p Payload, c *websocket.Conn) {
	color.Cyan.Println("Starting Heartbeat Cycle")
	for {
		color.Cyan.Println("Received opcode10, will send heartbeat after sleeping for", p.D["heartbeat_interval"])
		time.Sleep(time.Duration(p.D["heartbeat_interval"].(float64)) * time.Millisecond)
		hb := Heartbeat{Op: 1, D: p.S}
		hbJson, err := json.Marshal(hb)
		if err != nil {
			log.Println("error marshalling:", err)
		}
		color.Green.Println("Sending heartbeat: ", string(hbJson))
		err = c.WriteMessage(websocket.TextMessage, []byte(hbJson))
		if err != nil {
			log.Println("error:", err)
		}
	}
}
