package database

import (
	"context"

	"github.com/jmoiron/sqlx"
)

func FromContext(ctx context.Context) *sqlx.DB {
	return ctx.Value("db").(*sqlx.DB)
}
