package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/mdanzinger/mywhatsapp/internal/pkg/report"
	"time"
)

type cache struct {
	cl *redis.Client
}

var (
	defaultExpirationTime = time.Hour * 24 * 7
)

func (c *cache) Get(ctx context.Context, k string) (*report.Report, error) {
	//o := c.cl.Get(k)
	rjson, err := c.cl.Get(k).Result()
	if err != nil {
		return nil, err
	}
	r := &report.Report{}

	json.Unmarshal([]byte(rjson), r)
	return r, nil

}

func (c *cache) Store(ctx context.Context, r *report.Report) error {
	err := c.cl.Set(r.ReportID, r, defaultExpirationTime).Err()
	if err != nil {
		return err
	}
	return nil
}

func NewReportCache() *cache {
	c := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := c.Ping().Result()
	if err != nil {
		fmt.Errorf("Error creating cache client -> %s", err)
	}

	return &cache{
		cl: c,
	}
}