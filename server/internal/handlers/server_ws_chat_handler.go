package handlers

import (
	"log"
	"net/http"

	"github.com/thecodingmontana/go-community/internal/database/models"
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
		Conn: conn,
		Send: make(chan []byte, 256),
	}

	hub.Register <- client

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
