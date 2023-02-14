package inMemoryCache

type Cache struct {
	c map[string]interface{}
}

func NewCache() Cache {
	return Cache{
		c: make(map[string]interface{}),
	}
}

func (c Cache) Set(key string, value interface{}) {
	c.c[key] = value
}

func (c Cache) Get(key string) interface{} {
	return c.c[key]
}

func (c Cache) Delete(key string) {
	delete(c.c, key)
}
