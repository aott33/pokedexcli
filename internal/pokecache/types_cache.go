package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries	map[string]cacheEntry
	interval		time.Duration
	mu				sync.Mutex
}

type cacheEntry struct {
	createdAt		time.Time
	val				[]byte
}
