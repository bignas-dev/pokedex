// Package pokecache contains the caching logic needed
package pokecache

import (
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	data map[string]cacheEntry
	mu sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	data := make(map[string]cacheEntry)

	c := &Cache{
		data: data,
		mu: sync.Mutex{},
	}
	go func() {
		ticker := time.NewTicker(interval)
		for range ticker.C {
			c.mu.Lock()
			for key, val := range c.data {
				if time.Since(val.createdAt) > interval {
					delete(c.data, key)
				}
			}
			c.mu.Unlock()
		}
	}()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if val, ok := c.data[key]; ok {
		return val.val, true
	}
	return nil, false
}

