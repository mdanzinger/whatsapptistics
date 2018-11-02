package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/mdanzinger/whatsapptistics/report"
	"time"
)

// reportCacheRepo is an implementation of a ReportRepository to be used as a cache layer
type reportCacheRepo struct {
	client *redis.Client
}

const defaultExpirationTime = time.Hour * 24 * 7 // a week

func (c *reportCacheRepo) Get(ctx context.Context, key string) (*report.Report, error) {
	rjson, err := c.client.Get(key).Result()
	if err != nil {
		return nil, err
	}
	// No errors, create report and populate
	report := &report.Report{}

	json.Unmarshal([]byte(rjson), report)

	return report, nil
}

func (c *reportCacheRepo) Store(r *report.Report) error {
	data, err := json.Marshal(&r)
	if err != nil {
		return err
	}
	err = c.client.Set(r.ReportID, data, defaultExpirationTime).Err()
	if err != nil {
		return err
	}
	return nil
}

// NewReportCacheRepo returns a report repository to be used as a cache layer
func NewReportCacheRepo() *reportCacheRepo {
	c := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // TODO: Pull from config / env var
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	_, err := c.Ping().Result()
	if err != nil {
		fmt.Printf("Error creating cache client -> %s", err)
	}

	return &reportCacheRepo{
		client: c,
	}
}
