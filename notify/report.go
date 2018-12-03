package notify

import (
	"bytes"
	"html/template"
)

const (
	TEMPLATE_NAME = "notification.gohtml"
	SERVICE_URL   = "http://whatsapptistics.com"
)

type Notifier interface {
	Notify(message string, email string) error
}

type EmailNotifier struct {
	notifer Notifier
}

func (n *EmailNotifier) Notify(reportID string, email string) error {
	// Get Template
	t, err := n.parseTemplate(reportID)
	if err != nil {
		return err
	}

	if err := n.notifer.Notify(t, email); err != nil {
		return err
	}
	return nil
}

func (n *EmailNotifier) parseTemplate(reportID string) (string, error) {
	t, err := template.ParseFiles("../../web/template/" + TEMPLATE_NAME)
	if err != nil {
		return "", err
	}
	buffer := new(bytes.Buffer)
	url := template.URL(SERVICE_URL + "/report/" + reportID)
	if err = t.Execute(buffer, url); err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func NewNotifier(notifier Notifier) *EmailNotifier {
	return &EmailNotifier{notifer: notifier}
}
