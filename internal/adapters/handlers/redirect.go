package handlers

import "github.com/pchchv/go-url-shortener/internal/core/services"

type RedirectFunctionHandler struct {
	linkService  *services.LinkService
	statsService *services.StatsService
}

func NewRedirectFunctionHandler(l *services.LinkService, s *services.StatsService) *RedirectFunctionHandler {
	return &RedirectFunctionHandler{linkService: l, statsService: s}
}
