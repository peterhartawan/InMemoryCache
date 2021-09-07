package none

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

func (imc *inMemoryCache) Add(key, value string) int {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("Recovering from panic: %v \n", r)
		}
	}()

	// Check key already exists
	isExist := imc.EvictionManager.Push(key)

	// Update data if exists
	if isExist == EVICTION_MANAGER_KEY_EXISTS {
		imc.CacheValues = append(imc.CacheValues, value)

		return IN_MEMORY_CACHE_REPLACE_EXISTING_KEY
	}

	// Panic if exceed Limit
	if len(imc.CacheKeys) == imc.Limit {
		// Remove latest key
		imc.EvictionManager.Pop()

		// Throw panic
		panic("key_limit_exceeded")
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
			return imc.CacheValues[i]
		}
	}

	return ""
}

func (imc *inMemoryCache) Clear() (totalData int) {
	// Remove all eviction manager Keys
	totalData = imc.EvictionManager.Clear()

	// Make cache data empty
	imc.CacheValues = nil
	imc.CacheKeys = nil

	return
}

func (imc *inMemoryCache) Keys() (keys []string) {
	return imc.CacheKeys
}
