package storage

import (
	"context"
	"errors"
	"time"

	"github.com/Time-Tracker/internal"
)

var ErrNotExist = errors.New("item not exist in storage")

type Users interface {
	Create(ctx context.Context, name string, passwordHash string) (id int, err error)
	UpdatePassword(ctx context.Context, id int, passwordHash string) (err error)
	Update(ctx context.Context, id int, name string) (err error)
	GetId(ctx context.Context, name string, passwordHash string) (id int, err error)
	Delete(ctx context.Context, id int) (err error)
}

type Timers interface {
	Create(ctx context.Context, userId int, name string) (id int, err error)
	Toggle(ctx context.Context, id int, startTime time.Time, workTime time.Duration) (err error)
	Update(ctx context.Context, id int, name string) (err error)
	GetById(ctx context.Context, id int) (name string, startTime time.Time, workTime time.Duration, err error)
	GetByUserId(ctx context.Context, userId int) (timers []internal.Timer, err error)
	Delete(ctx context.Context, id int) (err error)
}

type Storage struct {
	Users
	Timers
}
