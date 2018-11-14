package sqs

import (
	"testing"

	"github.com/mdanzinger/whatsapptistics/src/job"
)

func TestNewAnalyzeJobSource(t *testing.T) {
	b := NewAnalyzeJobSource(nil)
	var _ job.AnalyzeJobSource = (*b)(nil)
}
