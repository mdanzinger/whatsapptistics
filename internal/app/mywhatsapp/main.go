package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-redis/redis"
	"github.com/mdanzinger/whatsapp/internal/pkg/store/dynamodb"
	redisStore "github.com/mdanzinger/whatsapp/internal/pkg/store/redis"
)

func main() {
	s, err := session.NewSession()
	if err != nil {
		fmt.Errorf("Error creating session -> %s", err)
	}
	c := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = c.Ping().Result()
	if err != nil {
		fmt.Errorf("Error creating cache client -> %s", err)
	}

	r := redisStore.NewReportCache(c)
	d := dynamodb.NewReportStore(s, r)

}
