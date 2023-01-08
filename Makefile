test:
	go test ./... -v -coverprofile=coverage.out

lint:
	golangci-lint run -v --color always --timeout 5m

rm-mock:
	find . -type d -name "mock_*" -exec rm -rf {} +

gen-mock:
	go generate ./...