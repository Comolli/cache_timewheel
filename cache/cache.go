package cache

import (
	"sync"
	"time"
)

type Cacheoption func(cache *Cache)
type Cache struct {
	name   string
	data   map[string]interface{}
	expire time.Duration
	lock   sync.Mutex
}

func NewCache(expire time.Duration, opts ...Cacheoption) *Cache {
	cache := &Cache{
		data:   make(map[string]interface{}),
		expire: expire,
	}
	for _, opt := range opts {
		opt(cache)
	}
	if cache.name == "" {
		//todo
	}
	return cache
}

func (c *Cache) Del(key string) {
	c.lock.Lock()
	delete(c.data, key)
	c.lock.Unlock()
}

func (c *Cache) Update(key string, value interface{}) {
	c.lock.Lock()
	c.data[key] = value
	c.lock.Unlock()
}

func (c *Cache) Get(key string) (interface{}, bool) {
	//todo
	//add lru
	value, ok := c.data[key]
	if !ok {
		return "", false
	}
	return value, true

}
