package impl

import (
	"github.com/anon-org/arithmetic/api"
	"github.com/google/wire"
	"sync"
)

var (
	repo     *repository
	repoOnce sync.Once

	svc         *service
	svcOnce     sync.Once
	ProviderSet wire.ProviderSet = wire.NewSet(
		ProvideRepository,
		ProvideService,

		wire.Bind(new(api.Repository), new(*repository)),
		wire.Bind(new(api.Service), new(*service)),
	)
)

func ProvideRepository(db map[string]*api.User) *repository {
	repoOnce.Do(func() {
		repo = &repository{
			db: db,
		}
	})

	return repo
}

func ProvideService(repo api.Repository) *service {
	svcOnce.Do(func() {
		svc = &service{
			repo: repo,
		}
	})

	return svc
}
