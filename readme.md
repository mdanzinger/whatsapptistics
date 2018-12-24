# whatsapptistics [![Go Report Card](https://goreportcard.com/badge/github.com/mdanzinger/whatsapptistics?cache=clear)](https://goreportcard.com/report/github.com/mdanzinger/whatsapptistics?cache=clear) [![Project Status](https://img.shields.io/badge/%20status:-launched-green.svg)](https://img.shields.io/badge/%20status:-launched-green.svg)
[Whatsapptistics.com](https://whatsapptistics.com)

whatsapptistics is hobby project created to provide a breakdown and detailed analysis of uploaded WhatsApp chats. This is a monorepo for the service, so all code related to whatsapptistics can be found here. The core technologies utilized for this app are [Go](https://golang.org/) (for the backend), [SQS / SNS](https://aws.amazon.com/sqs/) (for the job queue system), [DynamoDB](https://aws.amazon.com/dynamodb/) (to store reports), [Redis](https://redis.io/)(to cache reports), [Docker](https://www.docker.com/)(primarily for easy deployment) and an array of frontend technologies (**scss**,**gulp**,**es6**,**highchart.js**)

[![sample-report](https://i.imgur.com/sBUisf9.jpg)](http://whatsapptistics.com)

## Getting Started

To get whatsapptistics set up on your system locally, follow the instructions below:

### Prerequisites

1. A working `Go` environment (or docker)
2. An AWS account to create the necessary services  


### AWS

While the "business logic" of whatsapptistics is abstracted away using interfaces, AWS services are the only implementation of said interfaces. At some later time I may add some implementations that don't require any third party services, but for now, the following AWS services are required for whatsapptistics to function. 

1. S3 Bucket
2. SQS Queue
3. SNS Topic  (Subscribe the SQS queue created above to the SNS topic)
4. Dynamodb table
5. SES Email account


When the following services are created, rename the `.env.sample` file to `.env` and replace the variables with your own values
```$xslt
mv .env.sample .env
``` 


### Building

If Docker is not being used, binaries must be built into the project root:


```$xslt
cd cmd/whatsapptistics && go build -o ../../server 
cd ../whatsapptistics-consumer && go build -o ../../analyzer
```

Docker instructions coming soon...



## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details



### Some internal project notes:
---
Android Regex for breaking apart chat:

```(?P<datetime>\d{1,2}\/\d{1,2}\/\d{2}, \d{1,2}:\d{1,2} ((?i)[ap]m))(?: - )(?P<name>.*?)(?::) (?P<message>.+)```


IOs: 
```(?P<datetime>\d{4}\-\d{2}\-\d{2}, \d{1,2}:\d{1,2}:\d{2} ((?i)[ap]m))(?:\]) (?P<name>.*?)(?::) (?P<message>.+)```



