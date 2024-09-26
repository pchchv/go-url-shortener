package services

import "github.com/pchchv/go-url-shortener/internal/core/ports"

type LinkService struct {
	port  ports.LinkPort
	cache ports.Cache
}
