package impl

import (
	"context"
	"github.com/anon-org/arithmetic/api"
	"github.com/google/uuid"
	"sync"
)

type repository struct {
	sync.RWMutex
	db map[string]*api.User
}

func (r *repository) Fetch(ctx context.Context) api.Users {
	r.Lock()
	defer r.Unlock()

	users := make([]*api.User, len(r.db))
	i := 0
	for _, v := range r.db {
		users[i] = v
		i++
	}

	return users
}

func (r *repository) Store(ctx context.Context, name string) (*api.User, error) {
	r.Lock()
	defer r.Unlock()

	id := uuid.NewString()
	u := &api.User{
		ID:   id,
		Name: name,
	}

	r.db[id] = u

	return u, nil
}
