package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache    map[string]cacheEntry
	interval time.Duration
	mu       sync.Mutex
	stop     chan struct{}
	stopped  sync.Once
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	if interval <= 0 {
		interval = time.Second
	}
	c := Cache{
		cache:    make(map[string]cacheEntry),
		interval: interval,
		stop:     make(chan struct{}),
	}
	go c.reapCache()
	return &c
}

func (c *Cache) AddCache(key string, val []byte) {
	v := cacheEntry{
		val:       append([]byte(nil), val...),
		createdAt: time.Now(),
	}
	c.mu.Lock()
	c.cache[key] = v
	c.mu.Unlock()
}

func (c *Cache) GetCache(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	ce, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	if !ce.createdAt.Add(c.interval).After(time.Now()) {
		delete(c.cache, key)
		return nil, false
	}
	return append([]byte(nil), ce.val...), true
}

func (c *Cache) Close() {
	c.stopped.Do(func() {
		close(c.stop)
	})
}

func (c *Cache) reapCache() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for {
		select {
		case <-c.stop:
			return
		case t := <-ticker.C:
			c.reap(t)
		}
	}
}

func (c *Cache) reap(t time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, ce := range c.cache {
		if !ce.createdAt.Add(c.interval).After(t) {
			delete(c.cache, k)
		}
	}
}
