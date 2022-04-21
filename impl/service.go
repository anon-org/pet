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

func (s service) Store(ctx context.Context, name string) (*api.User, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("name can't be empty")
	}

	return s.repo.Store(ctx, name)
}
