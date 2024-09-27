package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	ddbtypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/pchchv/go-url-shortener/internal/core/domain"
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

func (d *StatsRepository) All(ctx context.Context) ([]domain.Stats, error) {
	input := &dynamodb.ScanInput{
		TableName: &d.tableName,
	}

	result, err := d.client.Scan(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to scan table: %w", err)
	}

	stats := []domain.Stats{}
	if err = attributevalue.UnmarshalListOfMaps(result.Items, &stats); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return stats, nil
}

func (d *StatsRepository) Get(ctx context.Context, id string) (domain.Stats, error) {
	input := &dynamodb.GetItemInput{
		TableName: &d.tableName,
		Key: map[string]ddbtypes.AttributeValue{
			"id": &ddbtypes.AttributeValueMemberS{Value: id},
		},
	}

	result, err := d.client.GetItem(ctx, input)
	if err != nil {
		return domain.Stats{}, fmt.Errorf("failed to get item from DynamoDB: %w", err)
	}

	stats := domain.Stats{}
	if err = attributevalue.UnmarshalMap(result.Item, &stats); err != nil {
		return domain.Stats{}, fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return stats, nil
}

func (d *StatsRepository) Create(ctx context.Context, stats domain.Stats) error {
	item, err := attributevalue.MarshalMap(stats)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	input := &dynamodb.PutItemInput{
		TableName: &d.tableName,
		Item:      item,
	}
	if _, err = d.client.PutItem(ctx, input); err != nil {
		return fmt.Errorf("failed to put item to DynamoDB: %w", err)
	}

	return nil
}

func (d *StatsRepository) Delete(ctx context.Context, id string) error {
	input := &dynamodb.DeleteItemInput{
		TableName: &d.tableName,
		Key: map[string]ddbtypes.AttributeValue{
			"id": &ddbtypes.AttributeValueMemberS{Value: id},
		},
	}
	if _, err := d.client.DeleteItem(ctx, input); err != nil {
		return fmt.Errorf("failed to delete item from DynamoDB: %w", err)
	}

	return nil
}
