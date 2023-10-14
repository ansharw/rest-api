package api

// API for session
type Session struct {
	UserId         string `json:"user_id"`
	UserAgent      string `json:"user_agent"`
	BrowserSession string `json:"browser_session"`
	Type           string `json:"type"`
}
