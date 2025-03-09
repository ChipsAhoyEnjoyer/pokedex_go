package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type PokeCache struct {
	Data map[string]CacheEntry
	Mu   *sync.RWMutex
}
