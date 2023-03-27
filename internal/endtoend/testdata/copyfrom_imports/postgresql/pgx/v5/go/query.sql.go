// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

const deleteValues = `-- name: DeleteValues :execresult
DELETE
FROM myschema.foo
`

func (q *Queries) DeleteValues(ctx context.Context) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, deleteValues)
}

const insertSingleValue = `-- name: InsertSingleValue :exec
INSERT INTO myschema.foo (a) VALUES ($1)
`

func (q *Queries) InsertSingleValue(ctx context.Context, a pgtype.Text) error {
	_, err := q.db.Exec(ctx, insertSingleValue, a)
	return err
}

type InsertValuesParams struct {
	A pgtype.Text
	B pgtype.Int4
}