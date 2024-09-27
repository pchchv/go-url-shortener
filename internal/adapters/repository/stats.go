package repository

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type StatsRepository struct {
	client    *dynamodb.Client
	tableName string
}

func NewStatsRepository(ctx context.Context, tableName string) *StatsRepository {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	client := dynamodb.NewFromConfig(cfg)
	return &StatsRepository{
		client:    client,
		tableName: tableName,
	}
}
