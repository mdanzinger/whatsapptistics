package cache

type ReportCache struct {
	cache Cache
}

func NewReportCache(c Cacher) *ReportCache {
	return &ReportCache{
		cache: c,
	}
}
