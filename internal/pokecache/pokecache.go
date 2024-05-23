package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu  sync.RWMutex
	val map[string]cacheEntry
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) *Cache {

	c := &Cache{
		val: make(map[string]cacheEntry),
		mu:  sync.RWMutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) (val []byte, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.val[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

// reapLoop periodically removes entries older than the interval.
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		for key, entry := range c.val {
			if now.Sub(entry.createdAt) > interval {
				delete(c.val, key)
			}
		}
		c.mu.Unlock()
	}
}
