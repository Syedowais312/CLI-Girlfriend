package model

type ChatMessage struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

type ChatHistory struct {
    Messages []ChatMessage `json:"messages"`
}
