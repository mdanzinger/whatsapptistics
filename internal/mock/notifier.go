package mock

type Notifier struct {
	NotifyFn func(message string, email string) error
}

func (n *Notifier) Notify(message string, email string) error {
	return n.NotifyFn(message, email)
}
