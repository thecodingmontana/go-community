package utils

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/thecodingmontana/go-community/internal/database/models"
	"github.com/thecodingmontana/go-community/pkg/types"
)

const (
	// Time allowed to write a message to the peer
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait
	pingPeriod = (pongWait * 9) / 10
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(req *http.Request) bool {
		origin := req.Header.Get("Origin")

		// List of allowed origins
		allowedOrigins := []string{
			"https://go-community.thecodingmontana.com",
			// Add any other allowed domains
		}

		// Check if the request origin is in the allowed list
		for _, allowed := range allowedOrigins {
			if origin == allowed {
				return true
			}
		}

		// Optional: Allow localhost for testing
		if strings.HasPrefix(origin, "http://localhost:") {
			return true
		}

		// Log unauthorized attempts
		log.Printf("Unauthorized WebSocket connection attempt from origin: %s", origin)
		return false
	},
}

// Client represents a WebSocket connection client

func (client *Client) ReadPump(hub *Hub) {
	defer func() {
		hub.Unregister <- client
		client.Conn.Close()
	}()

	// Set read deadline and pong handler
	client.Conn.SetReadDeadline(time.Now().Add(pongWait))
	client.Conn.SetPongHandler(func(string) error {
		client.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, rawMessage, err := client.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Parse the incoming message
		var msg types.SocketMessage
		if err := json.Unmarshal(rawMessage, &msg); err != nil {
			log.Printf("Error unmarshaling message: %v", err)
			continue
		}

		switch msg.Type {
		case "message":
			var chatMsg types.ChatMessage
			if err := json.Unmarshal(msg.Payload, &chatMsg); err != nil {
				log.Printf("Error unmarshaling chat message: %v", err)
				continue
			}

			// save message to DB
			savedMessage, errSavedMessage := client.Queries.AddNewMessage(ctx, models.AddNewMessageParams{
				ID:       uuid.New().String(),
				UserID:   chatMsg.User.ID,
				Content:  chatMsg.Content,
				ImageUrl: chatMsg.ImageURL,
			})

			if errSavedMessage != nil {
				log.Printf("failed to save message: %v", errSavedMessage)
				continue
			}

			chatResponse := types.ChatResponse{
				ID:      savedMessage.ID,
				Content: savedMessage.Content,
				User: types.User{
					ID:            chatMsg.User.ID,
					Email:         chatMsg.User.Email,
					Username:      chatMsg.User.Avatar,
					EmailVerified: chatMsg.User.EmailVerified,
					Avatar:        chatMsg.User.Avatar,
				},
				ImageURL:  savedMessage.ImageUrl,
				Deleted:   savedMessage.Deleted,
				UpdatedAt: savedMessage.UpdatedAt.Time.String(),
				CreatedAt: savedMessage.CreatedAt.Time.String(),
			}

			response := types.SocketMessage{
				Type: "message",
				Payload: func() json.RawMessage {
					data, _ := json.Marshal(chatResponse)
					return data
				}(),
			}

			data, marshallErr := json.Marshal(response)
			if marshallErr != nil {
				log.Printf("Error marshaling response: %v", err)
				continue
			}
			hub.Broadcast <- data
		}
		// Context is automatically cancelled here due to defer
	}
}

func (client *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		client.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-client.Send:
			client.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := client.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued messages to the current websocket message
			n := len(client.Send)
			for i := 0; i < n; i++ {
				w.Write([]byte{'\n'})
				w.Write(<-client.Send)
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			client.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := client.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
