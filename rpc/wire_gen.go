// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package rpc

import (
	"github.com/anon-org/pet/rpc/internal"
	"time"
)

// Injectors from wire.go:

func Wire(baseURL string, timeout time.Duration) *internal.Rpc {
	rpc := internal.ProvideRpc(baseURL, timeout)
	return rpc
}
