package cache

import (
	"slices"
	"sync"
	"time"
)

// Cache is a key-value storage
type Cache[K comparable, V any] struct {
	ttl time.Duration

	mu   sync.RWMutex
	data map[K]entryWithTimeout[V]

	maxSize           int
	chronologicalKeys []K
}

type entryWithTimeout[V any] struct {
	value   V
	expires time.Time
}

func New[K comparable, V any](maxSize int, ttl time.Duration) Cache[K, V] {
	return Cache[K, V]{
		ttl:               ttl,
		data:              make(map[K]entryWithTimeout[V]),
		maxSize:           maxSize,
		chronologicalKeys: make([]K, 0, maxSize),
	}
}

// Read returns the associated value for a key, and a boolean of false if
// the key is absent.
func (c *Cache[K, V]) Read(key K) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	v, ok := c.data[key]

	var zeroV V

	switch {
	case !ok:
		return zeroV, false
	case v.expires.Before(time.Now()):
		// The value has expired.
		c.deleteKeyValue(key)
		return zeroV, false
	default:
		return v.value, ok
	}

}

// Upsert overrides the value for a given key.
func (c *Cache[K, V]) Upsert(key K, value V) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, alreadyPresent := c.data[key]
	switch {
	case alreadyPresent:
		c.deleteKeyValue(key)
	case len(c.data) == c.maxSize:
		c.deleteKeyValue(c.chronologicalKeys[0])
	}

	c.addKeyValue(key, value)
	return nil
}

// Delete removes the entry for the given key.
func (c *Cache[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.deleteKeyValue(key)
}

func (c *Cache[K, V]) addKeyValue(key K, value V) {
	c.data[key] = entryWithTimeout[V]{
		value:   value,
		expires: time.Now().Add(c.ttl),
	}
	c.chronologicalKeys = append(c.chronologicalKeys, key)
}

func (c *Cache[K, V]) deleteKeyValue(key K) {
	c.chronologicalKeys = slices.DeleteFunc(
		c.chronologicalKeys,
		func(k K) bool { return k == key })
	delete(c.data, key)
}
