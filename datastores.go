package datastores

import "sync"

//###########################
// Memory storage structures
//###########################
// Generic data store
type DataStore[K comparable, V any] struct {
	sync.RWMutex
	cache map[K]V
}

func NewDataStore[K comparable, V any]() *DataStore[K, V] {
	return &DataStore[K, V]{
		cache: make(map[K]V),
	}
}

func (ds *DataStore[K, V]) set(key K, value V) {
	ds.Lock()
	defer ds.Unlock()

	ds.cache[key] = value
}

func (ds *DataStore[K, V]) get(key K) V {
	ds.RLock()
	defer ds.RUnlock()

	return ds.cache[key]
}

func (ds *DataStore[K, V]) setAll(data map[K]V) {
	ds.Lock()
	defer ds.Unlock()

	ds.cache = data
}

func (ds *DataStore[K, V]) unset(key K) bool {
	ds.Lock()
	defer ds.Unlock()

	_, ok := ds.cache[key]
	if ok {
		delete(ds.cache, key)
		return true
	}
	return false
}
