NAME=uuid4


install: get-deps test build


get-deps:
	go get ./...

build:
	go build -v

build-linux:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v

test:
	go test -v -coverprofile=cover.out ./...

bench:
	printf "BENCHMARK REPORT - $(shell date)\n" > benchmark_results.txt
	go test -v -bench=. -run=^$$ >> benchmark_results.txt

cover:
	go tool cover -html=cover.out
