//go:build wireinject

package main

import (
	"github.com/anon-org/pet/api"
	"github.com/anon-org/pet/service/internal"
	"github.com/google/wire"
)

func Wire(db map[string]*api.User) *internal.Service {
	panic(wire.Build(internal.ProviderSet))
}
