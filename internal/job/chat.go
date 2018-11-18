package job

// Chat represents a chat job that needs to be analyzed and created into a report
type Chat struct {
	ChatID        string
	UploaderEmail string
}

// Source represents a queue of chats that need to be analyzed
type Source interface {
	QueueJob(*Chat) error
	NextJob() (*Chat, error)
}
