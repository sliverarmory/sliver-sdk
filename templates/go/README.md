# Sliver extension template for Go

This is a template for a Sliver extensions written in Go. It is intended to be used as a base for new extensions.

## How to use this template

```shell
$ rm go.mod # or modify the module name in go.mod
$ go mod init <your-full-package-path>
```

## Repository structure

- `pkg/myextension`: Implement your logic in there, and export functions that will be called by the entrypoint
- `dll/main.go`: The extension entrypoint, which contains some helper functions to parse arguments and send data back to the implant
- `main.go`: Use this file to locally test your code without needing to build a DLL