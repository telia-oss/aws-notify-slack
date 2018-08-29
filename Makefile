TARGET ?= linux
ARCH ?= amd64

build:
	dep ensure -v
	env GOOS=$(TARGET) GOARCH=$(ARCH) go build -ldflags="-s -w" -o bin/main main.go

zip:
	zip -r bin/main.zip ./*

test:
	go test ./... --cover

run:
	sam local invoke NotifySlackFunction -e snsEvent.json

.PHONY: clean
clean:
	rm -rf ./bin ./vendor Gopkg.lock
