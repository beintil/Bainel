fmt:
	gofmt -w .

lint: fmt
	golangci-lint run --fix

test:
	go test ./...

