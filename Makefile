test:
	@go test -v ./tests/...

coverage:
	@go test -coverprofile=coverage.out ./tools/... ./tests/...
	@go tool cover -html=coverage.out -o ./tests/coverage.html
	@rm coverage.out
