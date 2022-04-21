service:
	@go run service/main.go

rpc:
	@go run rpc/cmd/api/main.go

wire:
	@go run github.com/google/wire/cmd/wire@v0.5.0 ./impl ./rpc