package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pchchv/go-url-shortener/internal/adapters/cache"
	"github.com/pchchv/go-url-shortener/internal/adapters/handlers"
	"github.com/pchchv/go-url-shortener/internal/adapters/repository"
	"github.com/pchchv/go-url-shortener/internal/config"
	"github.com/pchchv/go-url-shortener/internal/core/services"
)

func main() {
	appConfig := config.NewConfig()
	redisAddress, redisPassword, redisDB := appConfig.GetRedisParams()
	linkTableName := appConfig.GetLinkTableName()
	statsTableName := appConfig.GetStatsTableName()

	cache := cache.NewRedisCache(redisAddress, redisPassword, redisDB)

	linkRepo := repository.NewLinkRepository(context.TODO(), linkTableName)
	statsRepo := repository.NewStatsRepository(context.TODO(), statsTableName)

	linkService := services.NewLinkService(linkRepo, cache)
	statsService := services.NewStatsService(statsRepo, cache)

	handler := handlers.NewDeleteFunctionHandler(linkService, statsService)

	lambda.Start(handler.Delete)
}
