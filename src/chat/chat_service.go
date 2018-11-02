package chat

import (
	"context"
	"github.com/nu7hatch/gouuid"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

// ChatService represents a service for managing chat
type ChatService interface {
	New(ctx context.Context, r io.Reader) error
	Retrieve(id string) (*Chat, error)
}

type chatService struct {
	cr     ChatRepository
	logger *log.Logger
}

// New creates a new Chat entity and uploads it to an injected storage repo
func (cs *chatService) New(ctx context.Context, r io.Reader) error {
	// Generate ID of chat entity
	id, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}
	// Clean ID -- remove dashes
	cid := strings.Replace(id.String(), "-", "", -1)

	// Read chat content
	cc, err := ioutil.ReadAll(r)
	if err != nil {
		cs.logger.Fatal(err)
	}

	// New chat
	c := &Chat{
		ChatID:  cid,
		Content: cc,
	}

	if err = cs.cr.Upload(ctx, c); err != nil {
		cs.logger.Print(err)
		return err
	}
	return nil
}

// Chat Content retrieves a chat from an injected storage repo
func (cs *chatService) Retrieve(id string) (*Chat, error) {
	c, err := cs.cr.Download(id)
	if err != nil {
		cs.logger.Print(err)
		return nil, err
	}
	return c, nil
}

func NewChatService(cr ChatRepository, l *log.Logger) *chatService {
	return &chatService{
		cr:     cr,
		logger: l,
	}
}
