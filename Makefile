.PHONY: service rpc

service:
	@go run ./service/

rpc:
	@go run ./rpc/example/

wire:
	@go run github.com/google/wire/cmd/wire@v0.5.0 ./*/