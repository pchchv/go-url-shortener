package unit

import (
	"context"

	"github.com/pchchv/go-url-shortener/internal/adapters/cache"
	"github.com/pchchv/go-url-shortener/internal/core/domain"
)

func FillCache(cache *cache.RedisCache, links []domain.Link) (err error) {
	for _, link := range links {
		if err = cache.Set(context.Background(), link.Id, link.OriginalURL); err != nil {
			break
		}
	}
	return
}
