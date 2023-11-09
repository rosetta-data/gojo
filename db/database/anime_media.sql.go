// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: anime_media.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAnimeMedia = `-- name: CreateAnimeMedia :one
INSERT INTO anime_media (media_type, media_url, media_author, media_platform)
VALUES ($1, $2, $3, $4)
RETURNING  id, media_type, media_url, media_author, media_platform, created_at
`

type CreateAnimeMediaParams struct {
	MediaType     string `json:"media_type"`
	MediaUrl      string `json:"media_url"`
	MediaAuthor   string `json:"media_author"`
	MediaPlatform string `json:"media_platform"`
}

func (q *Queries) CreateAnimeMedia(ctx context.Context, arg CreateAnimeMediaParams) (AnimeMedium, error) {
	row := q.db.QueryRow(ctx, createAnimeMedia,
		arg.MediaType,
		arg.MediaUrl,
		arg.MediaAuthor,
		arg.MediaPlatform,
	)
	var i AnimeMedium
	err := row.Scan(
		&i.ID,
		&i.MediaType,
		&i.MediaUrl,
		&i.MediaAuthor,
		&i.MediaPlatform,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAnimeMedia = `-- name: DeleteAnimeMedia :exec
DELETE FROM anime_media
WHERE id = $1
`

func (q *Queries) DeleteAnimeMedia(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteAnimeMedia, id)
	return err
}

const getAnimeMedia = `-- name: GetAnimeMedia :one
SELECT id, media_type, media_url, media_author, media_platform, created_at FROM anime_media
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAnimeMedia(ctx context.Context, id int64) (AnimeMedium, error) {
	row := q.db.QueryRow(ctx, getAnimeMedia, id)
	var i AnimeMedium
	err := row.Scan(
		&i.ID,
		&i.MediaType,
		&i.MediaUrl,
		&i.MediaAuthor,
		&i.MediaPlatform,
		&i.CreatedAt,
	)
	return i, err
}

const updateAnimeMedia = `-- name: UpdateAnimeMedia :one
UPDATE anime_media
SET
  media_type = COALESCE($1, media_type),
  media_url = COALESCE($2, media_url),
  media_author = COALESCE($3, media_author),
  media_platform = COALESCE($4, media_platform)
WHERE
  id = $5
RETURNING id, media_type, media_url, media_author, media_platform, created_at
`

type UpdateAnimeMediaParams struct {
	MediaType     pgtype.Text `json:"media_type"`
	MediaUrl      pgtype.Text `json:"media_url"`
	MediaAuthor   pgtype.Text `json:"media_author"`
	MediaPlatform pgtype.Text `json:"media_platform"`
	ID            int64       `json:"id"`
}

func (q *Queries) UpdateAnimeMedia(ctx context.Context, arg UpdateAnimeMediaParams) (AnimeMedium, error) {
	row := q.db.QueryRow(ctx, updateAnimeMedia,
		arg.MediaType,
		arg.MediaUrl,
		arg.MediaAuthor,
		arg.MediaPlatform,
		arg.ID,
	)
	var i AnimeMedium
	err := row.Scan(
		&i.ID,
		&i.MediaType,
		&i.MediaUrl,
		&i.MediaAuthor,
		&i.MediaPlatform,
		&i.CreatedAt,
	)
	return i, err
}