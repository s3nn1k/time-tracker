package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Time-Tracker/internal"
	"github.com/Time-Tracker/internal/storage"
)

type Timers struct {
	db *sql.DB
}

func newTimers(db *sql.DB) *Timers {
	return &Timers{
		db: db,
	}
}

func (t *Timers) Create(ctx context.Context, userId int, name string) (id int, err error) {
	q := fmt.Sprintf("INSERT INTO %s (user_id, name, start_time, work_time) VALUES ($1, $2, $3, $4) RETURNING id", timersTable)

	err = t.db.QueryRowContext(ctx, q, userId, name, time.Time{}.Unix(), 0).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (t *Timers) Toggle(ctx context.Context, id int, startTime time.Time, workTime time.Duration) (err error) {
	q := fmt.Sprintf("UPDATE %s SET start_time=$1, work_time=$2 WHERE id=$3", timersTable)

	res, err := t.db.ExecContext(ctx, q, startTime.Unix(), int64(workTime.Seconds()), id)
	if err != nil {
		return err
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		return storage.ErrNotExist
	}

	return nil
}

func (t *Timers) Update(ctx context.Context, id int, name string) (err error) {
	q := fmt.Sprintf("UPDATE %s SET name=$1 WHERE id=$2", timersTable)

	res, err := t.db.ExecContext(ctx, q, name, id)
	if err != nil {
		return err
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		return storage.ErrNotExist
	}

	return nil
}

func (t *Timers) GetById(ctx context.Context, id int) (name string, startTime time.Time, workTime time.Duration, err error) {
	q := fmt.Sprintf("SELECT name, start_time, work_time FROM %s WHERE id=$1", timersTable)

	var unixStartTime int64
	var secWorkTime int64

	err = t.db.QueryRowContext(ctx, q, id).Scan(&name, &unixStartTime, &secWorkTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", time.Time{}, 0, storage.ErrNotExist
		}

		return "", time.Time{}, 0, err
	}

	startTime = time.Unix(unixStartTime, 0)
	workTime = time.Duration(secWorkTime * int64(time.Second))

	return name, startTime, workTime, nil
}

func (t *Timers) GetByUserId(ctx context.Context, userId int) (timers []internal.Timer, err error) {
	q := fmt.Sprintf("SELECT id, name, start_time, work_time FROM %s WHERE user_id=$1", timersTable)

	rows, err := t.db.QueryContext(ctx, q, userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var timer internal.Timer
		var unixStartTime int64
		var secWorkTime int64

		err = rows.Scan(&timer.Id, &timer.Name, &unixStartTime, &secWorkTime)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, storage.ErrNotExist
			}

			return nil, err
		}

		timer.LastStart = time.Unix(unixStartTime, 0)
		timer.WorkTime = time.Duration(secWorkTime * int64(time.Second))

		timers = append(timers, timer)
	}

	return timers, nil
}

func (t *Timers) Delete(ctx context.Context, id int) (err error) {
	q := fmt.Sprintf("DELETE FROM %s WHERE id=$1", timersTable)

	res, err := t.db.ExecContext(ctx, q, id)
	if err != nil {
		return err
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		return storage.ErrNotExist
	}

	return nil
}
