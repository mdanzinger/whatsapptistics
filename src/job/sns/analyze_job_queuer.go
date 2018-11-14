package sns

import (
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/mdanzinger/whatsapptistics/src/job"
)

type SnsQueuer struct {
	cl *sns.SNS
}

func (s *SnsQueuer) QueueJob(job *job.AnalyzeJob) {

}
