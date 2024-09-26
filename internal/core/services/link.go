package services

import "github.com/pchchv/go-url-shortener/internal/core/ports"

type LinkService struct {
	port  ports.LinkPort
	cache ports.Cache
}

func NewLinkService(p ports.LinkPort, c ports.Cache) *LinkService {
	return &LinkService{port: p, cache: c}
}
