package chat

import "context"

type Chat struct {
	ChatID  string
	Content []byte
}

type ChatRepository interface {
	Download(id string) (*Chat, error)
	Upload(context.Context, *Chat) error
}
