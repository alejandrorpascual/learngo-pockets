package cache_test

import (
	"fmt"
	cache "learngo-pockets/genericcache"
	"sync"
	"testing"
	"time"

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
	c := cache.New[int, string](5, time.Millisecond*500)

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
	c := cache.New[int, string](5, time.Millisecond*500)
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
	c := cache.New[int, string](5, time.Millisecond*500)

	t.Run("write six", func(t *testing.T) {
		t.Parallel()
		c.Upsert(6, "six")
	})

	t.Run("write kuus", func(t *testing.T) {
		t.Parallel()
		c.Upsert(6, "kuus")
	})
}

func TestCache_TTL(t *testing.T) {
	t.Parallel()

	c := cache.New[string, string](5, time.Millisecond*100)
	c.Upsert("Norwegian", "Blue")

	got, ok := c.Read("Norwegian")
	assert.True(t, ok)
	assert.Equal(t, got, "Blue")

	time.Sleep(time.Millisecond * 200)

	got, ok = c.Read("Norwegian")
	assert.False(t, ok)
	assert.Equal(t, got, "")
}

func TestCache_MaxSize(t *testing.T) {
	t.Parallel()

	c := cache.New[int, int](3, time.Minute)

	c.Upsert(1, 1)
	c.Upsert(2, 2)
	c.Upsert(3, 3)

	got, ok := c.Read(1)
	assert.True(t, ok)
	assert.Equal(t, got, 1)

	// Update 1, which will no longer make it the oldest
	c.Upsert(1, 10)

	// Adding a fourth element will discard the oldest - 2 in this case
	c.Upsert(4, 4)

	// Trying to retrieve an element that should've been discarded by now
	got, ok = c.Read(2)
	assert.False(t, ok)
	assert.Equal(t, got, 0)
}
