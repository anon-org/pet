package api

import "context"

type Repository interface {
	Fetch(ctx context.Context) Users
	Store(ctx context.Context, name string) (*User, error)
}
