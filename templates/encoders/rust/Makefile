CARGO ?= cargo
EXECUTABLES = $(CARGO)
K := $(foreach exec,$(EXECUTABLES),\
        $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH")))

traffic.wasm:
	$(CARGO) build --release --target wasm32-unknown-unknown && cp target/wasm32-unknown-unknown/release/traffic.wasm ./traffic.wasm

clean:
	rm -f traffic.wasm
