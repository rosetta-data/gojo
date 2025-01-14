-- name: CreateAnimeSerie :one
INSERT INTO anime_series (
    original_title,
    unique_id,
    first_year,
    last_year,
    mal_id,
    tvdb_id,
    tmdb_id,
    portrait_poster,
    portrait_blur_hash,
    landscape_poster,
    landscape_blur_hash
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING *;

-- name: GetAnimeSerie :one
SELECT * FROM anime_series 
WHERE id = $1 LIMIT 1;

-- name: UpdateAnimeSerie :one
UPDATE anime_series
SET
  original_title = COALESCE(sqlc.narg(original_title), original_title),
  first_year = COALESCE(sqlc.narg(first_year), first_year),
  last_year = COALESCE(sqlc.narg(last_year), last_year),
  mal_id = COALESCE(sqlc.narg(mal_id), mal_id),
  tvdb_id = COALESCE(sqlc.narg(tvdb_id), tvdb_id),
  tmdb_id = COALESCE(sqlc.narg(tmdb_id), tmdb_id),
  portrait_poster = COALESCE(sqlc.narg(portrait_poster), portrait_poster),
  portrait_blur_hash = COALESCE(sqlc.narg(portrait_blur_hash), portrait_blur_hash),
  landscape_poster = COALESCE(sqlc.narg(landscape_poster), landscape_poster),
  landscape_blur_hash = COALESCE(sqlc.narg(landscape_blur_hash), landscape_blur_hash)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeSerie :exec
DELETE FROM anime_series 
WHERE id = $1;

-- name: ListAnimeSeries :many
SELECT * FROM anime_series
WHERE $1 IN (first_year, last_year) OR $1 = 0
LIMIT $2
OFFSET $3;