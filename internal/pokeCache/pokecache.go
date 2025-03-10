package pokeCache

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

func NewPokeCache(interval time.Duration) *PokeCache {
	cache := &PokeCache{
		Data: make(map[string]CacheEntry),
		Mu:   &sync.RWMutex{},
	}
	go func(duration time.Duration) {
		ticker := time.NewTicker(duration)
		for {
			select {
			case <-ticker.C:
				cache.readLoop(duration)
			}
		}
	}(interval)
	return cache
}

func (pC *PokeCache) Add(key string, val []byte) {
	pC.Mu.Lock()
	defer pC.Mu.Unlock()
	pC.Data[key] = CacheEntry{
		CreatedAt: time.Now(),
		Val:       val,
	}
}

func (pC *PokeCache) Get(key string) ([]byte, bool) {
	pC.Mu.RLock()
	defer pC.Mu.RUnlock()
	cache, exists := pC.Data[key]
	if !exists {
		return nil, false
	}
	return cache.Val, true
}

func (pC *PokeCache) readLoop(interval time.Duration) {
	for key, val := range pC.Data {
		if time.Since(val.CreatedAt) >= interval {
			pC.Mu.Lock()
			delete(pC.Data, key)
			pC.Mu.Unlock()
		}
	}

}
