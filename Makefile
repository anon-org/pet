.PHONY: service rpc

service:
	@go run ./service/cmd

rpc:
	@go run ./rpc/cmd/

wire:
	@go run github.com/google/wire/cmd/wire@v0.5.0 ./rpc ./service/