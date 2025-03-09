package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type Cache struct {
	Data map[string]CacheEntry
	Mu   *sync.RWMutex
}
