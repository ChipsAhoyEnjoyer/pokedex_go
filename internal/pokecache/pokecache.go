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
	Data     map[string]CacheEntry
	Mu       *sync.RWMutex
	interval time.Time
}

func NewPokeCache(i time.Time) *PokeCache {
	return &PokeCache{
		Data:     make(map[string]CacheEntry),
		Mu:       &sync.RWMutex{},
		interval: i,
	}
}

func (pC *PokeCache) Add(key string, val []byte) {
	pC.Data[key] = CacheEntry{
		CreatedAt: time.Now(),
		Val:       val,
	}
}
