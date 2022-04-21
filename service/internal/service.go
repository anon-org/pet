package internal

import (
	"context"
	"errors"
	"github.com/anon-org/pet/api"
	"strings"
)

type Service struct {
	repo api.Repository
}

func (s Service) Fetch(ctx context.Context) (api.Users, error) {
	return s.repo.Fetch(ctx), nil
}

func (s Service) Store(ctx context.Context, name string) (*api.User, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("name can't be empty")
	}

	return s.repo.Store(ctx, name)
}
