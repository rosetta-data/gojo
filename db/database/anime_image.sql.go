// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: anime_image.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAnimeImage = `-- name: CreateAnimeImage :one
INSERT INTO anime_images (image_type, image_host, image_url, image_thumbnails, image_blurhash, image_height, image_width)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING  id, image_type, image_host, image_url, image_thumbnails, image_blurhash, image_height, image_width, created_at
`

type CreateAnimeImageParams struct {
	ImageType       int32
	ImageHost       string
	ImageUrl        string
	ImageThumbnails string
	ImageBlurhash   string
	ImageHeight     int32
	ImageWidth      int32
}

func (q *Queries) CreateAnimeImage(ctx context.Context, arg CreateAnimeImageParams) (AnimeImage, error) {
	row := q.db.QueryRow(ctx, createAnimeImage,
		arg.ImageType,
		arg.ImageHost,
		arg.ImageUrl,
		arg.ImageThumbnails,
		arg.ImageBlurhash,
		arg.ImageHeight,
		arg.ImageWidth,
	)
	var i AnimeImage
	err := row.Scan(
		&i.ID,
		&i.ImageType,
		&i.ImageHost,
		&i.ImageUrl,
		&i.ImageThumbnails,
		&i.ImageBlurhash,
		&i.ImageHeight,
		&i.ImageWidth,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAnimeImage = `-- name: DeleteAnimeImage :exec
DELETE FROM anime_images
WHERE id = $1
`

func (q *Queries) DeleteAnimeImage(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteAnimeImage, id)
	return err
}

const getAnimeImage = `-- name: GetAnimeImage :one
SELECT id, image_type, image_host, image_url, image_thumbnails, image_blurhash, image_height, image_width, created_at FROM anime_images
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAnimeImage(ctx context.Context, id int64) (AnimeImage, error) {
	row := q.db.QueryRow(ctx, getAnimeImage, id)
	var i AnimeImage
	err := row.Scan(
		&i.ID,
		&i.ImageType,
		&i.ImageHost,
		&i.ImageUrl,
		&i.ImageThumbnails,
		&i.ImageBlurhash,
		&i.ImageHeight,
		&i.ImageWidth,
		&i.CreatedAt,
	)
	return i, err
}

const updateAnimeImage = `-- name: UpdateAnimeImage :one
UPDATE anime_images
SET
  image_type = COALESCE($1, image_type),
  image_host = COALESCE($2, image_host),
  image_url = COALESCE($3, image_url),
  image_thumbnails = COALESCE($4, image_thumbnails),
  image_blurhash = COALESCE($5, image_blurhash),
  image_height = COALESCE($6, image_height),
  image_width = COALESCE($7, image_width)
WHERE
  id = $8
RETURNING id, image_type, image_host, image_url, image_thumbnails, image_blurhash, image_height, image_width, created_at
`

type UpdateAnimeImageParams struct {
	ImageType       pgtype.Int4
	ImageHost       pgtype.Text
	ImageUrl        pgtype.Text
	ImageThumbnails pgtype.Text
	ImageBlurhash   pgtype.Text
	ImageHeight     pgtype.Int4
	ImageWidth      pgtype.Int4
	ID              int64
}

func (q *Queries) UpdateAnimeImage(ctx context.Context, arg UpdateAnimeImageParams) (AnimeImage, error) {
	row := q.db.QueryRow(ctx, updateAnimeImage,
		arg.ImageType,
		arg.ImageHost,
		arg.ImageUrl,
		arg.ImageThumbnails,
		arg.ImageBlurhash,
		arg.ImageHeight,
		arg.ImageWidth,
		arg.ID,
	)
	var i AnimeImage
	err := row.Scan(
		&i.ID,
		&i.ImageType,
		&i.ImageHost,
		&i.ImageUrl,
		&i.ImageThumbnails,
		&i.ImageBlurhash,
		&i.ImageHeight,
		&i.ImageWidth,
		&i.CreatedAt,
	)
	return i, err
}