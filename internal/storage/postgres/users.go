package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/s3nn1k/time-tracker/internal/storage"
)

type Users struct {
	db *sql.DB
}

func newUsers(db *sql.DB) *Users {
	return &Users{
		db: db,
	}
}

func (u *Users) Create(ctx context.Context, name string, passwordHash string) (id int, err error) {
	q := fmt.Sprintf("INSERT INTO %s (name, password_hash) VALUES ($1, $2) RETURNING id", usersTable)

	err = u.db.QueryRowContext(ctx, q, name, passwordHash).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (u *Users) GetId(ctx context.Context, name string, passwordHash string) (id int, err error) {
	q := fmt.Sprintf("SELECT id FROM %s WHERE name=$1 AND password_hash=$2", usersTable)

	err = u.db.QueryRowContext(ctx, q, name, passwordHash).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, storage.ErrNotExist
		}

		return 0, err
	}

	return id, nil
}

func (u *Users) UpdatePassword(ctx context.Context, id int, passwordHash string) (err error) {
	q := fmt.Sprintf("UPDATE %s SET password_hash=$1 WHERE id=$2", usersTable)

	res, err := u.db.ExecContext(ctx, q, passwordHash, id)
	if err != nil {
		return err
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		return storage.ErrNotExist
	}

	return nil
}

func (u *Users) Update(ctx context.Context, id int, name string) (err error) {
	q := fmt.Sprintf("UPDATE %s SET name=$1 WHERE id=$2", usersTable)

	res, err := u.db.ExecContext(ctx, q, name, id)
	if err != nil {
		return err
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		return storage.ErrNotExist
	}

	return nil
}

func (u *Users) Delete(ctx context.Context, id int) (err error) {
	q := fmt.Sprintf("DELETE FROM %s WHERE id=$1", usersTable)

	res, err := u.db.ExecContext(ctx, q, id)
	if err != nil {
		return err
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		return storage.ErrNotExist
	}

	return nil
}
