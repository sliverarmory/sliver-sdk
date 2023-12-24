GO ?= go
CGO_ENABLED = 1
CC_X64 ?= x86_64-w64-mingw32-gcc
CC_X86 ?= i686-w64-mingw32-gcc
EXT_NAME ?= {{.ExtensionName}}

.PHONY: all
all: debug build

.PHONY: build
build: build_amd64 build_386
	@mkdir -p build

.PHONY: build_amd64
build_amd64:
	CC=$(CC_X64) CGO_ENABLED=$(CGO_ENABLED) GOOS=windows GOARCH=amd64 $(GO) build -o build/$(EXT_NAME).x64.dll -buildmode=c-shared dll/main.go

.PHONY: build_386
build_386:
	CC=$(CC_X86) CGO_ENABLED=$(CGO_ENABLED) GOOS=windows GOARCH=386 $(GO) build -o build/$(EXT_NAME).x86.dll -buildmode=c-shared dll/main.go

.PHONY: debug
debug: debug_amd64 debug_386

.PHONY: debug_amd64
debug_amd64:
	GOOS=windows GOARCH=amd64 $(GO) build -gcflags "-N -l" -o build/$(EXT_NAME).x64.exe

.PHONY: debug_386
debug_386:
	GOOS=windows GOARCH=386 $(GO) build -gcflags "-N -l" -o build/$(EXT_NAME).x86.exe

.PHONY: clean
clean:
	rm -rf build