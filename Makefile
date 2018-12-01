NAME = "svf"
REVISION = $(shell git describe --always --dirty --tags)
LDFLAGS = -ldflags "-X github.com/blp1526/go-shell-variables-file/pkg/cmd.revision=$(REVISION)"

.PHONY: all
all: build

.PHONY: clean
clean:
	rm -rf bin/
	mkdir -p bin/
	@echo

.PHONY: dep
dep:
	go get github.com/golang/dep/cmd/dep
	dep ensure
	@echo

.PHONY: gometalinter
gometalinter:
	go get github.com/alecthomas/gometalinter
	gometalinter --install
	@echo
	gometalinter --config .gometalinter.json ./...
	@echo

.PHONY: test
test: gometalinter
	go test ./... -v --cover -race -covermode=atomic -coverprofile=coverage.txt
	@echo

.PHONY: build
build: dep test clean
	go build $(LDFLAGS) -o bin/$(NAME)
	@echo
