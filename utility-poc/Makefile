all: cli

cli:
	@mkdir -p bin/
	@go get ./...
	@go install profile_tool

clean:
	@rm -rf bin/ pkg/ src/github.com/

.PHONY: all clean cli
