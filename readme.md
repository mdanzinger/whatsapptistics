# whatsapptistics [![Go Report Card](https://goreportcard.com/badge/github.com/mdanzinger/whatsapptistics?cache=clear)](https://goreportcard.com/report/github.com/mdanzinger/whatsapptistics?cache=clear) [![Project Status](https://img.shields.io/badge/%20status:-launched-green.svg)](https://img.shields.io/badge/%20status:-launched-green.svg)
[Whatsapptistics.com](https://whatsapptistics.com)

whatsapptistics is hobby project created to provide interesting and detailed analysis on uploaded WhatsApp chats. This is a monorepo for the service, so all code related to whatsapptistics can be found here. The core technologies utilized for this app are [Go](https://golang.org/), [SQS / SNS](https://aws.amazon.com/sqs/), [DynamoDB](https://aws.amazon.com/dynamodb/), [Redis](https://redis.io/), [Docker / Docker Swarm](https://www.docker.com/) and an array of frontend tools and libraries (**scss**,**gulp**,**es6**,**chart.js**)

## Getting Started

To get this project set up on your system locally, follow the instructions below:

### Prerequisites

TODO: provide list of prerequisites

```
Give examples
```

### Installing

A step by step series of examples that tell you how to get a development env running

TODO: Complete...

```
git clone https://github.com/mdanzinger/whatsapptistics
```

etc

```
some more steps..
```



## Running the tests

TODO: Implement tests first... 


### Compiling front-end assets

TODO: Add steps to install node modules, describe coding standards, add a few gulp commands

```
gulp
```

## Deployment

TODO: Implement deployment..

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/mdanzinger/whatsapptistics/tags).

## Authors

* **Mendy Danzinger** - [MDanzinger](https://github.com/mdanzinger)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* WILL ADD! *



### Some internal project notes:
---
Android Regex for breaking apart chat:

```(?P<datetime>\d{1,2}\/\d{1,2}\/\d{2}, \d{1,2}:\d{1,2} ((?i)[ap]m))(?: - )(?P<name>.*?)(?::) (?P<message>.+)```


IOs: 
```(?P<datetime>\d{4}\-\d{2}\-\d{2}, \d{1,2}:\d{1,2}:\d{2} ((?i)[ap]m))(?:\]) (?P<name>.*?)(?::) (?P<message>.+)```



