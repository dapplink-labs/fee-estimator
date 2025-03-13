SHELL := /bin/bash

fee-estimator:
	env GO111MODULE=on go build -v $(LDFLAGS) ./cmd/fee-estimator
.PHONY: fee-estimator

go-proto:
	./bin/go_compile.sh

clean:
	rm fee-estimator

test:
	go test -v ./...

lint:
	golangci-lint run ./...