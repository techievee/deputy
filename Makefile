deputy_CMD_PATH = "./cmd"
BUILD_NUMBER ?= v1.0.0
GO111MODULE = on
export GO111MODULE

ALL_LINUX = linux-amd64 \
	linux-386 \
	linux-ppc64le \
	linux-arm-5 \
	linux-arm-6 \
	linux-arm-7 \
	linux-arm64 \
	linux-mips \
	linux-mipsle \
	linux-mips64 \
	linux-mips64le

ALL = $(ALL_LINUX) \
	darwin-amd64 \
	windows-amd64

all: $(ALL:%=build/%/deputy)

release: $(ALL:%=build/deputy-%.tar.gz)

release-linux: $(ALL_LINUX:%=build/deputy-%.tar.gz)

bin-windows: build/windows-amd64/deputy.exe

bin-darwin: build/darwin-amd64/deputy

bin:
	go build -trimpath -ldflags "-X main.Build=$(BUILD_NUMBER)" -o ./deputy ${deputy_CMD_PATH}

install:
	go install -trimpath -ldflags "-X main.Build=$(BUILD_NUMBER)" ${deputy_CMD_PATH}

build/%/deputy: .FORCE
		GOOS=$(firstword $(subst -, , $*)) \
		GOARCH=$(word 2, $(subst -, ,$*)) \
		GOARM=$(word 3, $(subst -, ,$*)) \
		go build -trimpath -o $@ -ldflags "-X main.Build=$(BUILD_NUMBER)" ${deputy_CMD_PATH}
		cp -R test_samples/* build/$*


build/%/deputy.exe: build/%/deputy
	mv $< $@

build/deputy-%.tar.gz: build/%/deputy
	cp -R test_samples/* build/$* && tar -zcv -C build/$* -f $@ deputy

build/deputy-%.zip: build/%/deputy.exe
	cd build/$* && copy test_samples/*.json ./$* && zip ../deputy-$*.zip deputy.exe


vet:
	go vet -v ./...

test:
	go test -v ./...

test-cov-html:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out
	go tool cover -func=coverage.out

code-coverage:
	go get -t -v ./...
    go test -race -coverprofile=coverage.txt -covermode=atomic

bench:
	go test -bench . ./...

service:
	@echo > /dev/null
	$(eval deputy_CMD_PATH := "./cmd/deputy")
ifeq ($(words $(MAKECMDGOALS)),1)
	$(MAKE) service ${.DEFAULT_GOAL} --no-print-directory
endif

.FORCE:
.PHONY: test test-cov-html bench bench-cpu bench-cpu-long bin release service
.DEFAULT_GOAL := bin