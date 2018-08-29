package slack

import (
	"fmt"
	"os"

	"github.com/Jeffail/gabs"
	"github.com/aws/aws-lambda-go/events"
)

type slackAttachmentField struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
	Short bool   `json:"short,omitempty"`
}

// MessageAttachments is the Slack message stracture
type MessageAttachments struct {
	Color    string                 `json:"color,omitempty"`
	Pretext  string                 `json:"pretext,omitempty"`
	Username string                 `json:"username,omitempty"`
	Icon     string                 `json:"icon_emoji,omitempty"`
	Fields   []slackAttachmentField `json:"fields,omitempty"`
}

func mapColor(status string) string {
	var colorCode string
	switch status {
	case "ALARM":
		colorCode = "danger"
	case "INSUFFICIENT_DATA":
		colorCode = "warning"
	default:
		colorCode = "good"
	}

	return colorCode
}

// CreateSlackMessagAttachment is a function to create slack message
func CreateSlackMessagAttachment(snsEvent events.SNSEvent) MessageAttachments {
	records := snsEvent.Records
	snsRecord := records[0].SNS

	jsonParsed, _ := gabs.ParseJSON([]byte(snsRecord.Message))
	NewStateValue, _ := jsonParsed.Path("NewStateValue").Data().(string)
	NewStateReason, _ := jsonParsed.Path("NewStateReason").Data().(string)
	AlarmName, _ := jsonParsed.Path("AlarmName").Data().(string)
	Region, _ := jsonParsed.Path("Region").Data().(string)

	slackAttachmentFields := []slackAttachmentField{
		slackAttachmentField{
			Title: "Alarm",
			Value: AlarmName,
			Short: true,
		},
		slackAttachmentField{
			Title: "Status",
			Value: NewStateValue,
			Short: true,
		},
		slackAttachmentField{
			Title: "Reason",
			Value: NewStateReason,
			Short: false,
		},
	}

	pretext := fmt.Sprintf("%s: %s in %s", NewStateValue, AlarmName, Region)

	username := os.Getenv("USERNAME")
	if username == "" {
		username = "AWS-bot"
	}

	icon := os.Getenv("ICON")
	if icon == "" {
		icon = ":loudspeaker:"
	}

	slackMessageAttachments := MessageAttachments{
		Color:    mapColor(NewStateValue),
		Pretext:  pretext,
		Username: username,
		Icon:     icon,
		Fields:   slackAttachmentFields,
	}

	return slackMessageAttachments
}
