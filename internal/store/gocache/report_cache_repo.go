package gocache

import (
	"context"
	"fmt"
	"time"

	gocache "github.com/patrickmn/go-cache"

	"github.com/mdanzinger/whatsapptistics/internal/report"
)

type reportCacheRepo struct {
	cl *gocache.Cache
}

func (c reportCacheRepo) Get(ctx context.Context, id string) (*report.Report, error) {
	r, found := c.cl.Get(id)
	if found {
		cachereport := r.(*report.Report)
		return cachereport, nil
	}
	return nil, fmt.Errorf("gocache miss")
}

func (c reportCacheRepo) Store(r *report.Report) error {
	cp := r
	c.cl.Set(r.ReportID, cp, gocache.DefaultExpiration)
	return nil
}

func NewReportCache() *reportCacheRepo {
	c := gocache.New(1*time.Hour, 2*time.Hour)

	return &reportCacheRepo{
		cl: c,
	}
}
