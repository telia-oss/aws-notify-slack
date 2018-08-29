package main

import (
	"os"

	"github.com/aws-notify-slack/slack"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/parnurzeal/gorequest"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(snsEvent events.SNSEvent) {
	slackMessageAttachments := slack.CreateSlackMessagAttachment(snsEvent)

	slackHook := os.Getenv("SLACK_HOOK")
	if slackHook == "" {
		slackHook = "https://hooks.slack.com/services/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"
	}
	request := gorequest.New()
	request.
		Post(slackHook).
		Send(slackMessageAttachments).
		End()
}

func main() {
	lambda.Start(Handler)
}
