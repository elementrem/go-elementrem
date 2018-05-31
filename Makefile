# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: gele android ios gele-cross swarm evm all test clean
.PHONY: gele-linux gele-linux-386 gele-linux-amd64 gele-linux-mips64 gele-linux-mips64le
.PHONY: gele-linux-arm gele-linux-arm-5 gele-linux-arm-6 gele-linux-arm-7 gele-linux-arm64
.PHONY: gele-darwin gele-darwin-386 gele-darwin-amd64
.PHONY: gele-windows gele-windows-386 gele-windows-amd64

GOBIN = $(shell pwd)/build/bin
GO ?= latest

gele:
	build/env.sh go run build/ci.go install ./cmd/gele
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gele\" to launch gele."

swarm:
	build/env.sh go run build/ci.go install ./cmd/swarm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/swarm\" to launch swarm."

all:
	build/env.sh go run build/ci.go install

android:
	build/env.sh go run build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/gele.aar\" to use the library."

ios:
	build/env.sh go run build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/Gele.framework\" to use the library."

test: all
	build/env.sh go run build/ci.go test

lint: ## Run linters.
	build/env.sh go run build/ci.go lint

clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/kevinburke/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go get -u github.com/golang/protobuf/protoc-gen-go
	env GOBIN= go install ./cmd/abigen
	@type "npm" 2> /dev/null || echo 'Please install node.js and npm'
	@type "solc" 2> /dev/null || echo 'Please install solc'
	@type "protoc" 2> /dev/null || echo 'Please install protoc'

# Cross Compilation Targets (xgo)

gele-cross: gele-linux gele-darwin gele-windows gele-android gele-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/gele-*

gele-linux: gele-linux-386 gele-linux-amd64 gele-linux-arm gele-linux-mips64 gele-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-*

gele-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/gele
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-* | grep 386

gele-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/gele
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-* | grep amd64

gele-linux-arm: gele-linux-arm-5 gele-linux-arm-6 gele-linux-arm-7 gele-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-* | grep arm

gele-linux-arm-5:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/gele
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-* | grep arm-5

gele-linux-arm-6:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/gele
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-* | grep arm-6

gele-linux-arm-7:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/gele
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-* | grep arm-7

gele-linux-arm64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/gele
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-* | grep arm64

gele-linux-mips:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/gele
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-* | grep mips

gele-linux-mipsle:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/gele
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-* | grep mipsle

gele-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/gele
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-* | grep mips64

gele-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/gele
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/gele-linux-* | grep mips64le

gele-darwin: gele-darwin-386 gele-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/gele-darwin-*

gele-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/gele
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/gele-darwin-* | grep 386

gele-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/gele
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gele-darwin-* | grep amd64

gele-windows: gele-windows-386 gele-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/gele-windows-*

gele-windows-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/gele
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/gele-windows-* | grep 386

gele-windows-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/gele
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gele-windows-* | grep amd64
