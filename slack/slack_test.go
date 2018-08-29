package slack

import (
	"testing"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

var snsEntity = events.SNSEntity{
	SignatureVersion: "1",
	Timestamp:        time.Now(),
	Signature:        "EXAMPLE",
	SigningCertURL:   "EXAMPLE",
	MessageID:        "95df01b4-ee98-5cb9-9903-4c221d41eb5e",
	Message:          "{\"AlarmName\":\"sns-cloudwatch\",\"AlarmDescription\":null,\"AWSAccountId\":\"123456789012\",\"NewStateValue\":\"OK\",\"NewStateReason\":\"Threshold Crossed: 1 datapoint (7.9053535353535365) was not greater than or equal to the threshold (8.0).\",\"StateChangeTime\":\"2015-11-09T21:19:43.454+0000\",\"Region\":\"US - N. Virginia\",\"OldStateValue\":\"ALARM\",\"Trigger\":{\"MetricName\":\"CPUUtilization\",\"Namespace\":\"AWS/EC2\",\"Statistic\":\"AVERAGE\",\"Unit\":null,\"Dimensions\":[],\"Period\":300,\"EvaluationPeriods\":1,\"ComparisonOperator\":\"GreaterThanOrEqualToThreshold\",\"Threshold\":8.0}}",
	Type:             "Notification",
	UnsubscribeURL:   "EXAMPLE",
	TopicArn:         "arn:aws:sns:eu-west-1:000000000000:cloudwatch-alarms",
	Subject:          "example subject",
}
var snseventRecord = []events.SNSEventRecord{
	events.SNSEventRecord{
		EventVersion:         "1.0",
		EventSubscriptionArn: "arn:aws:sns:EXAMPLE",
		EventSource:          "aws:sns",
		SNS:                  snsEntity,
	},
}
var testEvent = events.SNSEvent{
	Records: snseventRecord,
}

func TestCreateSlackMessagAttachment(t *testing.T) {
	slackMessageAttachments := CreateSlackMessagAttachment(testEvent)
	attachemntsFields := []slackAttachmentField{
		slackAttachmentField{
			Title: "Alarm",
			Value: "sns-cloudwatch",
			Short: true,
		},
		slackAttachmentField{
			Title: "Status",
			Value: "OK",
			Short: true,
		},
		slackAttachmentField{
			Title: "Reason",
			Value: "Threshold Crossed: 1 datapoint (7.9053535353535365) was not greater than or equal to the threshold (8.0).",
			Short: false,
		},
	}

	assert.Equal(t, "good", slackMessageAttachments.Color)
	assert.Equal(t, "OK: sns-cloudwatch in US - N. Virginia", slackMessageAttachments.Pretext)
	assert.Equal(t, "AWS-bot", slackMessageAttachments.Username)
	assert.Equal(t, ":loudspeaker:", slackMessageAttachments.Icon)
	assert.Equal(t, attachemntsFields, slackMessageAttachments.Fields)
}
