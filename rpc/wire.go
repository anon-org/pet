//go:build wireinject

package rpc

import (
	"github.com/google/wire"
	"time"
)

func Wire(baseURL string, timeout time.Duration) *rpc {
	panic(wire.Build(ProviderSet))
}
