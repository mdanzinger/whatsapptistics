package cache

type Cache struct {
	Key    string
	Value  string
	Cacher Cacher
}

type Cacher interface {
	Get(string) (interface{}, error)
	Set(string, interface{}) error
}

func NewCache(c Cacher) *Cache {
	return &Cache{
		Cacher: c,
	}
}
