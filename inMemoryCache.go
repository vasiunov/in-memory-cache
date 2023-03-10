package cache

import (
	"sync"
	"time"
)

// Cache type contains map[string]interface{}, time-to-live and mutex fields.
type Cache struct {
	c   map[string]interface{}
	ttl time.Duration
	mu  *sync.Mutex
}

// Function NewCache() initialise Cache struct. Each key-value pair expires in 5 sec.
func NewCache() Cache {
	return Cache{
		c:   make(map[string]interface{}),
		ttl: time.Second * 5,
		mu:  new(sync.Mutex),
	}
}

// Method Set() sets the interface{} value for a string key. Method contains goroutine deleting data according to time-to-live.
func (c Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	c.c[key] = value
	c.mu.Unlock()

	go func() {
		time.Sleep(c.ttl)
		c.Delete(key)
	}()
}

// Method Get() returns the interface{} value corresponding to a string key.
func (c Cache) Get(key string) interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.c[key]
}

// Method Delete() deletes the key-interface{} pair. Use Delete() if you don't need to wait for time-to-live to expire.
func (c Cache) Delete(key string) {
	c.mu.Lock()
	delete(c.c, key)
	c.mu.Unlock()
}
