package service

import (
	"context"

	"github.com/Time-Tracker/internal/storage"
)

type Users struct {
	storage storage.Users
}

func newUsers(s storage.Users) *Users {
	return &Users{
		storage: s,
	}
}

func (u *Users) Create(ctx context.Context, name string, passwordHash string) (id int, err error) {
	return u.storage.Create(ctx, name, passwordHash)
}

func (u *Users) UpdatePassword(ctx context.Context, id int, passwordHash string) (err error) {
	return u.storage.UpdatePassword(ctx, id, passwordHash)
}

func (u *Users) Update(ctx context.Context, id int, name string) (err error) {
	return u.storage.Update(ctx, id, name)
}

func (u *Users) GetId(ctx context.Context, name string, passwordHash string) (id int, err error) {
	return u.storage.GetId(ctx, name, passwordHash)
}

func (u *Users) Delete(ctx context.Context, id int) (err error) {
	return u.storage.Delete(ctx, id)
}
