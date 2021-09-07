package lfu

type EvictionManager struct {
	Keys          []string
	AccessCounter []int
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
	for i := 0; i < len(em.AccessCounter); i++ {
		for j := 0; j < len(em.AccessCounter); j++ {
			// swap if counter first data lesser than second data
			if (em.AccessCounter[i] < em.AccessCounter[j]) {
				em.AccessCounter[i], em.AccessCounter[j] = em.AccessCounter[j], em.AccessCounter[i]
				em.Keys[i], em.Keys[j] = em.Keys[j], em.Keys[i]
			}
		}
	}
}

func (em *EvictionManager) Push(key string) int {
	// Search is key exists
	index := em.findIndexBySearch(key)

	// If exist add counter and re-order key
	if index != EVICTION_MANAGER_CACHE_KEY_INDEX_NOT_FOUND {
		em.AccessCounter[index]++

		em.reOrderKey(index, key)

		return EVICTION_MANAGER_KEY_EXISTS
	}

	// Add new key
	em.Keys = append(em.Keys, key)
	em.AccessCounter = append(em.AccessCounter,1)

	return EVICTION_MANAGER_KEY_NOT_EXISTS
}

func (em *EvictionManager) Pop() (key string) {
	// Get latest key
	key = em.Keys[len(em.Keys)-1]

	// Remove latest key
	em.Keys = em.Keys[:len(em.Keys)-1]
	em.AccessCounter = em.AccessCounter[:len(em.AccessCounter)-1]

	return
}

func (em *EvictionManager) Clear() (totalKeys int) {
	// Get total Keys
	totalKeys = len(em.Keys)

	// Make key empty
	em.Keys = nil
	em.AccessCounter = nil

	return
}
