export PATH := $(abspath ./vendor/bin):$(PATH)

BASE_PACKAGE_NAME  = github.com/omegion/go-db-backup
GIT_VERSION = $(shell git describe --tags --always 2> /dev/null || echo 0.0.0)
LDFLAGS            = -ldflags "-X $(BASE_PACKAGE_NAME)/pkg/info.Version=$(GIT_VERSION)"
BUFFER            := $(shell mktemp)

.PHONY: build
build:
	CGO_ENABLED=0 go build $(LDFLAGS) -installsuffix cgo -o dist/db-backup cmd/db-backup/main.go

build-for-container:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -a -installsuffix cgo -o dist/db-backup-linux cmd/db-backup/main.go


.PHONY: lint
lint:
	@echo "Checking code style"
	gofmt -l . | tee $(BUFFER)
	@! test -s $(BUFFER)
	go vet ./...
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.31.0
	@golangci-lint --version
	golangci-lint run

.PHONY: cut-tag
cut-tag:
	@echo "Cutting $(version)"
	git tag $(version)
	git push origin $(version)

.PHONY: release
release:
	@echo "Releasing $(GIT_VERSION)"
	docker build -t db-backup . --build-arg VERSION=$(GIT_VERSION)
	docker tag db-backup:latest omegion/go-db-backup:$(GIT_VERSION)
	docker push omegion/go-db-backup:$(GIT_VERSION)
