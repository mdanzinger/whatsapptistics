package ses_test

import (
	"testing"

	"github.com/mdanzinger/whatsapptistics/internal/notify/ses"
)

func TestSesNotifier_Notify(t *testing.T) {
	n := ses.NewSESNotifier()

	err := n.Notify("<h1> hi! </h1>", "mendydanzinger@gmail.com")
	if err != nil {
		t.Errorf("got error sending test email: %v", err)
	}
}
