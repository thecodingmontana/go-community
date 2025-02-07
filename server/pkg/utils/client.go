package utils

import "github.com/gorilla/websocket"

type Client struct {
	Conn *websocket.Conn // websocket connection
	Send chan []byte     // channel for sending messages
}
