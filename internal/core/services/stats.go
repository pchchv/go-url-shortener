package services

import "github.com/pchchv/go-url-shortener/internal/core/ports"

type StatsService struct {
	port  ports.StatsPort
	cache ports.Cache
}

func NewStatsService(p ports.StatsPort, c ports.Cache) *StatsService {
	return &StatsService{port: p, cache: c}
}
