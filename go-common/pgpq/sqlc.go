package pgpq

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// PGXPool returns a prepared connection pool, compatible with sqlc generated structs.
func PGXPool(ctx context.Context, dbUri string) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, dbUri)
}
