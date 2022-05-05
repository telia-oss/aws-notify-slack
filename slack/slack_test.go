package slack

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

var testAlarmEvent = events.SNSEvent{
	Records: []events.SNSEventRecord{
		{
			EventVersion:         "1.0",
			EventSubscriptionArn: "arn:aws:sns:EXAMPLE",
			EventSource:          "aws:sns",
			SNS: events.SNSEntity{
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
			},
		},
	},
}

var deactivatingStoppedEcsTaskEvent = events.SNSEvent{
	Records: []events.SNSEventRecord{
		{
			EventVersion:         "1.0",
			EventSubscriptionArn: "arn:aws:sns:EXAMPLE",
			EventSource:          "aws:sns",
			SNS: events.SNSEntity{
				SignatureVersion: "1",
				Timestamp:        time.Now(),
				Signature:        "EXAMPLE",
				SigningCertURL:   "EXAMPLE",
				MessageID:        "95df01b4-ee98-5cb9-9903-4c221d41eb5e",
				Message:          "{\"version\":\"0\",\"id\":\"b99683ac-c96e-a875-fa84-d3a50e04f78b\",\"detail-type\":\"ECS Task State Change\",\"source\":\"aws.ecs\",\"account\":\"123456789000\",\"time\":\"2022-05-03T07:29:20Z\",\"region\":\"eu-west-1\",\"resources\":[\"arn:aws:ecs:eu-west-1:123456789000:task/service/12345\"],\"detail\":{\"attachments\":[{\"id\":\"6e0c1c6e-180f-449a-b23c-92c1bef62f4c\",\"type\":\"sdi\",\"status\":\"ATTACHED\",\"details\":[]},{\"id\":\"123\",\"type\":\"eni\",\"status\":\"ATTACHED\",\"details\":[{\"name\":\"subnetId\",\"value\":\"subnet-12345\"},{\"name\":\"networkInterfaceId\",\"value\":\"eni-12345\"},{\"name\":\"macAddress\",\"value\":\"12345\"},{\"name\":\"privateDnsName\",\"value\":\"xzy.eu-west-1.compute.internal\"},{\"name\":\"privateIPv4Address\",\"value\":\"127.0.0.1\"}]}],\"attributes\":[{\"name\":\"ecs.cpu-architecture\",\"value\":\"x86_64\"}],\"availabilityZone\":\"eu-west-1b\",\"clusterArn\":\"arn:aws:ecs:eu-west-1:123456789000:cluster/service\",\"connectivity\":\"CONNECTED\",\"connectivityAt\":\"2022-05-02T10:48:55.655Z\",\"containers\":[{\"containerArn\":\"arn:aws:ecs:eu-west-1:123456789000:container/service/12345/1234\",\"lastStatus\":\"RUNNING\",\"name\":\"service\",\"image\":\"123456789000.dkr.ecr.eu-west-1.amazonaws.com/image-name:latest\",\"imageDigest\":\"sha256:123\",\"runtimeId\":\"12345-123\",\"taskArn\":\"arn:aws:ecs:eu-west-1:123456789000:task/service/12345\",\"networkInterfaces\":[{\"attachmentId\":\"123\",\"privateIpv4Address\":\"127.0.0.1\"}],\"cpu\":\"0\"}],\"cpu\":\"256\",\"createdAt\":\"2022-05-02T10:48:52.289Z\",\"desiredStatus\":\"STOPPED\",\"enableExecuteCommand\":false,\"ephemeralStorage\":{\"sizeInGiB\":20},\"group\":\"service:service\",\"launchType\":\"FARGATE\",\"lastStatus\":\"DEACTIVATING\",\"memory\":\"512\",\"overrides\":{\"containerOverrides\":[{\"name\":\"service\"}]},\"platformVersion\":\"1.4.0\",\"pullStartedAt\":\"2022-05-02T10:49:03.653Z\",\"pullStoppedAt\":\"2022-05-02T10:49:05.836Z\",\"startedAt\":\"2022-05-02T10:49:50.416Z\",\"startedBy\":\"ecs-svc/123\",\"stoppingAt\":\"2022-05-03T07:29:20.98Z\",\"stoppedReason\":\"Task stopped by user\",\"stopCode\":\"UserInitiated\",\"taskArn\":\"arn:aws:ecs:eu-west-1:123456789000:task/service/12345\",\"taskDefinitionArn\":\"arn:aws:ecs:eu-west-1:123456789000:task-definition/service:2\",\"updatedAt\":\"2022-05-03T07:29:20.98Z\",\"version\":5}}",
				Type:             "Notification",
				UnsubscribeURL:   "EXAMPLE",
				TopicArn:         "arn:aws:sns:eu-west-1:000000000000:cloudwatch-alarms",
				Subject:          "example subject",
			},
		},
	},
}

var provisioningRunningEcsTaskEvent = events.SNSEvent{
	Records: []events.SNSEventRecord{
		{
			EventVersion:         "1.0",
			EventSubscriptionArn: "arn:aws:sns:EXAMPLE",
			EventSource:          "aws:sns",
			SNS: events.SNSEntity{
				SignatureVersion: "1",
				Timestamp:        time.Now(),
				Signature:        "EXAMPLE",
				SigningCertURL:   "EXAMPLE",
				MessageID:        "95df01b4-ee98-5cb9-9903-4c221d41eb5e",
				Message:          "{\"version\":\"0\",\"id\":\"f597c1cc-1506-bce5-d509-69ef2633dc34\",\"detail-type\":\"ECS Task State Change\",\"source\":\"aws.ecs\",\"account\":\"123456789000\",\"time\":\"2022-05-03T07:29:51Z\",\"region\":\"eu-west-1\",\"resources\":[\"arn:aws:ecs:eu-west-1:123456789000:task/service/123\"],\"detail\":{\"attachments\":[{\"id\":\"e91b67e3-d01a-47f9-99ae-027b95c6b1ac\",\"type\":\"eni\",\"status\":\"PRECREATED\",\"details\":[{\"name\":\"subnetId\",\"value\":\"subnet-123\"}]},{\"id\":\"114de60f-dfad-4285-980c-35216ee9f614\",\"type\":\"sdi\",\"status\":\"PRECREATED\",\"details\":[]}],\"attributes\":[{\"name\":\"ecs.cpu-architecture\",\"value\":\"x86_64\"}],\"availabilityZone\":\"eu-west-1c\",\"clusterArn\":\"arn:aws:ecs:eu-west-1:123456789000:cluster/service\",\"containers\":[{\"containerArn\":\"arn:aws:ecs:eu-west-1:123456789000:container/service/123/123\",\"lastStatus\":\"PENDING\",\"name\":\"service\",\"image\":\"123456789000.dkr.ecr.eu-west-1.amazonaws.com/image-name:latest\",\"taskArn\":\"arn:aws:ecs:eu-west-1:123456789000:task/service/123\",\"networkInterfaces\":[],\"cpu\":\"0\"}],\"cpu\":\"256\",\"createdAt\":\"2022-05-03T07:29:51.723Z\",\"desiredStatus\":\"RUNNING\",\"enableExecuteCommand\":false,\"ephemeralStorage\":{\"sizeInGiB\":20},\"group\":\"service:service\",\"launchType\":\"FARGATE\",\"lastStatus\":\"PROVISIONING\",\"memory\":\"512\",\"overrides\":{\"containerOverrides\":[{\"name\":\"service\"}]},\"platformVersion\":\"1.4.0\",\"startedBy\":\"ecs-svc/123\",\"taskArn\":\"arn:aws:ecs:eu-west-1:123456789000:task/service/123\",\"taskDefinitionArn\":\"arn:aws:ecs:eu-west-1:123456789000:task-definition/service:2\",\"updatedAt\":\"2022-05-03T07:29:51.723Z\",\"version\":1}}",
				Type:             "Notification",
				UnsubscribeURL:   "EXAMPLE",
				TopicArn:         "arn:aws:sns:eu-west-1:000000000000:cloudwatch-alarms",
				Subject:          "example subject",
			},
		},
	},
}

var deprovisioniningStoppedEcsTaskEvent = events.SNSEvent{
	Records: []events.SNSEventRecord{
		{
			EventVersion:         "1.0",
			EventSubscriptionArn: "arn:aws:sns:EXAMPLE",
			EventSource:          "aws:sns",
			SNS: events.SNSEntity{
				SignatureVersion: "1",
				Timestamp:        time.Now(),
				Signature:        "EXAMPLE",
				SigningCertURL:   "EXAMPLE",
				MessageID:        "95df01b4-ee98-5cb9-9903-4c221d41eb5e",
				Message:          "{\"version\":\"0\",\"id\":\"52d64f87-0b04-cef3-4bf1-d8f5aa83e29d\",\"detail-type\":\"ECS Task State Change\",\"source\":\"aws.ecs\",\"account\":\"123456789000\",\"time\":\"2022-05-03T07:29:54Z\",\"region\":\"eu-west-1\",\"resources\":[\"arn:aws:ecs:eu-west-1:123456789000:task/service/12345\"],\"detail\":{\"attachments\":[{\"id\":\"6e0c1c6e-180f-449a-b23c-92c1bef62f4c\",\"type\":\"sdi\",\"status\":\"DELETED\",\"details\":[]},{\"id\":\"123\",\"type\":\"eni\",\"status\":\"ATTACHED\",\"details\":[{\"name\":\"subnetId\",\"value\":\"subnet-12345\"},{\"name\":\"networkInterfaceId\",\"value\":\"eni-12345\"},{\"name\":\"macAddress\",\"value\":\"12345\"},{\"name\":\"privateDnsName\",\"value\":\"xzy.eu-west-1.compute.internal\"},{\"name\":\"privateIPv4Address\",\"value\":\"127.0.0.1\"}]}],\"attributes\":[{\"name\":\"ecs.cpu-architecture\",\"value\":\"x86_64\"}],\"availabilityZone\":\"eu-west-1b\",\"clusterArn\":\"arn:aws:ecs:eu-west-1:123456789000:cluster/service\",\"connectivity\":\"CONNECTED\",\"connectivityAt\":\"2022-05-02T10:48:55.655Z\",\"containers\":[{\"containerArn\":\"arn:aws:ecs:eu-west-1:123456789000:container/service/12345/1234\",\"exitCode\":2,\"lastStatus\":\"STOPPED\",\"name\":\"service\",\"image\":\"123456789000.dkr.ecr.eu-west-1.amazonaws.com/image-name:latest\",\"imageDigest\":\"sha256:123\",\"runtimeId\":\"12345-123\",\"taskArn\":\"arn:aws:ecs:eu-west-1:123456789000:task/service/12345\",\"networkInterfaces\":[{\"attachmentId\":\"123\",\"privateIpv4Address\":\"127.0.0.1\"}],\"cpu\":\"0\"}],\"cpu\":\"256\",\"createdAt\":\"2022-05-02T10:48:52.289Z\",\"desiredStatus\":\"STOPPED\",\"enableExecuteCommand\":false,\"ephemeralStorage\":{\"sizeInGiB\":20},\"executionStoppedAt\":\"2022-05-03T07:29:44.856Z\",\"group\":\"service:service\",\"launchType\":\"FARGATE\",\"lastStatus\":\"DEPROVISIONING\",\"memory\":\"512\",\"overrides\":{\"containerOverrides\":[{\"name\":\"service\"}]},\"platformVersion\":\"1.4.0\",\"pullStartedAt\":\"2022-05-02T10:49:03.653Z\",\"pullStoppedAt\":\"2022-05-02T10:49:05.836Z\",\"startedAt\":\"2022-05-02T10:49:50.416Z\",\"startedBy\":\"ecs-svc/123\",\"stoppingAt\":\"2022-05-03T07:29:20.98Z\",\"stoppedReason\":\"Task stopped by user\",\"stopCode\":\"UserInitiated\",\"taskArn\":\"arn:aws:ecs:eu-west-1:123456789000:task/service/12345\",\"taskDefinitionArn\":\"arn:aws:ecs:eu-west-1:123456789000:task-definition/service:2\",\"updatedAt\":\"2022-05-03T07:29:54.884Z\",\"version\":7}}",
				Type:             "Notification",
				UnsubscribeURL:   "EXAMPLE",
				TopicArn:         "arn:aws:sns:eu-west-1:000000000000:cloudwatch-alarms",
				Subject:          "example subject",
			},
		},
	},
}

var pendingRunningEcsTaskEvent = events.SNSEvent{
	Records: []events.SNSEventRecord{
		{
			EventVersion:         "1.0",
			EventSubscriptionArn: "arn:aws:sns:EXAMPLE",
			EventSource:          "aws:sns",
			SNS: events.SNSEntity{
				SignatureVersion: "1",
				Timestamp:        time.Now(),
				Signature:        "EXAMPLE",
				SigningCertURL:   "EXAMPLE",
				MessageID:        "95df01b4-ee98-5cb9-9903-4c221d41eb5e",
				Message:          "{\"version\":\"0\",\"id\":\"05cd4a7e-f7df-60c6-e04a-799ce6078b88\",\"detail-type\":\"ECS Task State Change\",\"source\":\"aws.ecs\",\"account\":\"123456789000\",\"time\":\"2022-05-03T07:30:00Z\",\"region\":\"eu-west-1\",\"resources\":[\"arn:aws:ecs:eu-west-1:123456789000:task/service/123\"],\"detail\":{\"attachments\":[{\"id\":\"114de60f-dfad-4285-980c-35216ee9f614\",\"type\":\"sdi\",\"status\":\"PRECREATED\",\"details\":[]},{\"id\":\"e91b67e3-d01a-47f9-99ae-027b95c6b1ac\",\"type\":\"eni\",\"status\":\"ATTACHED\",\"details\":[{\"name\":\"subnetId\",\"value\":\"subnet-123\"},{\"name\":\"networkInterfaceId\",\"value\":\"eni-123\"},{\"name\":\"macAddress\",\"value\":\"123\"},{\"name\":\"privateDnsName\",\"value\":\"123.eu-west-1.compute.internal\"},{\"name\":\"privateIPv4Address\",\"value\":\"123\"}]}],\"attributes\":[{\"name\":\"ecs.cpu-architecture\",\"value\":\"x86_64\"}],\"availabilityZone\":\"eu-west-1c\",\"clusterArn\":\"arn:aws:ecs:eu-west-1:123456789000:cluster/service\",\"containers\":[{\"containerArn\":\"arn:aws:ecs:eu-west-1:123456789000:container/service/123/123\",\"lastStatus\":\"PENDING\",\"name\":\"service\",\"image\":\"123456789000.dkr.ecr.eu-west-1.amazonaws.com/image-name:latest\",\"taskArn\":\"arn:aws:ecs:eu-west-1:123456789000:task/service/123\",\"networkInterfaces\":[{\"attachmentId\":\"e91b67e3-d01a-47f9-99ae-027b95c6b1ac\",\"privateIpv4Address\":\"123\"}],\"cpu\":\"0\"}],\"cpu\":\"256\",\"createdAt\":\"2022-05-03T07:29:51.723Z\",\"desiredStatus\":\"RUNNING\",\"enableExecuteCommand\":false,\"ephemeralStorage\":{\"sizeInGiB\":20},\"group\":\"service:service\",\"launchType\":\"FARGATE\",\"lastStatus\":\"PENDING\",\"memory\":\"512\",\"overrides\":{\"containerOverrides\":[{\"name\":\"service\"}]},\"platformVersion\":\"1.4.0\",\"startedBy\":\"ecs-svc/123\",\"taskArn\":\"arn:aws:ecs:eu-west-1:123456789000:task/service/123\",\"taskDefinitionArn\":\"arn:aws:ecs:eu-west-1:123456789000:task-definition/service:2\",\"updatedAt\":\"2022-05-03T07:30:00.511Z\",\"version\":2}}",
				Type:             "Notification",
				UnsubscribeURL:   "EXAMPLE",
				TopicArn:         "arn:aws:sns:eu-west-1:000000000000:cloudwatch-alarms",
				Subject:          "example subject",
			},
		},
	},
}

var stoppedEcsTaskEvent = events.SNSEvent{
	Records: []events.SNSEventRecord{
		{
			EventVersion:         "1.0",
			EventSubscriptionArn: "arn:aws:sns:EXAMPLE",
			EventSource:          "aws:sns",
			SNS: events.SNSEntity{
				SignatureVersion: "1",
				Timestamp:        time.Now(),
				Signature:        "EXAMPLE",
				SigningCertURL:   "EXAMPLE",
				MessageID:        "95df01b4-ee98-5cb9-9903-4c221d41eb5e",
				Message:          "{\"version\":\"0\",\"id\":\"09824691-294a-8f91-0bd5-01e13c5a3c8d\",\"detail-type\":\"ECS Task State Change\",\"source\":\"aws.ecs\",\"account\":\"123456789000\",\"time\":\"2022-05-03T07:30:07Z\",\"region\":\"eu-west-1\",\"resources\":[\"arn:aws:ecs:eu-west-1:123456789000:task/service/12345\"],\"detail\":{\"attachments\":[{\"id\":\"6e0c1c6e-180f-449a-b23c-92c1bef62f4c\",\"type\":\"sdi\",\"status\":\"DELETED\",\"details\":[]},{\"id\":\"123\",\"type\":\"eni\",\"status\":\"DELETED\",\"details\":[{\"name\":\"subnetId\",\"value\":\"subnet-12345\"},{\"name\":\"networkInterfaceId\",\"value\":\"eni-12345\"},{\"name\":\"macAddress\",\"value\":\"12345\"},{\"name\":\"privateDnsName\",\"value\":\"xzy.eu-west-1.compute.internal\"},{\"name\":\"privateIPv4Address\",\"value\":\"127.0.0.1\"}]}],\"attributes\":[{\"name\":\"ecs.cpu-architecture\",\"value\":\"x86_64\"}],\"availabilityZone\":\"eu-west-1b\",\"clusterArn\":\"arn:aws:ecs:eu-west-1:123456789000:cluster/service\",\"connectivity\":\"CONNECTED\",\"connectivityAt\":\"2022-05-02T10:48:55.655Z\",\"containers\":[{\"containerArn\":\"arn:aws:ecs:eu-west-1:123456789000:container/service/12345/1234\",\"exitCode\":2,\"lastStatus\":\"STOPPED\",\"name\":\"service\",\"image\":\"123456789000.dkr.ecr.eu-west-1.amazonaws.com/image-name:latest\",\"imageDigest\":\"sha256:123\",\"runtimeId\":\"12345-123\",\"taskArn\":\"arn:aws:ecs:eu-west-1:123456789000:task/service/12345\",\"networkInterfaces\":[{\"attachmentId\":\"123\",\"privateIpv4Address\":\"127.0.0.1\"}],\"cpu\":\"0\"}],\"cpu\":\"256\",\"createdAt\":\"2022-05-02T10:48:52.289Z\",\"desiredStatus\":\"STOPPED\",\"enableExecuteCommand\":false,\"ephemeralStorage\":{\"sizeInGiB\":20},\"executionStoppedAt\":\"2022-05-03T07:29:44.856Z\",\"group\":\"service:service\",\"launchType\":\"FARGATE\",\"lastStatus\":\"STOPPED\",\"memory\":\"512\",\"overrides\":{\"containerOverrides\":[{\"name\":\"service\"}]},\"platformVersion\":\"1.4.0\",\"pullStartedAt\":\"2022-05-02T10:49:03.653Z\",\"pullStoppedAt\":\"2022-05-02T10:49:05.836Z\",\"startedAt\":\"2022-05-02T10:49:50.416Z\",\"startedBy\":\"ecs-svc/123\",\"stoppingAt\":\"2022-05-03T07:29:20.98Z\",\"stoppedAt\":\"2022-05-03T07:30:07.75Z\",\"stoppedReason\":\"Task stopped by user\",\"stopCode\":\"UserInitiated\",\"taskArn\":\"arn:aws:ecs:eu-west-1:123456789000:task/service/12345\",\"taskDefinitionArn\":\"arn:aws:ecs:eu-west-1:123456789000:task-definition/service:2\",\"updatedAt\":\"2022-05-03T07:30:07.75Z\",\"version\":8}}",
				Type:             "Notification",
				UnsubscribeURL:   "EXAMPLE",
				TopicArn:         "arn:aws:sns:eu-west-1:000000000000:cloudwatch-alarms",
				Subject:          "example subject",
			},
		},
	},
}

var activatingRunningEcsTaskEvent = events.SNSEvent{
	Records: []events.SNSEventRecord{
		{
			EventVersion:         "1.0",
			EventSubscriptionArn: "arn:aws:sns:EXAMPLE",
			EventSource:          "aws:sns",
			SNS: events.SNSEntity{
				SignatureVersion: "1",
				Timestamp:        time.Now(),
				Signature:        "EXAMPLE",
				SigningCertURL:   "EXAMPLE",
				MessageID:        "95df01b4-ee98-5cb9-9903-4c221d41eb5e",
				Message:          "{\"version\":\"0\",\"id\":\"da142eb5-0066-28ee-1bcb-87b088d66398\",\"detail-type\":\"ECS Task State Change\",\"source\":\"aws.ecs\",\"account\":\"123456789000\",\"time\":\"2022-05-03T07:30:10Z\",\"region\":\"eu-west-1\",\"resources\":[\"arn:aws:ecs:eu-west-1:123456789000:task/service/123\"],\"detail\":{\"attachments\":[{\"id\":\"114de60f-dfad-4285-980c-35216ee9f614\",\"type\":\"sdi\",\"status\":\"PRECREATED\",\"details\":[]},{\"id\":\"e91b67e3-d01a-47f9-99ae-027b95c6b1ac\",\"type\":\"eni\",\"status\":\"ATTACHED\",\"details\":[{\"name\":\"subnetId\",\"value\":\"subnet-123\"},{\"name\":\"networkInterfaceId\",\"value\":\"eni-123\"},{\"name\":\"macAddress\",\"value\":\"123\"},{\"name\":\"privateDnsName\",\"value\":\"123.eu-west-1.compute.internal\"},{\"name\":\"privateIPv4Address\",\"value\":\"123\"}]}],\"attributes\":[{\"name\":\"ecs.cpu-architecture\",\"value\":\"x86_64\"}],\"availabilityZone\":\"eu-west-1c\",\"clusterArn\":\"arn:aws:ecs:eu-west-1:123456789000:cluster/service\",\"containers\":[{\"containerArn\":\"arn:aws:ecs:eu-west-1:123456789000:container/service/123/123\",\"lastStatus\":\"RUNNING\",\"name\":\"service\",\"image\":\"123456789000.dkr.ecr.eu-west-1.amazonaws.com/image-name:latest\",\"imageDigest\":\"sha256:123\",\"runtimeId\":\"123-123\",\"taskArn\":\"arn:aws:ecs:eu-west-1:123456789000:task/service/123\",\"networkInterfaces\":[{\"attachmentId\":\"e91b67e3-d01a-47f9-99ae-027b95c6b1ac\",\"privateIpv4Address\":\"123\"}],\"cpu\":\"0\"}],\"cpu\":\"256\",\"createdAt\":\"2022-05-03T07:29:51.723Z\",\"desiredStatus\":\"RUNNING\",\"enableExecuteCommand\":false,\"ephemeralStorage\":{\"sizeInGiB\":20},\"group\":\"service:service\",\"launchType\":\"FARGATE\",\"lastStatus\":\"ACTIVATING\",\"memory\":\"512\",\"overrides\":{\"containerOverrides\":[{\"name\":\"service\"}]},\"platformVersion\":\"1.4.0\",\"pullStartedAt\":\"2022-05-03T07:30:05.962Z\",\"pullStoppedAt\":\"2022-05-03T07:30:07.915Z\",\"startedBy\":\"ecs-svc/123\",\"taskArn\":\"arn:aws:ecs:eu-west-1:123456789000:task/service/123\",\"taskDefinitionArn\":\"arn:aws:ecs:eu-west-1:123456789000:task-definition/service:2\",\"updatedAt\":\"2022-05-03T07:30:10.445Z\",\"version\":3}}",
				Type:             "Notification",
				UnsubscribeURL:   "EXAMPLE",
				TopicArn:         "arn:aws:sns:eu-west-1:000000000000:cloudwatch-alarms",
				Subject:          "example subject",
			},
		},
	},
}

var runningEcsTaskEvent = events.SNSEvent{
	Records: []events.SNSEventRecord{
		{
			EventVersion:         "1.0",
			EventSubscriptionArn: "arn:aws:sns:EXAMPLE",
			EventSource:          "aws:sns",
			SNS: events.SNSEntity{
				SignatureVersion: "1",
				Timestamp:        time.Now(),
				Signature:        "EXAMPLE",
				SigningCertURL:   "EXAMPLE",
				MessageID:        "95df01b4-ee98-5cb9-9903-4c221d41eb5e",
				Message:          "{\"version\":\"0\",\"id\":\"4ee36145-289e-878f-04fb-5eac648d2505\",\"detail-type\":\"ECS Task State Change\",\"source\":\"aws.ecs\",\"account\":\"123456789000\",\"time\":\"2022-05-03T07:30:53Z\",\"region\":\"eu-west-1\",\"resources\":[\"arn:aws:ecs:eu-west-1:123456789000:task/service/123\"],\"detail\":{\"attachments\":[{\"id\":\"114de60f-dfad-4285-980c-35216ee9f614\",\"type\":\"sdi\",\"status\":\"ATTACHED\",\"details\":[]},{\"id\":\"e91b67e3-d01a-47f9-99ae-027b95c6b1ac\",\"type\":\"eni\",\"status\":\"ATTACHED\",\"details\":[{\"name\":\"subnetId\",\"value\":\"subnet-123\"},{\"name\":\"networkInterfaceId\",\"value\":\"eni-123\"},{\"name\":\"macAddress\",\"value\":\"123\"},{\"name\":\"privateDnsName\",\"value\":\"123.eu-west-1.compute.internal\"},{\"name\":\"privateIPv4Address\",\"value\":\"123\"}]}],\"attributes\":[{\"name\":\"ecs.cpu-architecture\",\"value\":\"x86_64\"}],\"availabilityZone\":\"eu-west-1c\",\"clusterArn\":\"arn:aws:ecs:eu-west-1:123456789000:cluster/service\",\"connectivity\":\"CONNECTED\",\"connectivityAt\":\"2022-05-03T07:29:56.457Z\",\"containers\":[{\"containerArn\":\"arn:aws:ecs:eu-west-1:123456789000:container/service/123/123\",\"lastStatus\":\"RUNNING\",\"name\":\"service\",\"image\":\"123456789000.dkr.ecr.eu-west-1.amazonaws.com/image-name:latest\",\"imageDigest\":\"sha256:123\",\"runtimeId\":\"123-123\",\"taskArn\":\"arn:aws:ecs:eu-west-1:123456789000:task/service/123\",\"networkInterfaces\":[{\"attachmentId\":\"e91b67e3-d01a-47f9-99ae-027b95c6b1ac\",\"privateIpv4Address\":\"123\"}],\"cpu\":\"0\"}],\"cpu\":\"256\",\"createdAt\":\"2022-05-03T07:29:51.723Z\",\"desiredStatus\":\"RUNNING\",\"enableExecuteCommand\":false,\"ephemeralStorage\":{\"sizeInGiB\":20},\"group\":\"service:service\",\"launchType\":\"FARGATE\",\"lastStatus\":\"RUNNING\",\"memory\":\"512\",\"overrides\":{\"containerOverrides\":[{\"name\":\"service\"}]},\"platformVersion\":\"1.4.0\",\"pullStartedAt\":\"2022-05-03T07:30:05.962Z\",\"pullStoppedAt\":\"2022-05-03T07:30:07.915Z\",\"startedAt\":\"2022-05-03T07:30:53.752Z\",\"startedBy\":\"ecs-svc/123\",\"taskArn\":\"arn:aws:ecs:eu-west-1:123456789000:task/service/123\",\"taskDefinitionArn\":\"arn:aws:ecs:eu-west-1:123456789000:task-definition/service:2\",\"updatedAt\":\"2022-05-03T07:30:53.752Z\",\"version\":4}}",
				Type:             "Notification",
				UnsubscribeURL:   "EXAMPLE",
				TopicArn:         "arn:aws:sns:eu-west-1:000000000000:cloudwatch-alarms",
				Subject:          "example subject",
			},
		},
	},
}

func TestCreateSlackMessagAttachmentForAlarm(t *testing.T) {
	slackMessageAttachments := CreateSlackMessageAttachment(testAlarmEvent)
	attachemntsFields := []slackAttachmentField{
		{
			Title: "Alarm",
			Value: "sns-cloudwatch",
			Short: true,
		},
		{
			Title: "Status",
			Value: "OK",
			Short: true,
		},
		{
			Title: "Reason",
			Value: "Threshold Crossed: 1 datapoint (7.9053535353535365) was not greater than or equal to the threshold (8.0).",
			Short: false,
		},
	}

	var msg MessageAttachments

	json.Unmarshal([]byte(slackMessageAttachments), &msg)

	assert.Equal(t, "good", msg.Color)
	assert.Equal(t, "OK: sns-cloudwatch in US - N. Virginia", msg.Pretext)
	assert.Equal(t, "AWS-bot", msg.Username)
	assert.Equal(t, ":loudspeaker:", msg.Icon)
	assert.Equal(t, attachemntsFields, msg.Fields)
}

func TestCreateSlackMessageAttachmentForDeactivatingStoppedEcsTaskEvent(t *testing.T) {
	slackMessageAttachments := CreateSlackMessageAttachment(deactivatingStoppedEcsTaskEvent)
	attachemntsFields := []slackAttachmentField{
		{
			Title: "Last status",
			Value: "DEACTIVATING",
			Short: true,
		},
		{
			Title: "Desired status",
			Value: "STOPPED",
			Short: true,
		},
		{
			Title: "Cluster",
			Value: "service",
			Short: true,
		},
		{
			Title: "Task definition",
			Value: "service:2",
			Short: true,
		},
		{
			Title: "Task",
			Value: "12345",
			Short: true,
		},
		{
			Title: "Stopped reason",
			Value: "Task stopped by user",
			Short: true,
		},
	}

	var msg MessageAttachments

	json.Unmarshal([]byte(slackMessageAttachments), &msg)

	assert.Equal(t, "danger", msg.Color)
	assert.Equal(t, "Task service:2 in service cluster is changing state: DEACTIVATING -> STOPPED", msg.Pretext)
	assert.Equal(t, "AWS-bot", msg.Username)
	assert.Equal(t, ":loudspeaker:", msg.Icon)
	assert.Equal(t, attachemntsFields, msg.Fields)
}

func TestCreateSlackMessageAttachmentForProvisioningRunningEcsTaskEvent(t *testing.T) {
	slackMessageAttachments := CreateSlackMessageAttachment(provisioningRunningEcsTaskEvent)
	attachemntsFields := []slackAttachmentField{
		{
			Title: "Last status",
			Value: "PROVISIONING",
			Short: true,
		},
		{
			Title: "Desired status",
			Value: "RUNNING",
			Short: true,
		},
		{
			Title: "Cluster",
			Value: "service",
			Short: true,
		},
		{
			Title: "Task definition",
			Value: "service:2",
			Short: true,
		},
		{
			Title: "Task",
			Value: "123",
			Short: true,
		},
	}

	var msg MessageAttachments

	json.Unmarshal([]byte(slackMessageAttachments), &msg)

	assert.Equal(t, "good", msg.Color)
	assert.Equal(t, "Task service:2 in service cluster is changing state: PROVISIONING -> RUNNING", msg.Pretext)
	assert.Equal(t, "AWS-bot", msg.Username)
	assert.Equal(t, ":loudspeaker:", msg.Icon)
	assert.Equal(t, attachemntsFields, msg.Fields)
}

func TestCreateSlackMessageAttachmentForDeprovisioniningStoppedEcsTaskEvent(t *testing.T) {
	slackMessageAttachments := CreateSlackMessageAttachment(deprovisioniningStoppedEcsTaskEvent)
	attachemntsFields := []slackAttachmentField{
		{
			Title: "Last status",
			Value: "DEPROVISIONING",
			Short: true,
		},
		{
			Title: "Desired status",
			Value: "STOPPED",
			Short: true,
		},
		{
			Title: "Cluster",
			Value: "service",
			Short: true,
		},
		{
			Title: "Task definition",
			Value: "service:2",
			Short: true,
		},
		{
			Title: "Task",
			Value: "12345",
			Short: true,
		},
		{
			Title: "Stopped reason",
			Value: "Task stopped by user",
			Short: true,
		},
	}

	var msg MessageAttachments

	json.Unmarshal([]byte(slackMessageAttachments), &msg)

	assert.Equal(t, "danger", msg.Color)
	assert.Equal(t, "Task service:2 in service cluster is changing state: DEPROVISIONING -> STOPPED", msg.Pretext)
	assert.Equal(t, "AWS-bot", msg.Username)
	assert.Equal(t, ":loudspeaker:", msg.Icon)
	assert.Equal(t, attachemntsFields, msg.Fields)
}

func TestCreateSlackMessageAttachmentForPendingRunningEcsTaskEvent(t *testing.T) {
	slackMessageAttachments := CreateSlackMessageAttachment(pendingRunningEcsTaskEvent)
	attachemntsFields := []slackAttachmentField{
		{
			Title: "Last status",
			Value: "PENDING",
			Short: true,
		},
		{
			Title: "Desired status",
			Value: "RUNNING",
			Short: true,
		},
		{
			Title: "Cluster",
			Value: "service",
			Short: true,
		},
		{
			Title: "Task definition",
			Value: "service:2",
			Short: true,
		},
		{
			Title: "Task",
			Value: "123",
			Short: true,
		},
	}

	var msg MessageAttachments

	json.Unmarshal([]byte(slackMessageAttachments), &msg)

	assert.Equal(t, "good", msg.Color)
	assert.Equal(t, "Task service:2 in service cluster is changing state: PENDING -> RUNNING", msg.Pretext)
	assert.Equal(t, "AWS-bot", msg.Username)
	assert.Equal(t, ":loudspeaker:", msg.Icon)
	assert.Equal(t, attachemntsFields, msg.Fields)
}

func TestCreateSlackMessageAttachmentForStoppedEcsTaskEvent(t *testing.T) {
	slackMessageAttachments := CreateSlackMessageAttachment(stoppedEcsTaskEvent)
	attachemntsFields := []slackAttachmentField{
		{
			Title: "Last status",
			Value: "STOPPED",
			Short: true,
		},
		{
			Title: "Desired status",
			Value: "STOPPED",
			Short: true,
		},
		{
			Title: "Cluster",
			Value: "service",
			Short: true,
		},
		{
			Title: "Task definition",
			Value: "service:2",
			Short: true,
		},
		{
			Title: "Task",
			Value: "12345",
			Short: true,
		},
		{
			Title: "Stopped reason",
			Value: "Task stopped by user",
			Short: true,
		},
	}

	var msg MessageAttachments

	json.Unmarshal([]byte(slackMessageAttachments), &msg)

	assert.Equal(t, "danger", msg.Color)
	assert.Equal(t, "Task service:2 in service cluster changed state: STOPPED", msg.Pretext)
	assert.Equal(t, "AWS-bot", msg.Username)
	assert.Equal(t, ":loudspeaker:", msg.Icon)
	assert.Equal(t, attachemntsFields, msg.Fields)
}

func TestCreateSlackMessageAttachmentForActivatingRunningEcsTaskEvent(t *testing.T) {
	slackMessageAttachments := CreateSlackMessageAttachment(activatingRunningEcsTaskEvent)
	attachemntsFields := []slackAttachmentField{
		{
			Title: "Last status",
			Value: "ACTIVATING",
			Short: true,
		},
		{
			Title: "Desired status",
			Value: "RUNNING",
			Short: true,
		},
		{
			Title: "Cluster",
			Value: "service",
			Short: true,
		},
		{
			Title: "Task definition",
			Value: "service:2",
			Short: true,
		},
		{
			Title: "Task",
			Value: "123",
			Short: true,
		},
	}

	var msg MessageAttachments

	json.Unmarshal([]byte(slackMessageAttachments), &msg)

	assert.Equal(t, "good", msg.Color)
	assert.Equal(t, "Task service:2 in service cluster is changing state: ACTIVATING -> RUNNING", msg.Pretext)
	assert.Equal(t, "AWS-bot", msg.Username)
	assert.Equal(t, ":loudspeaker:", msg.Icon)
	assert.Equal(t, attachemntsFields, msg.Fields)
}

func TestCreateSlackMessageAttachmentForRunningEcsTaskEvent(t *testing.T) {
	slackMessageAttachments := CreateSlackMessageAttachment(runningEcsTaskEvent)
	attachemntsFields := []slackAttachmentField{
		{
			Title: "Last status",
			Value: "RUNNING",
			Short: true,
		},
		{
			Title: "Desired status",
			Value: "RUNNING",
			Short: true,
		},
		{
			Title: "Cluster",
			Value: "service",
			Short: true,
		},
		{
			Title: "Task definition",
			Value: "service:2",
			Short: true,
		},
		{
			Title: "Task",
			Value: "123",
			Short: true,
		},
	}

	var msg MessageAttachments

	json.Unmarshal([]byte(slackMessageAttachments), &msg)

	assert.Equal(t, "good", msg.Color)
	assert.Equal(t, "Task service:2 in service cluster changed state: RUNNING", msg.Pretext)
	assert.Equal(t, "AWS-bot", msg.Username)
	assert.Equal(t, ":loudspeaker:", msg.Icon)
	assert.Equal(t, attachemntsFields, msg.Fields)
}
