TARGET ?= linux
ARCH ?= amd64

build:
	env GOOS=$(TARGET) GOARCH=$(ARCH) go build -ldflags="-s -w" -o bin/main main.go

zip:
	zip -j bin/main.zip bin/main

test:
	go test ./... --cover

runAlarm:
	sam local invoke NotifySlackFunction -e alarmSnsEvent.json

runStateChange:
	sam local invoke NotifySlackFunction -e stateChangeSnsEvent.json

deploy:
	aws s3 cp bin/*.zip s3://telia-oss/aws-notify-slack/main.zip --acl "public-read"

.PHONY: clean
clean:
	rm -rf ./bin ./vendor Gopkg.lock
