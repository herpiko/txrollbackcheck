package tx

import (
	"context"
	"database/sql"
)

var (
	ctx context.Context
	db  *sql.DB
	tx  *sql.Tx
)
