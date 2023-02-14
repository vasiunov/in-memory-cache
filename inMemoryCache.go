// Package inMemoryCache is used for caching data.
package inMemoryCache

// Cache type contains the field map[string]interface{}
type Cache struct {
	c map[string]interface{}
}

// Function NewCache() initialise Cache struct
func NewCache() Cache {
	return Cache{
		c: make(map[string]interface{}),
	}
}

// Method Set() sets the interface{} value for a string key
func (c Cache) Set(key string, value interface{}) {
	c.c[key] = value
}

// Method Get() returns the interface{} value corresponding to a string key
func (c Cache) Get(key string) interface{} {
	return c.c[key]
}

// Method Delete() deletes the key-interface{} pair
func (c Cache) Delete(key string) {
	delete(c.c, key)
}
