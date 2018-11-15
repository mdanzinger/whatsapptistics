package sqs

// Message is a message received off of the queue. We use this to unmarshal the data.
type message struct {
	Message string
}
