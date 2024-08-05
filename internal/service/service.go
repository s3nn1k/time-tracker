package service

import (
	"errors"

	"github.com/s3nn1k/time-tracker/internal/storage"
)

var (
	errAlreadyRunning = errors.New("timer is already running")
	errNotRunning     = errors.New("timer is not running yet")
)

type Service struct {
	*Users
	*Timers
}

func New(s *storage.Storage) *Service {
	return &Service{
		Users:  newUsers(s.Users),
		Timers: newTimers(s.Timers),
	}
}
