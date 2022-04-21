package rpc

import (
	"github.com/anon-org/pet/api"
	"github.com/google/wire"
	"sync"
	"time"
)

var (
	r     *rpc
	rOnce sync.Once

	ProviderSet wire.ProviderSet = wire.NewSet(
		ProvideRpc,

		wire.Bind(new(api.Service), new(*rpc)),
	)
)

func ProvideRpc(baseURL string, timeout time.Duration) *rpc {
	rOnce.Do(func() {
		r = &rpc{
			BaseURL: baseURL,
			Timeout: timeout,
		}
	})

	return r
}
