.PHONY: test
test:
	@mkdir -p tmp
	@go test -cover -coverpkg=./... -coverprofile=tmp/cover.out ./...
