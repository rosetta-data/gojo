// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: anime_serie_seasons.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAnimeSerieSeason = `-- name: CreateAnimeSerieSeason :one
INSERT INTO anime_serie_seasons (anime_id, season_id)
VALUES ($1, $2)
RETURNING id, anime_id, season_id, created_at
`

type CreateAnimeSerieSeasonParams struct {
	AnimeID  int64
	SeasonID int64
}

func (q *Queries) CreateAnimeSerieSeason(ctx context.Context, arg CreateAnimeSerieSeasonParams) (AnimeSerieSeason, error) {
	row := q.db.QueryRow(ctx, createAnimeSerieSeason, arg.AnimeID, arg.SeasonID)
	var i AnimeSerieSeason
	err := row.Scan(
		&i.ID,
		&i.AnimeID,
		&i.SeasonID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAnimeSerieSeason = `-- name: DeleteAnimeSerieSeason :exec
DELETE FROM anime_serie_seasons
WHERE id = $1
`

func (q *Queries) DeleteAnimeSerieSeason(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteAnimeSerieSeason, id)
	return err
}

const getAnimeSerieSeason = `-- name: GetAnimeSerieSeason :one
SELECT id, anime_id, season_id, created_at FROM anime_serie_seasons
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetAnimeSerieSeason(ctx context.Context, id int64) (AnimeSerieSeason, error) {
	row := q.db.QueryRow(ctx, getAnimeSerieSeason, id)
	var i AnimeSerieSeason
	err := row.Scan(
		&i.ID,
		&i.AnimeID,
		&i.SeasonID,
		&i.CreatedAt,
	)
	return i, err
}

const listAnimeSerieSeasons = `-- name: ListAnimeSerieSeasons :many
SELECT season_id FROM anime_serie_seasons
WHERE anime_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListAnimeSerieSeasonsParams struct {
	AnimeID int64
	Limit   int32
	Offset  int32
}

func (q *Queries) ListAnimeSerieSeasons(ctx context.Context, arg ListAnimeSerieSeasonsParams) ([]int64, error) {
	rows, err := q.db.Query(ctx, listAnimeSerieSeasons, arg.AnimeID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var season_id int64
		if err := rows.Scan(&season_id); err != nil {
			return nil, err
		}
		items = append(items, season_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAnimeSerieSeason = `-- name: UpdateAnimeSerieSeason :one
UPDATE anime_serie_seasons
SET
  anime_id = COALESCE($1, anime_id),
  season_id = COALESCE($2, season_id)
WHERE
  id = $3
RETURNING id, anime_id, season_id, created_at
`

type UpdateAnimeSerieSeasonParams struct {
	AnimeID  pgtype.Int8
	SeasonID pgtype.Int8
	ID       int64
}

func (q *Queries) UpdateAnimeSerieSeason(ctx context.Context, arg UpdateAnimeSerieSeasonParams) (AnimeSerieSeason, error) {
	row := q.db.QueryRow(ctx, updateAnimeSerieSeason, arg.AnimeID, arg.SeasonID, arg.ID)
	var i AnimeSerieSeason
	err := row.Scan(
		&i.ID,
		&i.AnimeID,
		&i.SeasonID,
		&i.CreatedAt,
	)
	return i, err
}
