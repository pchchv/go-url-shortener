package handlers

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/pchchv/go-url-shortener/internal/config"
	"github.com/slack-go/slack"
)

func PostMessageToSlack(ctx context.Context, message string) error {
	appConfig := config.NewConfig()
	slackToken, slackChannelID := appConfig.GetSlackParams()
	api := slack.New(slackToken)
	channelID, timestamp, err := api.PostMessage(
		slackChannelID,
		slack.MsgOptionText(message, false),
	)
	if err != nil {
		log.Printf("Error posting to Slack: %s", err)
		return err
	}
	log.Printf("Message successfully sent to Slack channel %s at %s", channelID, timestamp)
	return nil
}

func HandleAPIGatewayRequest(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	if err := PostMessageToSlack(ctx, "Hello world! API Gateway message."); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Message successfully sent to Slack",
	}, nil
}

func HandleSQSMessage(ctx context.Context, message events.SQSMessage) error {
	return PostMessageToSlack(ctx, message.Body)
}
