package notify

import (
	"fmt"
	"testing"

	"github.com/mdanzinger/whatsapptistics/internal/mock"
)

func TestEmailNotifier_Notify(t *testing.T) {
	mockNotifier := &mock.Notifier{}
	mockNotifier.NotifyFn = func(message string, email string) error {
		t.Run("Sending Mock Email", func(t *testing.T) {
			if message == "" {
				t.Errorf("message is empty")
			}
		})
		return nil
	}
	n := NewNotifier(mockNotifier)

	err := n.Notify("55cc55", "mendyDanzinger@gmail.com")
	if err != nil {
		fmt.Println(err)
		t.Fatalf("got error sending notification: %v", err)
	}
}
