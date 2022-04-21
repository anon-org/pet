package api

import "context"

type Service interface {
	Fetch(ctx context.Context) (Users, error)
	Store(ctx context.Context, name string) (*User, error)
}
