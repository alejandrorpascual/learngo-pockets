package cache_test

import (
	"fmt"
	cache "learngo-pockets/genericcache"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	TestCache is an integration test.

- Create a cache and check that it is empty?
- Upsert a new key <5, fünf> in the cache
- Read the value for this entry
- Upsert for the same entry with the new value
- Read the new value
- Upsert another key <3, drei> and check that is doesn't override
- Delete 5 and check that 3 still exists
*/
func TestCache(t *testing.T) {
	c := cache.New[int, string]()

	_ = c.Upsert(5, "fünf")
	v, ok := c.Read(5)

	assert.True(t, ok)
	assert.Equal(t, v, "fünf")

	_ = c.Upsert(5, "pum")
	v, ok = c.Read(5)

	assert.True(t, ok)
	assert.Equal(t, v, "pum")

	_ = c.Upsert(3, "drei")

	v, ok = c.Read(3)

	assert.True(t, ok)
	assert.Equal(t, v, "drei")

	v, ok = c.Read(5)

	assert.True(t, ok)
	assert.Equal(t, v, "pum")

	c.Delete(5)

	v, ok = c.Read(5)

	assert.False(t, ok)
	assert.Equal(t, v, "")

	v, ok = c.Read(3)

	assert.True(t, ok)
	assert.Equal(t, v, "drei")
}

func TestCache_Parallel_goroutines(t *testing.T) {
	c := cache.New[int, string]()
	const parallelTasks = 10
	wg := sync.WaitGroup{}
	wg.Add(parallelTasks)

	for i := range parallelTasks {
		go func(j int) {
			defer wg.Done()
			c.Upsert(4, fmt.Sprint(j))
		}(i)
	}

	wg.Wait()
}

func TestCache_Parallel(t *testing.T) {
	c := cache.New[int, string]()

	t.Run("write six", func(t *testing.T) {
		t.Parallel()
		c.Upsert(6, "six")
	})

	t.Run("write kuus", func(t *testing.T) {
		t.Parallel()
		c.Upsert(6, "kuus")
	})
}
