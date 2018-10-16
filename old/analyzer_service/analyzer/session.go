package analyzer

import (
	"github.com/aws/aws-sdk-go/aws/session"
)

var Sess *session.Session

func Init() error {
	s, err := session.NewSession()
	if err != nil {
		return err
	}
	Sess = s
	return nil
}
