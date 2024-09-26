package mock

import (
	"context"
	"errors"
	"time"
)

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

func (m *MockRedisCache) Delete(ctx context.Context, key string) error {
	if _, ok := m.Store[key]; !ok {
		return errors.New("key not found")
	}

	delete(m.Store, key)
	delete(m.TTL, key)
	return nil
}
