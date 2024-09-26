package repository

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type LinkRepository struct {
	client    *dynamodb.Client
	tableName string
}

func NewLinkRepository(ctx context.Context, tableName string) *LinkRepository {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	client := dynamodb.NewFromConfig(cfg)
	return &LinkRepository{
		client:    client,
		tableName: tableName,
	}
}
