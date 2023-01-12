cli:
	@go build -o rlscli cmd/rlscli/*.go

.PHONY: install
install: cli
	@cp rlscli ${GOBIN}/rlscli

.PHONY: test
test:
	@go test -v ./...

.PHONY: lint
lint:
	@golint -set_exit_status