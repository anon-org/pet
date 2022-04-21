package impl

import (
	"context"
	"errors"
	"github.com/anon-org/arithmetic/api"
	"strings"
)

type service struct {
	repo api.Repository
}

func (s service) Fetch(ctx context.Context) (api.Users, error) {
	return s.repo.Fetch(ctx), nil
}

func (s service) FetchByID(ctx context.Context, id string) (*api.User, error) {
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("id can't be empty")
	}

	return s.repo.FetchByID(ctx, id)
}

func (s service) Store(ctx context.Context, name string) (*api.User, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("name can't be empty")
	}

	return s.repo.Store(ctx, name)
}

func (s service) Patch(ctx context.Context, id, name string) error {
	if strings.TrimSpace(id) == "" {
		return errors.New("id can't be empty")
	}

	if strings.TrimSpace(name) == "" {
		return errors.New("name can't be empty")
	}

	return s.repo.Patch(ctx, id, name)
}

func (s service) Destroy(ctx context.Context, id string) error {
	if strings.TrimSpace(id) == "" {
		return errors.New("id can't be empty")
	}

	return s.repo.Destroy(ctx, id)
}
