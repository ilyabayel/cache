package cache

type Cache struct {
	heap map[string]any
}

func New() *Cache {
	return &Cache{
		heap: make(map[string]any),
	}
}

func (c *Cache) Set(key string, value any) {
	c.heap[key] = value
}

func (c *Cache) Get(key string) any {
	return c.heap[key]
}

func (c *Cache) Delete(key string) {
	delete(c.heap, key)
}
