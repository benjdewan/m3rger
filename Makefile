# Constants
VERSION := $(shell git describe --tags || echo DEV-BUILD)
LDFLAGS := "-w -s -X main.version=$(VERSION)"
export CGO_ENABLED := 0
export GOARCH := amd64

# Targets/Source
TARGETS := m3rger-linux m3rger-windows.exe m3rger-darwin
SOURCE := $(shell find . -type f -iname '*.go')

# Executables
DEP := $(GOPATH)/bin/dep
GOMETALINTER := $(GOPATH)/bin/gometalinter

# Sanity Check
ifndef GOPATH
    $(error GOPATH must be defined to build this project)
endif

default: setup lint test $(TARGETS)
.PHONY: all

all: $(TARGETS)
.PHONY: all

setup:
	go get -u github.com/alecthomas/gometalinter github.com/golang/dep/cmd/dep
	$(GOMETALINTER) --install
	$(DEP) ensure
.PHONY: setup

m3rger-%.exe: $(SOURCE)
	GOOS=$* go build -ldflags $(LDFLAGS) -o "$@"

m3rger-%: $(SOURCE)
	GOOS=$* go build -ldflags $(LDFLAGS) -o "$@"

lint:
	$(GOMETALINTER) --disable=gotype --deadline=90s .
.PHONY: lint

test:
	go test -v .
.PHONY: test

clean:
	rm -rf $(TARGETS)
.PHONY: clean
