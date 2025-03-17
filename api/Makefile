GO_FILES = $(wildcard *.go)

API_BIN = ../lib/bin/api

$(API_BIN): $(GO_FILES)
	go build -o $(API_BIN)

all: $(API_BIN)
