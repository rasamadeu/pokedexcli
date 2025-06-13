package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	dataEntry []byte
	createdAt time.Time
}

type Cache struct {
	data     map[string]cacheEntry
	lifetime time.Duration          // Data lifetime duration in cache
	mut      *sync.Mutex            // Mutex to control I/O operations
	ticker   *time.Ticker           // Ticker used to call reapLoop
	kill     chan bool              // Channel to signal closing cache
}

func (cache *Cache) Add(key string, val []byte) {

	cache.mut.Lock()
	defer cache.mut.Unlock()
	newCacheEntry := cacheEntry{
		dataEntry: val,
		createdAt: time.Now(),
	}

	cache.data[key] = newCacheEntry
}

func (cache *Cache) Get(key string) ([]byte, bool) {

	cache.mut.Lock()
	defer cache.mut.Unlock()
	cacheEntry, ok := cache.data[key]
	if !ok {
		return []byte{}, false
	}
	return cacheEntry.dataEntry, true
}

// Function that kills the cache
func (cache *Cache) Kill() {
	// This line must be before signaling to kill cache, to avoid
	// the ticker being called accidentaly while we are killing the cache
	cache.ticker.Stop()
	close(cache.kill)
}

// Function that deletes data that outlived cache.lifetime
func (cache *Cache) reapLoop() {

	cache.mut.Lock()
	defer cache.mut.Unlock()

	keysOutdated := []string{}
	for key, val := range cache.data {
		// This condition checks if currentTime > createdAt + cache.lifetime
		if time.Now().After(val.createdAt.Add(cache.lifetime)) {
			keysOutdated = append(keysOutdated, key)
		}
	}

	// Remove outdated data from cache
	for _, key := range keysOutdated {
		delete(cache.data, key)
	}
}

func NewCache(lifetime time.Duration) *Cache {
	cache := Cache{
		data:     make(map[string]cacheEntry),
		lifetime: lifetime,
		mut:      &sync.Mutex{},
		ticker:   time.NewTicker(lifetime),
		kill:     make(chan bool),
	}
	
	go func(){
		// We must not defer the ticker outside this go subroutine.
		// Otherwise, we stop the ticker at the end of the CreateCache
		// function and we never run cache.reapLoop
		defer cache.ticker.Stop()
		for  {
			select{
			case <- cache.ticker.C:
				cache.reapLoop()
			case <- cache.kill:
				return
			}
		}
	}()
	
	return &cache
}
