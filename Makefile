.PHONY: build
build:
	CGO_ENABLED=0 go build -o bin/server \
        -ldflags "-w -s" ./src/cmd/server

.PHONY: build_indexer
build_indexer:
	CGO_ENABLED=0 go build -o bin/indexer \
        -ldflags "-w -s" ./src/cmd/indexer

.PHONY: setup
setup:
	export GO111MODULE=off
	go get gopkg.in/urfave/cli.v2@master
	go get github.com/oxequa/realize

.PHONY: run_dev
run_dev:
	realize start

.PHONY: gen_goa
gen_goa:
	goa gen github.com/k-yomo/elastic_blog_search/src/design -o src/
