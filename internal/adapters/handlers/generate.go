package handlers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/pchchv/go-url-shortener/internal/core/domain"
	"github.com/pchchv/go-url-shortener/internal/core/services"
)

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

func sendMessageToQueue(ctx context.Context, link domain.Link) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err.Error())
		return
	}

	sqsClient := sqs.NewFromConfig(cfg)
	queueUrl := os.Getenv("QueueUrl")
	if queueUrl == "" {
		log.Println("QueueUrl is not set")
		return
	}

	_, err = sqsClient.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    &queueUrl,
		MessageBody: aws.String("The system generated a short URL with the ID " + link.Id),
	})
	if err != nil {
		fmt.Printf("Failed to send message to SQS, %v", err.Error())
	}
}
