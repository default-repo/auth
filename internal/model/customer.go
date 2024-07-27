package model

import "time"

type Customer struct {
	ID        int64      `db:"id"`
	UUID      string     `db:"uuid"`
	Name      string     `db:"name"`
	Password  string     `db:"password"`
	Email     string     `db:"email"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
