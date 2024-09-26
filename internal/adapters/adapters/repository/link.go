package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	ddbtypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/pchchv/go-url-shortener/internal/core/domain"
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

func (d *LinkRepository) All(ctx context.Context) ([]domain.Link, error) {
	var links []domain.Link
	input := &dynamodb.ScanInput{
		TableName: &d.tableName,
		Limit:     aws.Int32(20),
	}

	result, err := d.client.Scan(ctx, input)
	if err != nil {
		return links, fmt.Errorf("failed to get items from DynamoDB: %w", err)
	}

	if err = attributevalue.UnmarshalListOfMaps(result.Items, &links); err != nil {
		return links, fmt.Errorf("failed to unmarshal data from DynamoDB: %w", err)
	}

	return links, nil
}

func (d *LinkRepository) Get(ctx context.Context, id string) (domain.Link, error) {
	link := domain.Link{}
	input := &dynamodb.GetItemInput{
		TableName: &d.tableName,
		Key: map[string]ddbtypes.AttributeValue{
			"id": &ddbtypes.AttributeValueMemberS{Value: id},
		},
	}

	result, err := d.client.GetItem(ctx, input)
	if err != nil {
		return link, fmt.Errorf("failed to get item from DynamoDB: %w", err)
	}

	if err = attributevalue.UnmarshalMap(result.Item, &link); err != nil {
		return link, fmt.Errorf("failed to unmarshal data from DynamoDB: %w", err)
	}

	return link, nil
}
