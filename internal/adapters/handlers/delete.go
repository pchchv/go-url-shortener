package handlers

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/pchchv/go-url-shortener/internal/core/services"
)

type DeleteFunctionHandler struct {
	statsService *services.StatsService
	linkService  *services.LinkService
}

func NewDeleteFunctionHandler(l *services.LinkService, s *services.StatsService) *DeleteFunctionHandler {
	return &DeleteFunctionHandler{linkService: l, statsService: s}
}

func (s *DeleteFunctionHandler) Delete(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	id := req.PathParameters["id"]

	if err := s.linkService.Delete(ctx, id); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	if err := s.statsService.Delete(ctx, id); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	return events.APIGatewayProxyResponse{StatusCode: 204}, nil
}
