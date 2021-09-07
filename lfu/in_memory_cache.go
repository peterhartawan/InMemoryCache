package lfu

import (
	"fmt"
)

type inMemoryCache struct {
	Limit           int
	EvictionManager EvictionManager
	CacheValues     []string
	CacheKeys       []string
}

func NewInMemoryCache(limit int, em EvictionManager) *inMemoryCache {
	imc := inMemoryCache{}

	imc.Limit = limit
	imc.EvictionManager = em

	return &imc
}

func (imc *inMemoryCache)searchIndexByKey(key string) int {
	for i:=0; i< len(imc.CacheKeys); i++ {
		// Remove cache data
		if imc.CacheKeys[i] == key {
			return i
		}
	}

	return IN_MEMORY_CACHE_KEY_INDEX_NOT_FOUND
}

func (imc *inMemoryCache) Add(key, value string) int {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("Recovering from panic: %v \n", r)
		}
	}()

	// Check key already exists
	isKeyInManagerExists := imc.EvictionManager.Push(key)
	//isKeyInCacheKeysExists := imc.searchIndexByKey(key)

	// Update data if exists
	if isKeyInManagerExists == EVICTION_MANAGER_KEY_EXISTS {
		// Search cache data
		index := imc.searchIndexByKey(key)

		// Update cache data
		imc.CacheValues[index] = value

		return IN_MEMORY_CACHE_REPLACE_EXISTING_KEY
	}

	// Remove least recently used data
	if len(imc.CacheKeys) == imc.Limit {
		// Remove recently added key
		imc.EvictionManager.Pop()

		// Remove latest key
		deletedKey := imc.EvictionManager.Pop()

		// Re add new key
		imc.EvictionManager.Push(key)

		// Search cache data
		index := imc.searchIndexByKey(deletedKey)

		// Remove cache data
		if index != IN_MEMORY_CACHE_KEY_INDEX_NOT_FOUND {
			imc.CacheValues = append(imc.CacheValues[:index], imc.CacheValues[index+1:]...)
			imc.CacheKeys = append(imc.CacheKeys[:index], imc.CacheKeys[index+1:]...)
		} else {
			panic(fmt.Sprintf("Cache data with key %s not found", deletedKey))
		}
	}

	// Add new cache
	imc.CacheValues = append(imc.CacheValues, value)
	imc.CacheKeys = append(imc.CacheKeys, key)

	return IN_MEMORY_CACHE_ADD_NEW_KEY
}

func (imc *inMemoryCache) Get(key string) string {
	// Search data based on key
	for i:=0; i< len(imc.CacheKeys); i++ {
		if imc.CacheKeys[i] == key {
			// Since eviction manager Push function can re arrange key order, we need to call the it
			imc.EvictionManager.Push(key)

			return imc.CacheValues[i]
		}
	}

	return ""
}

func (imc *inMemoryCache) Clear() (totalData int) {
	// Remove all eviction manager Keys
	totalData = imc.EvictionManager.Clear()

	// Remove all cache data
	imc.CacheValues = []string{}
	imc.CacheKeys = []string{}

	return
}

func (imc *inMemoryCache) Keys() (keys []string) {
	return imc.CacheKeys
}

