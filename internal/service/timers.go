package service

import (
	"context"
	"time"

	"github.com/Time-Tracker/internal"
	"github.com/Time-Tracker/internal/storage"
)

type Timers struct {
	storage storage.Timers
}

func newTimers(s storage.Timers) *Timers {
	return &Timers{
		storage: s,
	}
}

func (t *Timers) Create(ctx context.Context, userId int, name string) (id int, err error) {
	return t.storage.Create(ctx, userId, name)
}

func (t *Timers) Start(ctx context.Context, id int) (err error) {
	_, startTime, workTime, err := t.storage.GetById(ctx, id)
	if err != nil {
		return err
	}

	if !startTime.IsZero() {
		return errAlreadyRunning
	}

	startTime = time.Now()

	err = t.storage.Toggle(ctx, id, startTime, workTime)
	if err != nil {
		return err
	}

	return nil
}

func (t *Timers) Stop(ctx context.Context, id int) (err error) {
	_, startTime, workTime, err := t.storage.GetById(ctx, id)
	if err != nil {
		return err
	}

	if startTime.IsZero() {
		return errNotRunning
	}

	workTime = time.Since(startTime) + workTime
	startTime = time.Time{}

	err = t.storage.Toggle(ctx, id, startTime, workTime)
	if err != nil {
		return err
	}

	return nil
}

func (t *Timers) Update(ctx context.Context, id int, name string) (err error) {
	return t.storage.Update(ctx, id, name)
}

func (t *Timers) GetById(ctx context.Context, id int) (name string, startTime time.Time, workTime time.Duration, err error) {
	return t.storage.GetById(ctx, id)
}

func (t *Timers) GetByUserId(ctx context.Context, userId int) (timers []internal.Timer, err error) {
	return t.storage.GetByUserId(ctx, userId)
}

func (t *Timers) Delete(ctx context.Context, id int) (err error) {
	return t.storage.Delete(ctx, id)
}
