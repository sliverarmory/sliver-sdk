CARGO ?= cargo
EXECUTABLES = $(CARGO)
K := $(foreach exec,$(EXECUTABLES),\
        $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH")))

rename.wasm:
	$(CARGO) build --release --target wasm32-unknown-unknown && cp target/wasm32-unknown-unknown/release/rename.wasm ./rename.wasm

clean:
	rm -f rename.wasm
