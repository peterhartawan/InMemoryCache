package lru

type EvictionManager struct {
	Keys []string
}

func (em *EvictionManager) findIndexBySearch(key string) (index int) {
	index = EVICTION_MANAGER_CACHE_KEY_INDEX_NOT_FOUND

	for k, v := range em.Keys {
		if v == key {
			index = k
		}
	}

	return
}

func (em *EvictionManager) reOrderKey(index int, key string) {
	// Delete key
	em.Keys = append(em.Keys[:index], em.Keys[index+1:]...)

	// Put to the first data
	em.Keys = append([]string{key}, em.Keys...)
}

func (em *EvictionManager) Push(key string) int {
	// Search is key exists
	index := em.findIndexBySearch(key)

	// If exist move new key to front
	if index != EVICTION_MANAGER_CACHE_KEY_INDEX_NOT_FOUND {
		em.reOrderKey(index, key)

		return EVICTION_MANAGER_KEY_EXISTS
	}

	// Add new key
	em.Keys = append(em.Keys, key)

	return EVICTION_MANAGER_KEY_NOT_EXISTS
}

func (em *EvictionManager) Pop() (key string) {
	// Get latest key
	key = em.Keys[len(em.Keys)-1]

	// Remove latest key
	em.Keys = em.Keys[:len(em.Keys)-1]

	return
}

func (em *EvictionManager) Clear() (totalKeys int) {
	// Get total Keys
	totalKeys = len(em.Keys)

	// Make key empty
	em.Keys = nil

	return
}
