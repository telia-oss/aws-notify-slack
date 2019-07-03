TARGET ?= linux
ARCH ?= amd64

build:
	dep ensure -v
	env GOOS=$(TARGET) GOARCH=$(ARCH) go build -ldflags="-s -w" -o bin/main main.go

zip:
	zip -j bin/main.zip bin/main

test:
	go test ./... --cover

run:
	sam local invoke NotifySlackFunction -e snsEvent.json

deploy:
	aws s3 cp bin/*.zip s3://telia-oss/aws-notify-slack/main.zip --acl "public-read"

.PHONY: clean
clean:
	rm -rf ./bin ./vendor Gopkg.lock
