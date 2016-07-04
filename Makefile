# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: gele gele-cross evm all test travis-test-with-coverage xgo clean
.PHONY: gele-linux gele-linux-386 gele-linux-amd64
.PHONY: gele-linux-arm gele-linux-arm-5 gele-linux-arm-6 gele-linux-arm-7 gele-linux-arm64
.PHONY: gele-darwin gele-darwin-386 gele-darwin-amd64
.PHONY: gele-windows gele-windows-386 gele-windows-amd64
.PHONY: gele-android gele-ios

GOBIN = build/bin
GO ?= latest

gele:
	build/env.sh go build -i -v $(shell build/flags.sh) -o $(GOBIN)/gele ./cmd/gele
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gele\" to launch gele."

gele-cross: gele-linux gele-darwin gele-windows gele-android gele-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/gele-*

gele-linux: gele-linux-386 gele-linux-amd64 gele-linux-arm
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-*

gele-linux-386: xgo
	build/env.sh $(GOBIN)/xgo --go=$(GO) --dest=$(GOBIN) --targets=linux/386 -v $(shell build/flags.sh) ./cmd/gele
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-* | grep 386

gele-linux-amd64: xgo
	build/env.sh $(GOBIN)/xgo --go=$(GO) --dest=$(GOBIN) --targets=linux/amd64 -v $(shell build/flags.sh) ./cmd/gele
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-* | grep amd64

gele-linux-arm: gele-linux-arm-5 gele-linux-arm-6 gele-linux-arm-7 gele-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-* | grep arm

gele-linux-arm-5: xgo
	build/env.sh $(GOBIN)/xgo --go=$(GO) --dest=$(GOBIN) --targets=linux/arm-5 -v $(shell build/flags.sh) ./cmd/gele
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-* | grep arm-5

gele-linux-arm-6: xgo
	build/env.sh $(GOBIN)/xgo --go=$(GO) --dest=$(GOBIN) --targets=linux/arm-6 -v $(shell build/flags.sh) ./cmd/gele
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-* | grep arm-6

gele-linux-arm-7: xgo
	build/env.sh $(GOBIN)/xgo --go=$(GO) --dest=$(GOBIN) --targets=linux/arm-7 -v $(shell build/flags.sh) ./cmd/gele
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-* | grep arm-7

gele-linux-arm64: xgo
	build/env.sh $(GOBIN)/xgo --go=$(GO) --dest=$(GOBIN) --targets=linux/arm64 -v $(shell build/flags.sh) ./cmd/gele
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-* | grep arm64

gele-darwin: gele-darwin-386 gele-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/gele-darwin-*

gele-darwin-386: xgo
	build/env.sh $(GOBIN)/xgo --go=$(GO) --dest=$(GOBIN) --targets=darwin/386 -v $(shell build/flags.sh) ./cmd/gele
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/gele-darwin-* | grep 386

gele-darwin-amd64: xgo
	build/env.sh $(GOBIN)/xgo --go=$(GO) --dest=$(GOBIN) --targets=darwin/amd64 -v $(shell build/flags.sh) ./cmd/gele
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gele-darwin-* | grep amd64

gele-windows: gele-windows-386 gele-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/gele-windows-*

gele-windows-386: xgo
	build/env.sh $(GOBIN)/xgo --go=$(GO) --dest=$(GOBIN) --targets=windows/386 -v $(shell build/flags.sh) ./cmd/gele
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/gele-windows-* | grep 386

gele-windows-amd64: xgo
	build/env.sh $(GOBIN)/xgo --go=$(GO) --dest=$(GOBIN) --targets=windows/amd64 -v $(shell build/flags.sh) ./cmd/gele
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gele-windows-* | grep amd64

gele-android: xgo
	build/env.sh $(GOBIN)/xgo --go=$(GO) --dest=$(GOBIN) --targets=android-21/aar -v $(shell build/flags.sh) ./cmd/gele
	@echo "Android cross compilation done:"
	@ls -ld $(GOBIN)/gele-android-*

gele-ios: xgo
	build/env.sh $(GOBIN)/xgo --go=$(GO) --dest=$(GOBIN) --targets=ios-7.0/framework -v $(shell build/flags.sh) ./cmd/gele
	@echo "iOS framework cross compilation done:"
	@ls -ld $(GOBIN)/gele-ios-*

evm:
	build/env.sh $(GOROOT)/bin/go install -v $(shell build/flags.sh) ./cmd/evm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/evm to start the evm."

all:
	for cmd in `ls ./cmd/`; do \
		 build/env.sh go build -i -v $(shell build/flags.sh) -o $(GOBIN)/$$cmd ./cmd/$$cmd; \
	done

test: all
	build/env.sh go test ./...

travis-test-with-coverage: all
	build/env.sh go vet ./...
	build/env.sh build/test-global-coverage.sh

xgo:
	build/env.sh go get github.com/karalabe/xgo

clean:
	rm -fr build/_workspace/pkg/ Godeps/_workspace/pkg $(GOBIN)/*
