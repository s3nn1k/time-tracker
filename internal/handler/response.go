package handler

import (
	"fmt"
	"time"

	"github.com/Time-Tracker/internal"
)

const timeLayout = "2006-01-02 03:04PM"

type Response struct {
	Token   string  `json:"token,omitempty"`
	TimerId int     `json:"timer_id,omitempty"`
	Timers  []Timer `json:"timers,omitempty"`
}

type Timer struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	LastStart string `json:"last_start,omitempty"`
	WorkTime  string `json:"work_time"`
}

func convertTimers(timers []internal.Timer) (cleanTimers []Timer) {
	for _, timer := range timers {
		cleanTimer := Timer{
			Id:   timer.Id,
			Name: timer.Name,
		}

		if !timer.LastStart.IsZero() {
			cleanTimer.LastStart = timer.LastStart.Format(timeLayout)
		}

		cleanTimer.WorkTime = fmtDuration(timer.WorkTime)

		cleanTimers = append(cleanTimers, cleanTimer)
	}

	return cleanTimers
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%02dh%02dm%02ds", h, m, s)
}
