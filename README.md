# AWS-Notify-Slack
[![Build Status](https://travis-ci.org/telia-oss/aws-notify-slack.svg?branch=master)](https://travis-ci.org/telia-oss/aws-notify-slack)![](https://img.shields.io/maintenance/yes/2018.svg)

A lambda function that format and forward a AWS cloudwatch event to slack

![alt text](https://github.com/telia-oss/aws-notify-slack/blob/master/media/warn.png)
![alt text](https://github.com/telia-oss/aws-notify-slack/blob/master/media/ok.png)

## Supported event types
- [x] CloudWatch
- [ ] Autoscaling

## Run unit tests 
$ make test

## Build binary 
$ make build

## Run function using sam local with snsEvent payload  
$ make run

## Environment Variables

|      Name     |     Required  |     Type      |   Description  |
| ------------- | ------------- | ------------- | -------------- |
| SLACK_HOOK    | Yes           | String        | Slack hook url |
| USERNAME      | No            | String        | Slack username |
| ICON          | No            | String        | Slack icon     |

## Terraform
The Terraform module for this lambda can be found [here](https://github.com/telia-oss/terraform-aws-lambda-slack)

## Authors

Currently maintained by [these contributors](../../graphs/contributors).

## License

MIT License. See [LICENSE](LICENSE) for full details.
