# Sliver extension template for Rust

This is a template for a Sliver extensions written in Rust. It is intended to be used as a base for new extensions.

## Repository structure

- `src/lib.rs`: Entrypoint of the extension is declared there. It is named `start` by default, but can be renamed to anything you want.
- `src/main.rs`: Use this file to locally test your code without needing to build a DLL.
- `src/{{.ExtensionName}}.rs`: Put your extension logic in there.

## Building

To build the extension, run `cargo build --lib --release` in the root directory of the repository. The resulting DLL will be located in `target/release/`.
To debug the extension, run `cargo build` instead. The resulting DLL and executable will be located in `target/debug/`.