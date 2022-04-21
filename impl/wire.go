//go:build wireinject
// +build wireinject

package impl

import (
	"github.com/anon-org/arithmetic/api"
	"github.com/google/wire"
)

var (
	ProviderSet wire.ProviderSet = wire.NewSet(
		ProvideService,
		ProvideHandler,

		wire.Bind(new(api.Service), new(*Service)),
		wire.Bind(new(api.Handler), new(*Handler)),
	)
)

func ProvideService() *Service {
	return &Service{}
}

func ProvideHandler(service api.Service) *Handler {
	return &Handler{
		Svc: service,
	}
}

func Wire() *Handler {
	panic(wire.Build(ProviderSet))
}
