package redis

import (
	"github.com/go-redis/redis"
	"github.com/mdanzinger/whatsapp/internal/pkg/report"
)

type cache struct {
	cl *redis.Client
}

func (c *cache) Get(k string) (interface{}, error) {
	//o := c.cl.Get(k)
	return &report.Report{}, nil

}

func (c *cache) Set(k string, v interface{}) error {

}

func NewReportCache(client *redis.Client) *cache {
	return &cache{
		cl: client,
	}
}
