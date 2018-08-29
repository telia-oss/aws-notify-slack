package main

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/parnurzeal/gorequest"
	"github.com/telia-oss/aws-notify-slack/slack"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(snsEvent events.SNSEvent) {
	slackMessageAttachments := slack.CreateSlackMessagAttachment(snsEvent)

	slackHook := os.Getenv("SLACK_HOOK")
	request := gorequest.New()
	request.
		Post(slackHook).
		Send(slackMessageAttachments).
		End()
}

func main() {
	lambda.Start(Handler)
}
