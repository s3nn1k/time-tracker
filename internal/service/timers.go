package service

import (
	"context"
	"time"

	"github.com/s3nn1k/time-tracker/internal"
	"github.com/s3nn1k/time-tracker/internal/storage"
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

func (t *Timers) Start(ctx context.Context, id int, userId int) (err error) {
	_, startTime, workTime, err := t.storage.GetById(ctx, id, userId)
	if err != nil {
		return err
	}

	if !startTime.IsZero() {
		return errAlreadyRunning
	}

	startTime = time.Now()

	err = t.storage.Toggle(ctx, id, userId, startTime, workTime)
	if err != nil {
		return err
	}

	return nil
}

func (t *Timers) Stop(ctx context.Context, id int, userId int) (err error) {
	_, startTime, workTime, err := t.storage.GetById(ctx, id, userId)
	if err != nil {
		return err
	}

	if startTime.IsZero() {
		return errNotRunning
	}

	workTime = time.Since(startTime) + workTime
	startTime = time.Time{}

	err = t.storage.Toggle(ctx, id, userId, startTime, workTime)
	if err != nil {
		return err
	}

	return nil
}

func (t *Timers) Update(ctx context.Context, id int, userId int, name string) (err error) {
	return t.storage.Update(ctx, id, userId, name)
}

func (t *Timers) GetById(ctx context.Context, id int, userId int) (name string, startTime time.Time, workTime time.Duration, err error) {
	return t.storage.GetById(ctx, id, userId)
}

func (t *Timers) GetByUserId(ctx context.Context, userId int) (timers []internal.Timer, err error) {
	return t.storage.GetByUserId(ctx, userId)
}

func (t *Timers) Delete(ctx context.Context, id int, userId int) (err error) {
	return t.storage.Delete(ctx, id, userId)
}
