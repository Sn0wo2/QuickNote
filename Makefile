VERSION := $(shell git describe --tags --always --dirty)
COMMIT  := $(shell git rev-parse HEAD)
DATE    := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)

EXT :=
MOVE := mv
RM_RF := rm -rf ./static

UNAME_S := $(shell uname -s)

RACE_PLATFORMS := \
	linux/amd64 linux/arm64 linux/ppc64le linux/s390x \
	freebsd/amd64 netbsd/amd64 darwin/amd64 darwin/arm64 windows/amd64
GO_TEST_PLATFORM := $(shell go env GOOS)/$(shell go env GOARCH)
GO_TEST_RACE := $(if $(and \
	$(filter 1,$(shell go env CGO_ENABLED)), \
	$(filter $(GO_TEST_PLATFORM),$(RACE_PLATFORMS))),-race)

ifeq ($(OS),Windows_NT)
	EXT := .exe
	MOVE := move
	RM_RF := if exist static ( rmdir /s /q static )
endif

.PHONY: test_release run_docker update_depends update_frontend_depends build lint lint-fix test security ci

test_release:
	@echo "<<< Starting test release"
	@goreleaser release --snapshot --clean
	@echo "<<< Test release complete"

run_docker:
	@echo "<<< Starting Docker containers"
	@docker-compose up -d
	@echo "<<< Docker containers are up"

update_depends:
	@echo "<<< Updating Go dependencies"
	@go get all
	@go mod tidy
	@echo "<<< Go dependencies updated"

update_frontend_depends:
	@echo "<<< Updating frontend dependencies"
	@cd Frontend && bun update
	@echo "<<< Frontend dependencies updated"

build:
	@echo "<<< Building React Frontend"
	@(cd Frontend && bun install && bun run build)

	@echo "<<< Cleaning static folder"
	@$(RM_RF)

	@echo "<<< Moving new static files"
	@$(MOVE) Frontend/static .

	@echo "<<< Building Go Backend"
	@go build \
		-mod=readonly \
		-trimpath \
		-tags="mysql postgres sqlite sqlserver" \
		-o=QuicKNote$(EXT) \
		-ldflags="-s -w -buildid= -extldflags=-static \
		-X github.com/Sn0wo2/QuickNote/pkg/version.version=$(VERSION) \
		-X github.com/Sn0wo2/QuickNote/pkg/version.commit=$(COMMIT) \
		-X github.com/Sn0wo2/QuickNote/pkg/version.date=$(DATE)" \
		-gcflags="all=-d=ssa/check_bce/debug=0" \
		-asmflags="-trimpath" main.go

	@echo "<<< Build complete: ./QuicKNote$(EXT)"

lint:
	go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.15.0 run

lint-fix:
	go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.15.0 run --fix

test:
	go test $(GO_TEST_RACE) ./...

security:
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

ci: lint test security
