package handlers

import "github.com/pchchv/go-url-shortener/internal/core/services"

type RequestBody struct {
	Long string `json:"long"`
}

type GenerateLinkFunctionHandler struct {
	linkService  *services.LinkService
	statsService *services.StatsService
}

func NewGenerateLinkFunctionHandler(l *services.LinkService, s *services.StatsService) *GenerateLinkFunctionHandler {
	return &GenerateLinkFunctionHandler{linkService: l, statsService: s}
}
