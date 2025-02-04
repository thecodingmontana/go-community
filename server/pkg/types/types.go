package types

import (
	"github.com/gorilla/websocket"
)

type APIResponse struct {
	StatusCode    int         `json:"statusCode"`
	StatusMessage string      `json:"statusMessage"`
	Data          interface{} `json:"data,omitempty"`
}

type Client struct {
	Conn *websocket.Conn // websocket connection
	Send chan []byte     // channel for sending messages
}
