package mock

import "time"

type MockRedisCache struct {
	Store map[string]string
	TTL   map[string]time.Time
}

func NewMockRedisCache() *MockRedisCache {
	return &MockRedisCache{
		Store: make(map[string]string),
		TTL:   make(map[string]time.Time),
	}
}
