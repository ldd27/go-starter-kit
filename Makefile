PACKAGE=github.com/ldd27/go-starter-kit
VERSION_PACKAGE=github.com/ldd27/go-starter-kit
PREFIX=$(shell pwd)
OUTPUT_DIR=${PREFIX}/bin
OUTPUT_FILE=${OUTPUT_DIR}/go-starter-kit

VERSION=${CI_COMMIT_REF_NAME}
GIT_REVISION=$(shell git rev-parse --short HEAD)
GIT_BRANCH=$(shell git name-rev --name-only HEAD)
BUILD_TIME=$(shell date '+%Y-%m-%d %H:%M:%S')

BUILD_ARGS := \
    -ldflags "-X ${VERSION_PACKAGE}/pkg/version.version=$(VERSION) \
   	-X ${VERSION_PACKAGE}/pkg/version.gitRevision=$(GIT_REVISION) \
   	-X ${VERSION_PACKAGE}/pkg/version.gitBranch=$(GIT_BRANCH) \
  	-X '${VERSION_PACKAGE}/pkg/version.buildTime=$(BUILD_TIME)'" # 添加单引号解决日期带空格问题
EXTRA_BUILD_ARGS=

.PHONY: build
build:
	@go build $(BUILD_ARGS) $(EXTRA_BUILD_ARGS) -o ${OUTPUT_FILE} ./cmd/main.go

.PHONY: docs
docs:
	@#swag init -g router.go -d ./internal/router -ot yaml
	@swag init -g ./internal/router/router.go -ot yaml

.PHONY: wire
wire:
	@wire internal/wire/wire.go