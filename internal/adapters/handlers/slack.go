package handlers

import (
	"context"
	"log"

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
