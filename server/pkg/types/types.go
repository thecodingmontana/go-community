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

type MailBody struct {
	Message []byte
	From    string
}

type SMTPAuth struct {
	Username string
	Password string
	Host     string
	Port     string
}

type User struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	Username      string `json:"username"`
	EmailVerified bool   `json:"emailVerified"`
	Avatar        string `json:"avatar"`
}
