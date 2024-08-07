package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entry map[string]CacheEntry
	mu    *sync.RWMutex
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entry: map[string]CacheEntry{},
		mu:    &sync.RWMutex{},
	}

	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.entry[key] = CacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	dat, ok := c.entry[key]
	if !ok {
		return nil, false
	}
	return dat.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.mu.Lock()
		for k, v := range c.entry {
			if v.createdAt.Before(time.Now().UTC().Add(-interval)) {
				delete(c.entry, k)
			}
		}

		c.mu.Unlock()
	}
}
