package services

import (
	"context"
	"fmt"

	"github.com/pchchv/go-url-shortener/internal/core/domain"
	"github.com/pchchv/go-url-shortener/internal/core/ports"
)

type LinkService struct {
	port  ports.LinkPort
	cache ports.Cache
}

func NewLinkService(p ports.LinkPort, c ports.Cache) *LinkService {
	return &LinkService{port: p, cache: c}
}

func (service *LinkService) GetAll(ctx context.Context) (links []domain.Link, err error) {
	if links, err = service.port.All(ctx); err != nil {
		return nil, fmt.Errorf("failed to get all links: %w", err)
	}
	return links, nil
}

func (service *LinkService) GetOriginalURL(ctx context.Context, shortLinkKey string) (*string, error) {
	data, err := service.port.Get(ctx, shortLinkKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get short URL for identifier '%s': %w", shortLinkKey, err)
	}
	return &data.OriginalURL, nil
}

func (service *LinkService) Create(ctx context.Context, link domain.Link) error {
	if err := service.port.Create(ctx, link); err != nil {
		return fmt.Errorf("failed to create short URL: %w", err)
	}
	return nil
}

func (service *LinkService) Delete(ctx context.Context, short string) error {
	if err := service.port.Delete(ctx, short); err != nil {
		return fmt.Errorf("failed to delete short URL for identifier '%s': %w", short, err)
	}
	return nil
}
