package utils

import (
	"github.com/gorilla/websocket"
	"github.com/thecodingmontana/go-community/internal/database/models"
)

type Client struct {
	Conn    *websocket.Conn // websocket connection
	Send    chan []byte     // channel for sending messages
	Queries *models.Queries // handle database operations
}
