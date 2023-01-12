package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type Client struct {
	id   string
	room *Room
	ws   *websocket.Conn
}

type Message struct {
	Body  string `json:"body,omitempty"`
	Event string `json:"event,omitempty"`
}

func (c *Client) onMessage() {
	defer func() {
		c.room.Unregister <- c
		c.ws.Close()
	}()

	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		msg := &Message{}
		err := c.ws.ReadJSON(msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		c.room.Broadcast <- &Message{
			Body:  string(msg.Body),
			Event: string(msg.Event),
		}
	}
}

func createClient(room *Room, ws *websocket.Conn) {

	client := &Client{id: uuuid(), room: room, ws: ws}
	client.room.Register <- client

	fmt.Println("new client connected on room: ", client.room.ID)

	go client.onMessage()
}
