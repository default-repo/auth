package model

import "time"

type Customer struct {
	ID        int64      `db:"id"`
	Name      string     `db:"name"`
	Password  string     `db:"password"`
	Email     string     `db:"email"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
