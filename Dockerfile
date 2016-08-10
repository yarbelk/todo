FROM golang:1.7-wheezy

RUN mkdir -p "$GOPATH/src/github.com/yarbelk/todo/"
COPY . "$GOPATH/src/github.com/yarbelk/todo/"
WORKDIR "$GOPATH/src/github.com/yarbelk/todo/"

RUN make
