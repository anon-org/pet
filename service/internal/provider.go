package internal

import (
	"github.com/anon-org/pet/api"
	"github.com/google/wire"
	"sync"
)

var (
	repo     *Repository
	repoOnce sync.Once

	svc         *Service
	svcOnce     sync.Once
	ProviderSet wire.ProviderSet = wire.NewSet(
		ProvideRepository,
		ProvideService,

		wire.Bind(new(api.Repository), new(*Repository)),
		wire.Bind(new(api.Service), new(*Service)),
	)
)

func ProvideRepository(db map[string]*api.User) *Repository {
	repoOnce.Do(func() {
		repo = &Repository{
			db: db,
		}
	})

	return repo
}

func ProvideService(repo api.Repository) *Service {
	svcOnce.Do(func() {
		svc = &Service{
			repo: repo,
		}
	})

	return svc
}
