DOCKER_REGISTRY   ?= us.gcr.io
IMAGE_PREFIX      ?= azuqua-alpha
TARGETS           ?= darwin/amd64 linux/amd64 linux/386 linux/arm linux/arm64 linux/ppc64le linux/s390x windows/amd64
DIST_DIRS         = find * -type d -exec
APP               = mc-helper

# go option
GO        ?= go
TAGS      :=
TESTS     := .
TESTFLAGS :=
LDFLAGS   := -w -s
GOFLAGS   :=
BINDIR    := $(CURDIR)/bin
BINARIES  := mchelperd mchelper
SHORT_NAME := mc-helper

# Required for globs to work correctly
SHELL=/bin/bash


.PHONY: all
all: build

get-cluster-credentials:
	gcloud container clusters get-credentials "$(CLUSTER)" --project="$(PROJECT)" --zone="$(ZONE)"

.PHONY: dep
dep:
	@dep ensure
	@echo "cleaning up sub-deps ..."
	find vendor/ -path '*/vendor' -type d | xargs -IX rm -r X 2>/dev/null || true
	@echo "dep finished"

.PHONY: lint
lint:
	@echo "  running linter ..."
	@revive -config default.toml -formatter stylish

.PHONY: build
build: lint
	@echo "  building releases in ./bin/..."
	GOBIN=$(BINDIR) $(GO) install $(GOFLAGS) -tags '$(TAGS)' -ldflags '$(LDFLAGS)' github.com/jaredallard/k8s-mc-helper/cmd/...

# usage: make clean build-cross dist APP=mc-helper|mc-helper VERSION=v2.0.0-alpha.3
.PHONY: build-cross
build-cross: LDFLAGS += -extldflags "-static"
build-cross:
	CGO_ENABLED=1 gox -parallel=3 -output="_dist/{{.OS}}-{{.Arch}}/{{.Dir}}" -osarch='$(TARGETS)' $(GOFLAGS) $(if $(TAGS),-tags '$(TAGS)',) -ldflags '$(LDFLAGS)' github.com/jaredallard/k8s-mc-helper/cmd/$(APP)


.PHONY: dist
dist: build
	mkdir -p _release _dist || true
	( \
		rm -rf _dist/* && \
		cd _dist && \
		cp ../LICENSE . && \
		cp ../README.md . && \
		cp ../bin/* . && \
		7z a mc-helper-$(VERSION)-amd64.tar * && \
		xz mc-helper-$(VERSION)-amd64.tar && \
		mv mc-helper-$(VERSION)-amd64.tar.xz ../_release/ \
	)

.PHONY: check-docker
check-docker:
	@if [ -z $$(which docker) ]; then \
	  echo "Missing \`docker\` client which is required for development"; \
	  exit 2; \
	fi

.PHONY: docker-build
docker-build: check-docker 
	docker build --build-arg "TAGS=$(TAGS)" --build-arg "LDFLAGS=$(LDFLAGS)" --rm -t ${IMAGE} . -f Dockerfile
	docker tag ${IMAGE} ${MUTABLE_IMAGE}

.PHONY: docker-push
docker-push:
	docker push ${MUTABLE_IMAGE}

.PHONY: test
test: build
test: TESTFLAGS += -race -v
test: test-unit

.PHONY: test-unit
test-unit: get-cluster-credentials
	@echo
	echo "==> Running unit tests <=="
	mc-helper_HOME=/no/such/dir $(GO) test $(GOFLAGS) -run $(TESTS) $(PKG) $(TESTFLAGS)

.PHONY: test-integration
test-integration: get-cluster-credentials
	@echo
	echo "==> Running integration tests <=="
	mc-helper_HOME=/no/such/dir $(GO) test $(GOFLAGS) -run $(TESTS) $(PKG) $(TESTFLAGS) -integration

.PHONY: test-unit-coverage
test-unit-coverage: get-cluster-credentials
	@echo
	@echo "==> Running unit tests with coverage profiling <=="
	mc-helper_HOME=/no/such/dir $(GO) test $(GOFLAGS) -run $(TESTS) $(PKG) $(TESTFLAGS) -coverprofile coverage.out
	go tool cover -func coverage.out


.PHONY: protoc
protoc:
	$(MAKE) -C _proto/ all

.PHONY: clean
clean:
	@rm -rf $(BINDIR) ./rootfs/mc-helper ./_dist

.PHONY: coverage
coverage:
	@scripts/coverage.sh

HAS_GLIDE := $(shell command -v glide;)
HAS_GOX := $(shell command -v gox;)
HAS_GIT := $(shell command -v git;)

.PHONY: bootstrap
bootstrap:
ifndef HAS_GLIDE
	go get -u github.com/Masterminds/glide
endif
ifndef HAS_GOX
	go get -u github.com/mitchellh/gox
endif

ifndef HAS_GIT
	$(error You must install Git)
endif
	dep ensure

include versioning.mk
