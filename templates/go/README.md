# Sliver extension template for Go

This is a template for a Sliver extensions written in Go. It is intended to be used as a base for new extensions.

## Repository structure

- `pkg/<EXTENSION_NAME>`: Implement your logic in there, and export functions that will be called by the entrypoint.
- `dll/main.go`: The extension entrypoint, which contains some helper functions to parse arguments and send data back to the implant.
- `main.go`: Use this file to locally test your code without needing to build a DLL.

## How to use this repo template?

```shell
# Make sure to download all dependencies
$ go mod tidy
# Setup cross-compiler for CGO
$ export CC_X64=$(which x86_64-w64-mingw32-gcc)
$ export CC_X86=$(which i686-w64-mingw32-gcc)
# Build the DLLs for both amd64 and 386 architectures
$ make build
```

Build artifacts will be available in the `build` directory.