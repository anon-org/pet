package api

import "context"

type Service interface {
	Fetch(ctx context.Context) (Users, error)
	FetchByID(ctx context.Context, id string) (*User, error)
	Store(ctx context.Context, name string) (*User, error)
	Patch(ctx context.Context, id, name string) error
	Destroy(ctx context.Context, id string) error
}
