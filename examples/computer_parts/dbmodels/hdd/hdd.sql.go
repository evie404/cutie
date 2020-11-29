// Code generated by sqlc. DO NOT EDIT.
// source: hdd.sql

package hdd

import (
	"context"
)

const getByID = `-- name: GetByID :one
SELECT id, make_id, name, size_bytes, rpm, created_at, updated_at FROM hdd WHERE id = $1 LIMIT 1
`

func (q *Queries) GetByID(ctx context.Context, id int64) (HDD, error) {
	row := q.db.QueryRowContext(ctx, getByID, id)
	var i HDD
	err := row.Scan(
		&i.ID,
		&i.MakeID,
		&i.Name,
		&i.SizeBytes,
		&i.Rpm,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
