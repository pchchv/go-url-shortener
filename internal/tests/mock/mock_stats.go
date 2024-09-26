package mock

import "github.com/pchchv/go-url-shortener/internal/core/domain"

type MockStatsRepo struct {
	Stats []domain.Stats
}

func NewMockStatsRepo() *MockStatsRepo {
	return &MockStatsRepo{
		Stats: MockStatsData,
	}
}
