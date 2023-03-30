BIN := bin

PHONY: build install test

$(BIN):
	mkdir -p $@

build: $(BIN)
	go build -o $(BIN)/txrollbackcheck .

install:
	go install

test: build
	go test ./...
	# Due to an issue with importing in a anaylsistest's test data some hoop jumping is required
	# I call twice to avoid collecting package downloads in output
	-go vet -vettool=$(BIN)/txrollbackcheck ./testdata/sqlx_examples
	-go vet -vettool=$(BIN)/txrollbackcheck ./testdata/sqlx_examples 2> sqlx_examples_results.txt
	diff -a sqlx_examples_results.txt ./testdata/sqlx_examples/expected_results.txt

lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.50.0
	./bin/golangci-lint run
