
.PHONY: test

test:
	@go get ./...

test-coverage:
	@go test ./... -coverprofile=coverage.out
