package slack

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

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

func mapAlarmColor(status string) string {
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

func mapEcsTaskChangeColor(desiredStatus string) string {
	var colorCode string
	switch desiredStatus {
	case "STOPPED":
		colorCode = "danger"
	default:
		colorCode = "good"
	}

	return colorCode
}

func ecsTaskStateChange(message *gabs.Container) string {
	detail := message.Path("detail")

	clusterArn, _ := detail.Path("clusterArn").Data().(string)
	desiredStatus, _ := detail.Path("desiredStatus").Data().(string)
	lastStatus, _ := detail.Path("lastStatus").Data().(string)
	stoppedReason, _ := detail.Path("stoppedReason").Data().(string)
	taskArn, _ := detail.Path("taskArn").Data().(string)
	taskDefinitionArn, _ := detail.Path("taskDefinitionArn").Data().(string)

	clusterName := clusterArn[strings.LastIndex(clusterArn, "/")+1:]
	taskName := taskArn[strings.LastIndex(taskArn, "/")+1:]
	taskDefinitionName := taskDefinitionArn[strings.LastIndex(taskDefinitionArn, "/")+1:]

	slackAttachmentFields := []slackAttachmentField{
		{
			Title: "Last status",
			Value: lastStatus,
			Short: true,
		},
		{
			Title: "Desired status",
			Value: desiredStatus,
			Short: true,
		},
		{
			Title: "Cluster",
			Value: clusterName,
			Short: true,
		},
		{
			Title: "Task definition",
			Value: taskDefinitionName,
			Short: true,
		},
		{
			Title: "Task",
			Value: taskName,
			Short: true,
		},
	}

	if stoppedReason != "" {
		slackAttachmentFields = append(slackAttachmentFields, slackAttachmentField{
			Title: "Stopped reason",
			Value: stoppedReason,
			Short: true,
		})
	}

	pretext := fmt.Sprintf("Task %s in %s cluster is changing state: %s -> %s", taskDefinitionName, clusterName, lastStatus, desiredStatus)

	if lastStatus == desiredStatus {
		pretext = fmt.Sprintf("Task %s in %s cluster changed state: %s", taskDefinitionName, clusterName, lastStatus)
	}

	username := os.Getenv("USERNAME")
	if username == "" {
		username = "AWS-bot"
	}

	icon := os.Getenv("ICON")
	if icon == "" {
		icon = ":loudspeaker:"
	}

	slackMessageAttachments := MessageAttachments{
		Color:    mapEcsTaskChangeColor(desiredStatus),
		Pretext:  pretext,
		Username: username,
		Icon:     icon,
		Fields:   slackAttachmentFields,
	}

	resp, err := json.Marshal(slackMessageAttachments)
	if err != nil {
		log.Fatal("Error building Slack attachments", err)
	}

	return string(resp)
}

func alarm(message *gabs.Container) string {
	NewStateValue, _ := message.Path("NewStateValue").Data().(string)
	NewStateReason, _ := message.Path("NewStateReason").Data().(string)
	AlarmName, _ := message.Path("AlarmName").Data().(string)
	Region, _ := message.Path("Region").Data().(string)

	slackAttachmentFields := []slackAttachmentField{
		{
			Title: "Alarm",
			Value: AlarmName,
			Short: true,
		},
		{
			Title: "Status",
			Value: NewStateValue,
			Short: true,
		},
		{
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
		Color:    mapAlarmColor(NewStateValue),
		Pretext:  pretext,
		Username: username,
		Icon:     icon,
		Fields:   slackAttachmentFields,
	}

	resp, err := json.Marshal(slackMessageAttachments)
	if err != nil {
		log.Fatal("Error building Slack attachments", err)
	}

	return string(resp)
}

// CreateSlackMessagAttachment is a function to create slack message
func CreateSlackMessageAttachment(snsEvent events.SNSEvent) string {
	log.Println("snsEvent", snsEvent)
	records := snsEvent.Records
	snsRecord := records[0].SNS

	message, _ := gabs.ParseJSON([]byte(snsRecord.Message))

	if message.Exists("detail-type") && message.Path("detail-type").Data().(string) == "ECS Task State Change" {
		return ecsTaskStateChange(message)
	}

	if message.Exists("AlarmName") {
		return alarm(message)
	}

	return ""
}
