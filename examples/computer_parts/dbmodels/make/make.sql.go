// Code generated by sqlc. DO NOT EDIT.
// source: make.sql

package make

import (
	"context"
)

const getByID = `-- name: GetByID :one
SELECT id, name, created_at, updated_at FROM make WHERE id = $1 LIMIT 1
`

func (q *Queries) GetByID(ctx context.Context, id int64) (Make, error) {
	row := q.db.QueryRowContext(ctx, getByID, id)
	var i Make
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
