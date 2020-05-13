GO ?= go
GOBUILD ?= $(GO) build
BINARY_NAME ?= fx2preset
GOFILES := $(shell find . -name "*.go")
GOPATH_BIN = ~/go/bin/

all: build

.PHONY: build
build:
		$(GOBUILD) -o $(BINARY_NAME) -v $(GOFILES)

.PHONY: clean
clean:
		rm $(BINARY_NAME)

.PHONY: install
install: build
		cp -f $(BINARY_NAME) $(GOPATH_BIN)

.PHONY: uninstall
uninstall:
		rm $(GOPATH_BIN)$(BINARY_NAME)