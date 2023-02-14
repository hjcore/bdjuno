VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT  := $(shell git log -1 --format='%H')

export GO111MODULE = on

###############################################################################
###                                   All                                   ###
###############################################################################

all: lint build test-unit

###############################################################################
###                                Build flags                              ###
###############################################################################

LD_FLAGS = -X github.com/forbole/juno/v4/cmd.Version=$(VERSION) \
	-X github.com/forbole/juno/v4/cmd.Commit=$(COMMIT)
BUILD_FLAGS :=  -ldflags '$(LD_FLAGS)'

ifeq ($(LINK_STATICALLY),true)
  LD_FLAGS += -linkmode=external -extldflags "-Wl,-z,muldefs -static"
endif

build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

BUILD_FLAGS :=  -ldflags '$(LD_FLAGS)' -tags "$(build_tags)"

###############################################################################
###                                  Build                                  ###
###############################################################################

build: go.sum
ifeq ($(OS),Windows_NT)
	@echo "building gjuno binary..."
	@go build -mod=readonly $(BUILD_FLAGS) -o build/gjuno.exe ./cmd/gjuno
else
	@echo "building gjuno binary..."
	@go build -mod=readonly $(BUILD_FLAGS) -o build/gjuno ./cmd/gjuno
endif
.PHONY: build

###############################################################################
###                                 Install                                 ###
###############################################################################

install: go.sum
	@echo "installing gjuno binary..."
	@go install -mod=readonly $(BUILD_FLAGS) ./cmd/gjuno
.PHONY: install

###############################################################################
###                           Tests & Simulation                            ###
###############################################################################

stop-docker-test:
	@echo "Stopping Docker container..."
	@docker stop gjuno-test-db || true && docker rm gjuno-test-db || true
.PHONY: stop-docker-test

start-docker-test: stop-docker-test
	@echo "Starting Docker container..."
	@docker run --name gjuno-test-db -e POSTGRES_USER=gjuno -e POSTGRES_PASSWORD=password -e POSTGRES_DB=gjuno -d -p 6433:5432 postgres
.PHONY: start-docker-test

test-unit: start-docker-test
	@echo "Executing unit tests..."
	@go test -mod=readonly -v -coverprofile coverage.txt ./...
.PHONY: test-unit

lint:
	golangci-lint run --out-format=tab

lint-fix:
	golangci-lint run --fix --out-format=tab --issues-exit-code=0
.PHONY: lint lint-fix

format:
	find . -name '*.go' -type f -not -path "*.git*" | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "*.git*" | xargs misspell -w
	find . -name '*.go' -type f -not -path "*.git*" | xargs goimports -w -local github.com/gotabit/gjuno
.PHONY: format

clean:
	rm -f tools-stamp ./build/**
.PHONY: clean
