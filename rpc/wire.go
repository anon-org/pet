//go:build wireinject

package rpc

import (
	"github.com/anon-org/pet/rpc/internal"
	"github.com/google/wire"
	"time"
)

func Wire(baseURL string, timeout time.Duration) *internal.Rpc {
	panic(wire.Build(internal.ProviderSet))
}
