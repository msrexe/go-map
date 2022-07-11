package gomap

import (
	"sync"
)

type Map[K comparable, V any] map[K]V

type GoMap[K comparable, V any] struct {
	mu sync.Mutex

	goMap Map[K, V]
}

// Creates a new GoMap instance
func New[K comparable, V any](size int) *GoMap[K, V] {
	return &GoMap[K, V]{
		mu:    sync.Mutex{},
		goMap: Map[K, V](make(map[K]V, size)),
	}
}

// Get a value from the map
func (gm *GoMap[K, V]) Get(key K) V {
	gm.mu.Lock()
	defer gm.mu.Unlock()

	return gm.goMap[key]
}

// Set a key-value in the map
func (gm *GoMap[K, V]) Set(key K, val V) {
	gm.mu.Lock()
	defer gm.mu.Unlock()

	gm.goMap[key] = val
}

// Delete a key-value from the map
func (gm *GoMap[K, V]) Delete(key K) {
	gm.mu.Lock()
	defer gm.mu.Unlock()

	delete(gm.goMap, key)
}

// Get size of the map
func (gm *GoMap[K, V]) Length() int {
	gm.mu.Lock()
	defer gm.mu.Unlock()

	return len(gm.goMap)
}

// Get all keys from the map
func (gm *GoMap[K, V]) Keys() []K {
	gm.mu.Lock()
	defer gm.mu.Unlock()

	keys := make([]K, len(gm.goMap))
	i := 0
	for k := range gm.goMap {
		keys[i] = k
		i++
	}
	return keys
}

// Get all values from the map
func (gm *GoMap[K, V]) Values() []V {
	gm.mu.Lock()
	defer gm.mu.Unlock()

	values := make([]V, len(gm.goMap))
	i := 0
	for _, v := range gm.goMap {
		values[i] = v
		i++
	}
	return values
}

// Clear the map
func (gm *GoMap[K, V]) Clear() {
	gm.mu.Lock()
	defer gm.mu.Unlock()

	gm.goMap = make(map[K]V)
}

// Check key exists in the map
func (gm *GoMap[K, V]) IsExists(key K) bool {
	gm.mu.Lock()
	defer gm.mu.Unlock()

	_, ok := gm.goMap[key]
	return ok
}
