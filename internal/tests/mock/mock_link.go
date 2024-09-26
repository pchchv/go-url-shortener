package mock

import "github.com/pchchv/go-url-shortener/internal/core/domain"

type MockLinkRepo struct {
	Links []domain.Link
	Stats []domain.Stats
}

func NewMockLinkRepo() *MockLinkRepo {
	return &MockLinkRepo{
		Links: MockLinkData,
		Stats: MockStatsData,
	}
}
