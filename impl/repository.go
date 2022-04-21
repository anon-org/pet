package impl

import (
	"context"
	"fmt"
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

func (r *repository) FetchByID(ctx context.Context, id string) (*api.User, error) {
	r.Lock()
	defer r.Unlock()

	u, ok := r.db[id]

	if !ok {
		return nil, fmt.Errorf("users with ID: %s does not exist", id)
	}

	return u, nil
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

func (r *repository) Patch(ctx context.Context, id, name string) error {
	_, err := r.FetchByID(ctx, id)
	if err != nil {
		return err
	}

	r.db[id].Name = name
	return nil
}

func (r *repository) Destroy(ctx context.Context, id string) error {
	_, err := r.FetchByID(ctx, id)
	if err != nil {
		return err
	}

	delete(r.db, id)
	return nil
}
