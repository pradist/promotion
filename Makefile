test:
	go test ./... -v -coverprofile=coverage.out

test-coverage:
	go test ./... -v -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

html-coverage:
	go tool cover -html=coverage.out

lint:
	golangci-lint run -v --color always --timeout 5m

rm-mock:
	find . -type d -name "mock_*" -exec rm -rf {} +

gen-mock:
	go generate ./...

run:
	go run ./cmd/promotion
