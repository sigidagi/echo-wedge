.PHONY: build clean serve 

VERSION=0.2.1
all: linux 

linux:
	go build $(GO_EXTRA_BUILD_ARGS) -ldflags "-s -w -X main.version=$(VERSION)" -o build/echo-wedge backend/cmd/main.go

clean:
	@echo "Cleaning up"
	@rm -rf build 

serve: build
	@echo "Starting Echo Rest service"
	./build/echo-wedge

