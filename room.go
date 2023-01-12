package main

var Rooms []Room

type Room struct {
	ID         string
	Members    map[*Client]bool
	Broadcast  chan *Message
	Register   chan *Client
	Unregister chan *Client
}

func prepare(id string) *Room {

	// is room exist
	for _, room := range Rooms {
		if room.ID == id {
			return &room
		}
	}

	room := &Room{
		ID:         id,
		Members:    make(map[*Client]bool),
		Broadcast:  make(chan *Message),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}

	Rooms = append(Rooms, *room)

	go room.run()
	return room
}

func (r *Room) run() {
	for {
		select {
		case client := <-r.Register:
			r.Members[client] = true
		case client := <-r.Unregister:
			if _, ok := r.Members[client]; ok {
				client.ws.Close()
				delete(r.Members, client)
			}
		case msg := <-r.Broadcast:
			for m := range r.Members {
				m.ws.WriteJSON(msg)
			}
		}
	}
}
