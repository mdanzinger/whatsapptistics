package chat

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"strings"

	"github.com/mdanzinger/whatsapptistics/src/job"

	"github.com/nu7hatch/gouuid"
)

// ChatService represents a service for managing chat
type ChatService interface {
	New(ctx context.Context, r io.Reader) error
	Retrieve(id string) (*Chat, error)
}

type chatService struct {
	cr        ChatRepository
	jobSource job.Source
	logger    *log.Logger
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

	// Upload to storage
	if err = cs.cr.Upload(ctx, c); err != nil {
		cs.logger.Print(err)
		return err
	}

	// Create Analyze Job
	j := &job.Chat{ChatID: cid}
	if err = cs.jobSource.QueueJob(j); err != nil {
		cs.logger.Println(err)
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
