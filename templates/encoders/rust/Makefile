CARGO ?= cargo
EXECUTABLES = $(CARGO)
K := $(foreach exec,$(EXECUTABLES),\
        $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH")))

{{.EncoderName}}.wasm:
	$(CARGO) build --release --target wasm32-unknown-unknown && cp target/wasm32-unknown-unknown/release/{{.EncoderName}}.wasm ./{{.EncoderName}}.wasm

clean:
	rm -f {{.EncoderName}}.wasm
