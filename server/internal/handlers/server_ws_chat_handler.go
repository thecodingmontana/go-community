package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/thecodingmontana/go-community/internal/database/models"
	"github.com/thecodingmontana/go-community/pkg/types"
	"github.com/thecodingmontana/go-community/pkg/utils"
)

func ServeWSChatHandler(res http.ResponseWriter, req *http.Request, hub *utils.Hub, queries *models.Queries) {
	conn, err := utils.Upgrader.Upgrade(res, req, nil)

	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		utils.RespondWithError(res, http.StatusInternalServerError, "Could not upgrade connection")
		return
	}

	client := &utils.Client{
		Conn:    conn,
		Send:    make(chan []byte, 256),
		Queries: queries,
	}

	hub.Register <- client

	// Send message history to new client
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		savedMesages, err := queries.GetAllMessages(ctx)
		if err != nil {
			log.Printf("Error fetching message history: %v", err)
			return
		}

		var messages []types.ChatResponse

		for _, message := range savedMesages {
			messages = append(messages, types.ChatResponse{
				ID:        message.ID,
				Content:   message.Content,
				ImageURL:  message.ImageUrl,
				FileURL:   message.FileUrl,
				Deleted:   message.Deleted,
				CreatedAt: message.CreatedAt.Time.String(),
				UpdatedAt: message.UpdatedAt.Time.String(),
				User: types.User{
					ID:            message.ByID,
					Email:         message.ByEmail,
					Username:      message.ByUsername,
					EmailVerified: message.ByEmailVerified,
					Avatar:        message.ByAvatar.String,
				},
			})
		}

		response := types.SocketMessage{
			Type: "history",
			Payload: func() json.RawMessage {
				data, _ := json.Marshal(messages)
				return data
			}(),
		}

		data, marshalErr := json.Marshal(response)
		if marshalErr != nil {
			log.Printf("Error marshaling history: %v", err)
			return
		}

		client.Send <- data
	}()

	// Start client pumps
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered from panic in client pumps: %v", r)
			}
		}()
		go client.ReadPump(hub)
		go client.WritePump()
	}()
}
