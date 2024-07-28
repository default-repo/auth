package pg_db

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	sq "github.com/Masterminds/squirrel"
	"github.com/brianvoe/gofakeit"
	"github.com/default-repo/auth/internal/model"
	"github.com/default-repo/auth/internal/repository"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var _ repository.Repo = (*PGStore)(nil)

const (
	customerTableName = "customer"

	idColumn       = "id"
	uuidColumn     = "uuid"
	nameColumn     = "name"
	passwordColumn = "password"
	emailColumn    = "email"
)

type PGStore struct {
	pool *pgxpool.Pool
	log  slog.Logger
}

func NewPGStore(log slog.Logger, dbDSN string) (*PGStore, error) {
	pool, err := pgxpool.Connect(context.Background(), dbDSN)
	if err != nil {
		return nil, errors.New("could not connect to database: " + err.Error())
	}

	if err = pool.Ping(context.Background()); err != nil {
		return nil, errors.New("could not ping database: " + err.Error())
	}

	return &PGStore{
		pool: pool,
		log:  log,
	}, nil
}
func (db *PGStore) Close() {
	db.pool.Close()
}

func (db *PGStore) Somthing() error { return nil }

func (db *PGStore) InsertData(ctx context.Context, c model.Customer) (int, error) {
	builder := sq.Insert(customerTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(
			uuidColumn,
			nameColumn,
			passwordColumn,
			emailColumn,
		).
		Values(
			c.UUID,
			c.Name,
			c.Password,
			c.Email,
		).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("creating the pg insert query failed: %w", err)
	}

	var lastID int
	err = db.pool.QueryRow(ctx, query, args...).Scan(&lastID)
	if err != nil {
		return 0, fmt.Errorf("pg query row insering failed: %w", err)
	}

	return lastID, nil
}

func (db *PGStore) List(ctx context.Context, limit uint64) (pgx.Rows, error) {
	builder := sq.Select(
		idColumn,
		uuidColumn,
		nameColumn,
		passwordColumn,
		emailColumn,
	).
		PlaceholderFormat(sq.Dollar).
		From(customerTableName).
		OrderBy("id DESC").
		Limit(limit)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("creating list pg query failed: %w", err)
	}

	rows, err := db.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("executing list pg query failed: %w", err)
	}

	return rows, nil
}

func (db *PGStore) UpdateByID(ctc context.Context, ID int) (int64, error) {
	builder := sq.Update(customerTableName).
		PlaceholderFormat(sq.Dollar).
		Set("name", gofakeit.Name()).
		Where(sq.Eq{"id": ID})

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("creating UpdateByID pg query failed: %w", err)
	}

	result, err := db.pool.Exec(ctc, query, args...)
	if err != nil {
		return 0, fmt.Errorf("executing UpdateByID pg query failed: %w", err)
	}

	return result.RowsAffected(), nil
}

func (db *PGStore) GetCustomerByUID(ctx context.Context, ID int) (*model.Customer, error) {
	result := new(model.Customer)

	stmt := fmt.Sprintf("SELECT id, uuid, name, email, password FROM %s WHERE id = $1", customerTableName)

	row := db.pool.QueryRow(ctx, stmt, ID)

	err := row.Scan(&result.ID, &result.UUID, &result.Name, &result.Email, &result.Password)
	if err != nil {
		return nil, fmt.Errorf("scan GetByUID pg query failed: %w", err)
	}

	return result, nil
}
