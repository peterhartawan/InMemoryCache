package none

type EvictionManager struct {
	Keys []string
}

func (em *EvictionManager) Push(key string) int {
	// Search is key exists
	for _, v := range em.Keys {
		if v == key {
			return EVICTION_MANAGER_KEY_EXISTS
		}
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
