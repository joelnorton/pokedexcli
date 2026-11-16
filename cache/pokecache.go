package pokecache

type Cache struct {
	cache map[string][]byte
}

func NewCache() Cache {
	return Cache{cache: make(map[string][]byte)}
}

func (c *Cache) addCache(key string, val []byte) {
	c.cache[key] = val
}

func (c *Cache) getCache(key string) ([]byte, bool) {
	v, ok := c.cache[key]
	return v, ok
}
