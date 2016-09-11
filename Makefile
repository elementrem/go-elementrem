# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: gele gele-cross evm all test clean
.PHONY: gele-linux gele-linux-386 gele-linux-amd64 gele-linux-mips64 gele-linux-mips64le
.PHONY: gele-linux-arm gele-linux-arm-5 gele-linux-arm-6 gele-linux-arm-7 gele-linux-arm64
.PHONY: gele-darwin gele-darwin-386 gele-darwin-amd64
.PHONY: gele-windows gele-windows-386 gele-windows-amd64
.PHONY: gele-android gele-ios

GOBIN = build/bin
GO ?= latest

gele:
	build/env.sh go run build/ci.go install ./cmd/gele
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gele\" to launch gele."

evm:
	build/env.sh go run build/ci.go install ./cmd/evm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/evm to start the evm."

all:
	build/env.sh go run build/ci.go install

test: all
	build/env.sh go run build/ci.go test

clean:
	rm -fr build/_workspace/pkg/ Godeps/_workspace/pkg $(GOBIN)/*

# Cross Compilation Targets (xgo)

gele-cross: gele-linux gele-darwin gele-windows gele-android gele-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/gele-*

gele-linux: gele-linux-386 gele-linux-amd64 gele-linux-arm gele-linux-mips64 gele-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-*

gele-linux-386:
	build/env.sh go run build/ci.go xgo --go=$(GO) --dest=$(GOBIN) --targets=linux/386 -v ./cmd/gele
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-* | grep 386

gele-linux-amd64:
	build/env.sh go run build/ci.go xgo --go=$(GO) --dest=$(GOBIN) --targets=linux/amd64 -v ./cmd/gele
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-* | grep amd64

gele-linux-arm: gele-linux-arm-5 gele-linux-arm-6 gele-linux-arm-7 gele-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-* | grep arm

gele-linux-arm-5:
	build/env.sh go run build/ci.go xgo --go=$(GO) --dest=$(GOBIN) --targets=linux/arm-5 -v ./cmd/gele
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-* | grep arm-5

gele-linux-arm-6:
	build/env.sh go run build/ci.go xgo --go=$(GO) --dest=$(GOBIN) --targets=linux/arm-6 -v ./cmd/gele
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-* | grep arm-6

gele-linux-arm-7:
	build/env.sh go run build/ci.go xgo --go=$(GO) --dest=$(GOBIN) --targets=linux/arm-7 -v ./cmd/gele
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-* | grep arm-7

gele-linux-arm64:
	build/env.sh go run build/ci.go xgo --go=$(GO) --dest=$(GOBIN) --targets=linux/arm64 -v ./cmd/gele
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-* | grep arm64
	
gele-linux-mips64:
	build/env.sh go run build/ci.go xgo --go=$(GO) --dest=$(GOBIN) --targets=linux/mips64 -v ./cmd/gele
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-* | grep mips64

gele-linux-mips64le:
	build/env.sh go run build/ci.go xgo --go=$(GO) --dest=$(GOBIN) --targets=linux/mips64le -v ./cmd/gele
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-* | grep mips64le

gele-darwin: gele-darwin-386 gele-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/gele-darwin-*

gele-darwin-386:
	build/env.sh go run build/ci.go xgo --go=$(GO) --dest=$(GOBIN) --targets=darwin/386 -v ./cmd/gele
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/gele-darwin-* | grep 386

gele-darwin-amd64:
	build/env.sh go run build/ci.go xgo --go=$(GO) --dest=$(GOBIN) --targets=darwin/amd64 -v ./cmd/gele
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gele-darwin-* | grep amd64

gele-windows: gele-windows-386 gele-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/gele-windows-*

gele-windows-386:
	build/env.sh go run build/ci.go xgo --go=$(GO) --dest=$(GOBIN) --targets=windows/386 -v ./cmd/gele
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/gele-windows-* | grep 386

gele-windows-amd64:
	build/env.sh go run build/ci.go xgo --go=$(GO) --dest=$(GOBIN) --targets=windows/amd64 -v ./cmd/gele
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gele-windows-* | grep amd64

gele-android:
	build/env.sh go run build/ci.go xgo --go=$(GO) --dest=$(GOBIN) --targets=android-21/aar -v ./cmd/gele
	@echo "Android cross compilation done:"
	@ls -ld $(GOBIN)/gele-android-*

gele-ios:
	build/env.sh go run build/ci.go xgo --go=$(GO) --dest=$(GOBIN) --targets=ios-7.0/framework -v ./cmd/gele
	@echo "iOS framework cross compilation done:"
	@ls -ld $(GOBIN)/gele-ios-*