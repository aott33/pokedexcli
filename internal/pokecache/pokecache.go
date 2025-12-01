package pokecache

import (
	"time"
)

func NewCache(interval time.Duration) *Cache{
	cache := Cache{
		cacheEntries: map[string]cacheEntry{},
		interval: interval,
	}
	
	go cache.reapLoop()

	return &cache
}

func (cache *Cache) Add(key string, entryByte []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val: entryByte,
	}

	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.cacheEntries[key] = entry
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	entry, ok := cache.cacheEntries[key]
	
	if !ok {
		return []byte{}, false
	}

	return entry.val, ok
}

func (cache *Cache) reapLoop() {
	ticker := time.NewTicker(cache.interval)

	defer ticker.Stop()

	for range ticker.C {
		cache.mu.Lock()

		for key, entry := range cache.cacheEntries {
			if time.Since(entry.createdAt) > cache.interval {
				delete(cache.cacheEntries, key)
			}
		}

		cache.mu.Unlock()
	}
}
