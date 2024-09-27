package handlers

import "github.com/pchchv/go-url-shortener/internal/core/services"

type DeleteFunctionHandler struct {
	statsService *services.StatsService
	linkService  *services.LinkService
}

func NewDeleteFunctionHandler(l *services.LinkService, s *services.StatsService) *DeleteFunctionHandler {
	return &DeleteFunctionHandler{linkService: l, statsService: s}
}
