export
APP_VERSION := $(shell git describe --tags)
LD_FLAGS="-X 'rufin/config.Version=$(APP_VERSION)' -X 'rufin/config.BuildDate=$(shell date)'"
CGO_ARGS=CGO_ENABLED=0


all: build

.PHONY: build

bin/rufin-linux-amd64:
	mkdir -p bin
	cd src && $(CGO_ARGS) GOOS=linux GOARCH=amd64 go build \
	-ldflags=$(LD_FLAGS) \
	-o ../bin/rufin-linux-$(APP_VERSION)-amd64 \
	*.go

bin/rufin-linux-arm64:
	mkdir -p bin
	cd src && $(CGO_ARGS) GOOS=linux GOARCH=arm64 go build \
	-ldflags=$(LD_FLAGS) \
	-o ../bin/rufin-linux-$(APP_VERSION)-arm64 \
	*.go

bin/rufin-darwin-arm64:
	mkdir -p bin
	cd src && $(CGO_ARGS) GOOS=darwin GOARCH=arm64 go build \
	-ldflags=$(LD_FLAGS) \
	-o ../bin/rufin-darwin-$(APP_VERSION)-arm64 \
	*.go

build: clean bin/rufin-linux-amd64 bin/rufin-linux-arm64 bin/rufin-darwin-arm64
	@echo build version $(APP_VERSION)

#test:
#	CC_loglevel="error" go test -v ./... -cover

lint:
	docker run --rm -v $$(pwd):/app -w /app golangci/golangci-lint:v1.62.2 golangci-lint run -v
	

.PHONY: clean
clean:
	@rm -Rf bin/rufin*

.PHONY: release
release:
	@echo "Creating release for version $(APP_VERSION)"
	gh release create $(APP_VERSION) --generate-notes ./bin/rufin-*