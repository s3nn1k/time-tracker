package internal

import "time"

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

type Timer struct {
	Id        int           `json:"id"`
	Name      string        `json:"name"`
	LastStart time.Time     `json:"last_start,omitempty"`
	WorkTime  time.Duration `json:"work_time"`
}
