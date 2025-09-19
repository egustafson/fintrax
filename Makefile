# Makefile
# -------------------------------------------------------------------

.PHONY: all
all: build test

GIT_SUMMARY := $(shell git describe --tags --dirty --always)
BUILD_DATE  := $(shell date -u "+%Y-%m-%dT%H:%M:%SZ")

DIST_DIR = dist

# -------------------------------------------------------------------
GO_FLAGS =

.PHONY: build
build: fintrax

${DIST_DIR}:
	mkdir -p ${DIST_DIR}

.PHONY: fintrax  # force a rebuld always
fintrax: ${DIST_DIR}
	go build -ldflags "-X main.GitSummary=$(GIT_SUMMARY) -X main.BuildDate=$(BUILD_DATE)" -o ${DIST_DIR}/$@

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:

.PHONY: clean
clean:
	go clean ./...
	rm -rf ${DIST_DIR}

.PHONY: real_clean
real_clean: clean
	go clean -cache
