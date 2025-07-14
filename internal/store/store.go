package store

type GedisStore struct {
	data map[string]string
}

func NewStore() *GedisStore {
	return &GedisStore{
		data: make(map[string]string),
	}
}

func (store *GedisStore) Set(key, value string) {
	store.data[key] = value
}

func (store *GedisStore) Get(key string) (string, bool) {
	value, exists := store.data[key]
	return value, exists
}

func (store *GedisStore) Delete(key string) {
	delete(store.data, key)
}

func (store *GedisStore) Keys() []string {
	keys := make([]string, 0, len(store.data)) // Initialize a slice to hold the keys
	for key := range store.data {
		keys = append(keys, key)
	}
	return keys
}
