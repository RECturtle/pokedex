package pokecache

import "time"

func (c *Cache) Add(key string, val []byte) {
	// Add an item to the cache
	// use a RWMutex
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{time.Now(), val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	// get an item and return it in bytes and a bool if found
	// use a RWMutex
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) clearCache(retentionTime time.Time) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.data {
		if entry.createdAt.Before(retentionTime) {
			delete(c.data, key)
		}
	}
	return nil
}

func (c *Cache) realLoop() {
	// loop through and remove entries older than cache interval
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for tickTime := range ticker.C {
		c.clearCache(tickTime)
	}
}
