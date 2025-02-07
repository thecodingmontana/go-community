package utils

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/thecodingmontana/go-community/pkg/types"
)

type Hub struct {
	Clients    map[*Client]bool // maps all connected clients
	Broadcast  chan []byte      // channel for broadcasting messages
	Register   chan *Client     // channel for registering clients
	Unregister chan *Client     // channel for unregistering clients
	Mutex      sync.Mutex       // thread-safe operations
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.Register:
			// Register new client
			hub.Mutex.Lock()
			hub.Clients[client] = true
			hub.Mutex.Unlock()
		case client := <-hub.Unregister:
			// Unregister client
			hub.Mutex.Lock()
			if _, ok := hub.Clients[client]; ok {
				delete(hub.Clients, client)
				close(client.Send)
			}
			hub.Mutex.Unlock()
		case message := <-hub.Broadcast:
			// Send message to all clients
			hub.Mutex.Lock()
			var msg types.SocketMessage
			json.Unmarshal(message, &msg)
			log.Print(msg)
			log.Printf("Broadcasting message to %d clients", len(hub.Clients))
			for client := range hub.Clients {
				select {
				case client.Send <- message:
					// Message sent to client channel
					log.Printf("Message sent to client channel")
				default:
					// Client buffer is full, clean up
					log.Printf("Failed to send to client, removing")
					close(client.Send)
					delete(hub.Clients, client)
				}
			}
			hub.Mutex.Unlock()
		}
	}
}
