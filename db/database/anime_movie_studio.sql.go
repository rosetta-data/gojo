// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: anime_movie_studio.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAnimeMovieStudio = `-- name: CreateAnimeMovieStudio :one
INSERT INTO anime_movie_studios (anime_id, studio_id)
VALUES ($1, $2)
RETURNING id, anime_id, studio_id
`

type CreateAnimeMovieStudioParams struct {
	AnimeID  int64       `json:"anime_id"`
	StudioID pgtype.Int4 `json:"studio_id"`
}

func (q *Queries) CreateAnimeMovieStudio(ctx context.Context, arg CreateAnimeMovieStudioParams) (AnimeMovieStudio, error) {
	row := q.db.QueryRow(ctx, createAnimeMovieStudio, arg.AnimeID, arg.StudioID)
	var i AnimeMovieStudio
	err := row.Scan(&i.ID, &i.AnimeID, &i.StudioID)
	return i, err
}

const deleteAnimeMovieStudio = `-- name: DeleteAnimeMovieStudio :exec
DELETE FROM anime_movie_studios
WHERE anime_id = $1 AND studio_id = $2
`

type DeleteAnimeMovieStudioParams struct {
	AnimeID  int64       `json:"anime_id"`
	StudioID pgtype.Int4 `json:"studio_id"`
}

func (q *Queries) DeleteAnimeMovieStudio(ctx context.Context, arg DeleteAnimeMovieStudioParams) error {
	_, err := q.db.Exec(ctx, deleteAnimeMovieStudio, arg.AnimeID, arg.StudioID)
	return err
}

const getAnimeMovieStudio = `-- name: GetAnimeMovieStudio :one
SELECT id, anime_id, studio_id FROM anime_movie_studios
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAnimeMovieStudio(ctx context.Context, id int64) (AnimeMovieStudio, error) {
	row := q.db.QueryRow(ctx, getAnimeMovieStudio, id)
	var i AnimeMovieStudio
	err := row.Scan(&i.ID, &i.AnimeID, &i.StudioID)
	return i, err
}

const listAnimeMovieStudios = `-- name: ListAnimeMovieStudios :many
SELECT studio_id
FROM anime_movie_studios
WHERE anime_id = $1
LIMIT $2
OFFSET $3
`

type ListAnimeMovieStudiosParams struct {
	AnimeID int64 `json:"anime_id"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

func (q *Queries) ListAnimeMovieStudios(ctx context.Context, arg ListAnimeMovieStudiosParams) ([]pgtype.Int4, error) {
	rows, err := q.db.Query(ctx, listAnimeMovieStudios, arg.AnimeID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []pgtype.Int4{}
	for rows.Next() {
		var studio_id pgtype.Int4
		if err := rows.Scan(&studio_id); err != nil {
			return nil, err
		}
		items = append(items, studio_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}