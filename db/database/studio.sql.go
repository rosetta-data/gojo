// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: studio.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createStudio = `-- name: CreateStudio :one
INSERT INTO studios (studio_name)
VALUES ($1)
ON CONFLICT (studio_name)
DO UPDATE SET studio_name = excluded.studio_name
RETURNING  id, studio_name, created_at
`

func (q *Queries) CreateStudio(ctx context.Context, studioName string) (Studio, error) {
	row := q.db.QueryRow(ctx, createStudio, studioName)
	var i Studio
	err := row.Scan(&i.ID, &i.StudioName, &i.CreatedAt)
	return i, err
}

const deleteStudio = `-- name: DeleteStudio :exec
DELETE FROM studios
WHERE id = $1
`

func (q *Queries) DeleteStudio(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteStudio, id)
	return err
}

const getStudio = `-- name: GetStudio :one
SELECT id, studio_name, created_at FROM studios
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetStudio(ctx context.Context, id int32) (Studio, error) {
	row := q.db.QueryRow(ctx, getStudio, id)
	var i Studio
	err := row.Scan(&i.ID, &i.StudioName, &i.CreatedAt)
	return i, err
}

const listStudios = `-- name: ListStudios :many
SELECT id FROM studios
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListStudiosParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListStudios(ctx context.Context, arg ListStudiosParams) ([]int32, error) {
	rows, err := q.db.Query(ctx, listStudios, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int32{}
	for rows.Next() {
		var id int32
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateStudio = `-- name: UpdateStudio :one
UPDATE studios
SET
    studio_name = COALESCE($2, studio_name)
WHERE id = $1
RETURNING id, studio_name, created_at
`

type UpdateStudioParams struct {
	ID         int32
	StudioName pgtype.Text
}

func (q *Queries) UpdateStudio(ctx context.Context, arg UpdateStudioParams) (Studio, error) {
	row := q.db.QueryRow(ctx, updateStudio, arg.ID, arg.StudioName)
	var i Studio
	err := row.Scan(&i.ID, &i.StudioName, &i.CreatedAt)
	return i, err
}
