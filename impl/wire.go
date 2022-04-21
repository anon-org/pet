//go:build wireinject

package impl

import (
	"github.com/anon-org/pet/api"
	"github.com/google/wire"
)

func Wire(db map[string]*api.User) *service {
	panic(wire.Build(ProviderSet))
}
