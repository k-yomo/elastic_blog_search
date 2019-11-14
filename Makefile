.PHONY: build
build:
	CGO_ENABLED=0 go build -o bin/server \
        -ldflags "-w -s" ./src/cmd/server

.PHONY: setup
setup:
	export GO111MODULE=off
	go get gopkg.in/urfave/cli.v2@master
	go get github.com/oxequa/realize
