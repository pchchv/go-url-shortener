package mock

import (
	"context"

	"github.com/pchchv/go-url-shortener/internal/core/domain"
)

type MockStatsRepo struct {
	Stats []domain.Stats
}

func NewMockStatsRepo() *MockStatsRepo {
	return &MockStatsRepo{
		Stats: MockStatsData,
	}
}

func (m *MockStatsRepo) All(ctx context.Context) ([]domain.Stats, error) {
	return m.Stats, nil
}

func (m *MockStatsRepo) Get(ctx context.Context, id string) (domain.Stats, error) {
	for _, stats := range m.Stats {
		if stats.Id == id {
			return stats, nil
		}
	}
	return domain.Stats{}, nil
}
