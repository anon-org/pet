package internal

import (
	"github.com/anon-org/pet/api"
	"github.com/google/wire"
	"sync"
	"time"
)

var (
	r     *Rpc
	rOnce sync.Once

	ProviderSet wire.ProviderSet = wire.NewSet(
		ProvideRpc,

		wire.Bind(new(api.Service), new(*Rpc)),
	)
)

func ProvideRpc(baseURL string, timeout time.Duration) *Rpc {
	rOnce.Do(func() {
		r = &Rpc{
			BaseURL: baseURL,
			Timeout: timeout,
		}
	})

	return r
}
