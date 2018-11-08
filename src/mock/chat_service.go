package mock

import (
	"context"
	"io"

	"github.com/mdanzinger/whatsapptistics/src/chat"
)

//ChatService represents a mock of chat.ChatService
type ChatService struct {
	NewFn      func(ctx context.Context, r io.Reader) error
	NewInvoked bool

	RetrieveFn      func(id string) (*chat.Chat, error)
	RetrieveInvoked bool
}

// New invokes the mock implementation and marks the function as invoked
func (cs *ChatService) New(ctx context.Context, r io.Reader) error {
	cs.NewInvoked = true
	return cs.NewFn(ctx, r)
}

// Retrieve invokes the mock implementation and marks the function as invoked
func (cs *ChatService) Retrieve(id string) (*chat.Chat, error) {
	cs.RetrieveInvoked = true
	return cs.RetrieveFn(id)
}
