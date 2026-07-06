package config

import (
	"fmt"
	"sync"
	"time"
)

type MemoryCache struct {
	mu   sync.RWMutex
	data map[string]item
}

type item struct {
	value  any
	expiry time.Time
}

func NewMemoryCache() *MemoryCache {
	memoryCache := &MemoryCache{data: make(map[string]item)}
	go memoryCache.cleanup()
	return memoryCache
}

// memory cleanup job runs every 2 min

func (m *MemoryCache) cleanup() {
	ticker := time.NewTicker(2 * time.Minute)
	defer ticker.Stop()
	for range ticker.C {
		m.mu.Lock()
		for key, item := range m.data {
			if time.Now().After(item.expiry) {
				delete(m.data, key)
			}
		}
		m.mu.Unlock()
	}
}

// memory cache methods

func (m *MemoryCache) SetUserAccessToken(userID int, token string, duration time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[fmt.Sprintf("user:%d:token", userID)] = item{
		value:  HashToken(token),
		expiry: time.Now().Add(duration),
	}
}

func (m *MemoryCache) GetUserAccessToken(userID int) string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	item, exists := m.data[fmt.Sprintf("user:%d:token", userID)]
	if !exists || time.Now().After(item.expiry) {
		return ""
	}
	return item.value.(string)
}

func (m *MemoryCache) DeleteUserAccessToken(userID int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.data, fmt.Sprintf("user:%d:token", userID))
}
