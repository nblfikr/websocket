package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsocket(w http.ResponseWriter, r *http.Request) {

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	room := r.URL.Query().Get("r")
	createClient(prepare(room), ws)

	// for each connection clien on connection event call
	// 1. create an unique ID for the connectionr
	// 2. store its socket object into a global object

	// for each subscribe eent call
	// 1. store the conn ID to a global object keeping track of all channels subscriptions

	// so when publishing data to a channel
	// 1. filter all active connection id subscribe to that channel
	// 2 get all conn object for thse id
	// 3 use the connection object to send the message
}
