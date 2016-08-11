GIT_VERSION := $(shell git describe --abbrev=4 --dirty --always --tags)

all: clean store-srv command-srv query-srv
.PHONY: all

store-srv: store/cli/store-srv/main.go
	go build -ldflags "-X main.version=$(GIT_VERSION)" -o $@ $<

command-srv: command/cli/command-srv/main.go
	go build -ldflags "-X main.version=$(GIT_VERSION)" -o $@ $<

query-srv: query/cli/query-srv/main.go
	go build -ldflags "-X main.version=$(GIT_VERSION)" -o $@ $<
.PHONY: clean
	
clean:
	rm command-srv store-srv query-srv || true
