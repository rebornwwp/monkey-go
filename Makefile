
.PHONY: test

test:
	@go get ./...

test_coverage:
	@go test ./... -coverprofile=coverage.out
