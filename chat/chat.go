package chat

import "context"

// Chat represents an exported chat
type Chat struct {
	ChatID  string
	Content []byte
	Email   string
}

// ChatRepository represents a repository of chats
type ChatRepository interface {
	Download(id string) (*Chat, error)
	Upload(context.Context, *Chat) error
}
