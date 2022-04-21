//go:build wireinject

package main

import (
	"github.com/anon-org/pet/api"
	"github.com/anon-org/pet/service/internal/impl"
	"github.com/google/wire"
)

func Wire(db map[string]*api.User) *impl.Service {
	panic(wire.Build(impl.ProviderSet))
}
