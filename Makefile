export MY_WORKSPACE?=$(PWD)
export MY_USER?=root

OUTPUT_BASE?=./build

all: deps test build

build: build-server

clean: clean-output-dir

run:
	@SSH_PRIVATE_KEY=$(pk) docker-compose build app
	docker-compose up app

deps:
	go mod download

deps-fix:
	go mod tidy

check-fmt:
	test -z $(shell gofmt -l ./)

test: deps
	go vet ./...
	go test ./...

test-race:
	go test -race ./...

test-coverage: output-dir
	go test ./... -coverpkg=./... -coverprofile=$(COVERAGE_OUTPUT)

build-server: output-dir
	go build -o $(OUTPUT_BASE)/server drafteame/cmd/

output-dir:
	mkdir -p $(OUTPUT_BASE)

clean-server:
	rm $(OUTPUT_BASE)server

clean-output-dir:
	rm -rf $(OUTPUT_BASE)
