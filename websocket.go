package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	color "gopkg.in/gookit/color.v1"
)

func HandleWebSocket(conf Config) {
	wsConfig := wsConfig{Token: os.Getenv("DISCORDTOKEN")}
	addr, err := conf.getGateway()
	if err != nil {
		fmt.Errorf("Unable to get Discord gateway. %+v", err)
	}
	url := fmt.Sprintf("%s/?v=6&encoding=json", addr)
	wsConn, _, err := websocket.DefaultDialer.Dial(url, http.Header{"Authorization": []string{wsConfig.Token}})
	if err != nil {
		log.Fatal("dial error:", err)
	}
	defer wsConn.Close()

	interruptChannel := make(chan os.Signal, 1)
	signal.Notify(interruptChannel, os.Interrupt)
	doneChannel := make(chan struct{})

	go func() {
		defer close(doneChannel)
		for {
			color.Cyan.Println("reading...")
			_, message, err := wsConn.ReadMessage()
			color.Cyan.Println("ReadMessage got a new message...")
			if err != nil {
				log.Println("read error: ", err)
				return
			}
			var payload Payload
			err = json.Unmarshal(message, &payload)
			if err != nil {
				log.Println("Unmarshal error: ", err)

				log.Fatal(err)
			}
			color.Cyan.Println("Payload: ")
			color.Green.Printf("%+v\n", payload)
			color.Cyan.Println("Switching payload.Op: ", payload.Op, payload.T, payload.S)
			switch payload.Op {
			case 0:
				_, ok := payload.D["content"]
				if ok && payload.T == "MESSAGE_CREATE" {
					color.Cyan.Println(payload.Op, payload.D["author"], payload.D["content"])
					handleWSPayload(payload)
				}
			case 7:
				return
			case 9:
				return
			case 10:
				go sendWSIdentify(wsConfig, wsConn)
				go sendWSHeartbeat(payload, wsConn)
			default:
				color.Cyan.Println("No handler defined for payload.Op: ", payload.Op)
			}
		}
	}()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-doneChannel:
			return
		case <-interruptChannel:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := wsConn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-doneChannel:
			case <-time.After(time.Second):
			}
			return
		}
	}

}

func handleWSPayload(payload Payload) {
	color.Blue.Printf("%+v\n", payload)
	authorData := payload.D["author"].(map[string]interface{})
	channelID := payload.D["channel_id"]
	content := payload.D["content"]
	userName := (authorData["username"])
	color.Green.Printf("%s[%s]: %s\n", userName, channelID, content)
}

func sendWSIdentify(wsConfig wsConfig, wsConn *websocket.Conn) {
	color.Green.Println("Sending Identify")

	properties := make(map[string]string)
	properties["$os"] = "linux"
	properties["$browser"] = "mybot"
	properties["$device"] = "mybot"

	identifyData := IdentifyData{
		Token:      wsConfig.Token,
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
	err = wsConn.WriteMessage(websocket.TextMessage, []byte(identifyJson))
	if err != nil {
		log.Println("error WriteMessage:", err)
	}
	color.Green.Println("Sent Identify", string(identifyJson))
}

func sendWSHeartbeat(payload Payload, wsConn *websocket.Conn) {
	color.Cyan.Println("Starting Heartbeat Cycle")
	for {
		color.Cyan.Println("Received opcode10, will send heartbeat after sleeping for", payload.D["heartbeat_interval"])
		time.Sleep(time.Duration(payload.D["heartbeat_interval"].(float64)) * time.Millisecond)
		hb := Heartbeat{Op: 1, D: payload.S}
		hbJson, err := json.Marshal(hb)
		if err != nil {
			log.Println("error marshalling:", err)
		}
		color.Green.Println("Sending heartbeat: ", string(hbJson))
		err = wsConn.WriteMessage(websocket.TextMessage, []byte(hbJson))
		if err != nil {
			log.Println("error:", err)
		}
	}
}

// getGateway returns the current API gateway published by discord
func (conf Config) getGateway() (string, error) {
	// Build the appropriate request
	fullURL := conf.baseURL + "/gateway"
	request, _ := http.NewRequest("GET", fullURL, nil)

	// Make the request
	resp, err := conf.client.Do(request)
	if err != nil {
		log.Printf("Error making request %+v", err)
		return "", err
	}
	defer resp.Body.Close()

	// Parse the response into the data we return
	var gateway Gateway
	bodyJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("%+v", err)
		return "", err
	}
	err = json.Unmarshal([]byte(bodyJson), &gateway)
	if err != nil {
		fmt.Errorf("%+v", err)
		return "", err
	}
	return string(gateway.Url), nil
}
