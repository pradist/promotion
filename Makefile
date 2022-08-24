test:
	go test ./... -v -coverprofile=coverage.out

lint:
	golangci-lint run -v --color always --timeout 5m

mock:
	go generate ./...