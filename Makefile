DESTINATION=_dst/meister
default: build

build:
	GO111MODULE=on go build -o $(DESTINATION) ./cmd/cli

