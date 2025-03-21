package postgres

import (
	"context"

	"github.com/catalystgo/logger/logger"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ DB = (*db)(nil)

type db struct {
	db *pgxpool.Pool
}

func (db *db) Ping(ctx context.Context) error {
	return db.db.Ping(ctx)
}

func (db *db) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	return db.db.Exec(ctx, sql, arguments...)
}

func (db *db) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return db.db.Query(ctx, sql, args...)
}

func (db *db) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return db.db.QueryRow(ctx, sql, args...)
}

func (db *db) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	return db.db.SendBatch(ctx, b)
}

func (db *db) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	return db.db.CopyFrom(ctx, tableName, columnNames, rowSrc)
}

func (db *db) Begin(ctx context.Context, f func(tx pgx.Tx) error) error {
	tx, err := db.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		err := tx.Rollback(ctx)
		if err != nil {
			logger.Errorf(ctx, "rollback tx: %v", err)
		}
	}()

	if err = f(tx); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (db *db) Close() error {
	db.db.Close()
	return nil
}
