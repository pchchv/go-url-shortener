package benchmark

import (
	"github.com/pchchv/go-url-shortener/internal/adapters/cache"
	"github.com/pchchv/go-url-shortener/internal/core/services"
	"github.com/pchchv/go-url-shortener/internal/tests/mock"
)

func GetService() *services.LinkService {
	cache := cache.NewRedisCache("localhost:6379", "", 0)
	mockLinkRepo := mock.NewMockLinkRepo()
	linkService := services.NewLinkService(mockLinkRepo, cache)
	return linkService
}
