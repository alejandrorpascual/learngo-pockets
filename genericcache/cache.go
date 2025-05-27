package cache

// Cache is a key-value storage
type Cache[K comparable, V any] struct {
	data map[K]V
}

func New[K comparable, V any]() Cache[K, V] {
	return Cache[K, V]{
		data: make(map[K]V),
	}
}

// Read returns the associated value for a key, and a boolean of false if
// the key is absent.
func (c *Cache[K, V]) Read(key K) (V, bool) {
	v, ok := c.data[key]
	return v, ok
}

// Upsert overrides the value for a given key.
func (c *Cache[K, V]) Upsert(key K, value V) error {
	c.data[key] = value

	// Do not return an error for the moment, but it can happen in the near
	// future.
	return nil
}

// Delete removes the entry for the given key.
func (c *Cache[K, V]) Delete(key K) {
	delete(c.data, key)
}
