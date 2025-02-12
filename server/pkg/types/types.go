package types

import "encoding/json"

type APIResponse struct {
	StatusCode    int         `json:"statusCode"`
	StatusMessage string      `json:"statusMessage"`
	Data          interface{} `json:"data,omitempty"`
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

type ChatMessage struct {
	ID       int    `json:"id"`
	Content  string `json:"content"`
	ImageURL string `json:"image_url"`
	User     User   `json:"user"`
}

type SocketMessage struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type ChatResponse struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	User      User   `json:"user"`
	ImageURL  string `json:"image_url"`
	FileURL   string `json:"file_url"`
	Deleted   bool   `json:"deleted"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}
