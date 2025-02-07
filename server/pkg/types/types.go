package types

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

type SocketMessage struct {
	ID       int    `json:"id"`
	Content  string `json:"content"`
	UserID   string `json:"user_id"`
	ImageURL string `json:"image_url"`
}
