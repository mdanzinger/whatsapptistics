package notify

// JobNotifier notifies a notification service that a chat was added to storage.
type JobNotifier interface {
	JobNotify(chatid string) error
}
