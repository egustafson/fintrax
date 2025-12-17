# Makefile
# -------------------------------------------------------------------

.PHONY: all
all: build test

GIT_SUMMARY := $(shell git describe --tags --dirty --always)
BUILD_VER   := $(shell git describe --tags --always)
BUILD_DATE  := $(shell date -u "+%Y-%m-%dT%H:%M:%SZ")

DIST_DIR = dist

DOCKER = podman

# -------------------------------------------------------------------
GO = go
GO_FLAGS =

.PHONY: build
build: fintrax

.PHONY: preflight
preflight:
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest

${DIST_DIR}:
	mkdir -p ${DIST_DIR}

.PHONY: fintrax  # force a rebuld always
fintrax: ${DIST_DIR}
	go build -ldflags "-X main.GitSummary=$(GIT_SUMMARY) -X main.BuildDate=$(BUILD_DATE)" -o ${DIST_DIR}/$@

.PHONY: test
test: test-lint
test: unit-test

.PHONY: test-lint
test-lint:
	golangci-lint run ./...

.PHONY: unit-test
unit-test:
	go test ./...

.PHONY: package
package:
	${DOCKER} build -t fintrax:${BUILD_VER} .

.PHONY: clean
clean:
	go clean ./...
	rm -rf ${DIST_DIR}

.PHONY: real_clean
real_clean: clean
	go clean -cache
