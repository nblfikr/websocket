package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":3000", "http service address")

func welcome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not Found Woy", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("websocket service"))
}

func main() {
	flag.Parse()

	http.HandleFunc("/", welcome)
	http.HandleFunc("/ws", wsocket)

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("failed to launch the server")
	}
}
