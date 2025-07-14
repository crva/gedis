package store

import "sync"

type GedisStore struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewStore() *GedisStore {
	return &GedisStore{
		data: make(map[string]string),
	}
}

func (store *GedisStore) Set(key, value string) {
	store.mu.Lock()
	defer store.mu.Unlock()

	store.data[key] = value
}

func (store *GedisStore) Get(key string) (string, bool) {
	store.mu.RLock()
	defer store.mu.RUnlock()

	value, exists := store.data[key]
	return value, exists
}

func (store *GedisStore) Delete(key string) {
	store.mu.Lock()
	defer store.mu.Unlock()

	delete(store.data, key)
}

func (store *GedisStore) Keys() []string {
	store.mu.RLock()
	defer store.mu.RUnlock()

	keys := make([]string, 0, len(store.data)) // Initialize a slice to hold the keys
	for key := range store.data {
		keys = append(keys, key)
	}
	return keys
}
