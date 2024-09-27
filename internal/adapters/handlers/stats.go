package handlers

import "github.com/pchchv/go-url-shortener/internal/core/services"

type StatsFunctionHandler struct {
	statsService *services.StatsService
	linkService  *services.LinkService
}

func NewStatsFunctionHandler(l *services.LinkService, s *services.StatsService) *StatsFunctionHandler {
	return &StatsFunctionHandler{linkService: l, statsService: s}
}
