// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: anime_season_official_title.sql

package db

import (
	"context"
)

const createAnimeSeasonOfficialTitle = `-- name: CreateAnimeSeasonOfficialTitle :one
INSERT INTO anime_season_official_titles (season_id, title_text)
VALUES ($1, $2)
ON CONFLICT (season_id, title_text)
DO UPDATE SET 
    season_id = excluded.season_id
RETURNING id, season_id, title_text, created_at
`

type CreateAnimeSeasonOfficialTitleParams struct {
	SeasonID  int64
	TitleText string
}

func (q *Queries) CreateAnimeSeasonOfficialTitle(ctx context.Context, arg CreateAnimeSeasonOfficialTitleParams) (AnimeSeasonOfficialTitle, error) {
	row := q.db.QueryRow(ctx, createAnimeSeasonOfficialTitle, arg.SeasonID, arg.TitleText)
	var i AnimeSeasonOfficialTitle
	err := row.Scan(
		&i.ID,
		&i.SeasonID,
		&i.TitleText,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAnimeSeasonOfficialTitle = `-- name: DeleteAnimeSeasonOfficialTitle :exec
DELETE FROM anime_season_official_titles
WHERE id = $1
`

func (q *Queries) DeleteAnimeSeasonOfficialTitle(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteAnimeSeasonOfficialTitle, id)
	return err
}

const getAnimeSeasonOfficialTitles = `-- name: GetAnimeSeasonOfficialTitles :many
SELECT id, season_id, title_text, created_at FROM anime_season_official_titles
WHERE season_id = $1
`

func (q *Queries) GetAnimeSeasonOfficialTitles(ctx context.Context, seasonID int64) ([]AnimeSeasonOfficialTitle, error) {
	rows, err := q.db.Query(ctx, getAnimeSeasonOfficialTitles, seasonID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AnimeSeasonOfficialTitle{}
	for rows.Next() {
		var i AnimeSeasonOfficialTitle
		if err := rows.Scan(
			&i.ID,
			&i.SeasonID,
			&i.TitleText,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const queryAnimeSeasonOfficialTitles = `-- name: QueryAnimeSeasonOfficialTitles :many
WITH search_documents AS (
  SELECT
    season_id,
    title_text,
    to_tsvector('pg_catalog.english', title_text) AS title_text_tsv
  FROM anime_season_official_titles
)
SELECT season_id
FROM search_documents
WHERE (
  $1::varchar IS NOT NULL AND $1::varchar <> '' AND
  (
    SELECT bool_and(
      to_tsvector('pg_catalog.english', lower(translate(title_text, '[:punct:]', ''))) 
      @@ plainto_tsquery(word)
    )
    FROM unnest(regexp_split_to_array($1::varchar, '\\W+')) AS word
  )
  OR title_text ILIKE '%' || $1::varchar || '%'
)
ORDER BY
  ts_rank(title_text_tsv, phraseto_tsquery($1::varchar)) DESC,
  similarity(title_text, $1::varchar) DESC
  LIMIT $2
  OFFSET $3
`

type QueryAnimeSeasonOfficialTitlesParams struct {
	Column1 string
	Limit   int32
	Offset  int32
}

func (q *Queries) QueryAnimeSeasonOfficialTitles(ctx context.Context, arg QueryAnimeSeasonOfficialTitlesParams) ([]int64, error) {
	rows, err := q.db.Query(ctx, queryAnimeSeasonOfficialTitles, arg.Column1, arg.Limit, arg.Offset)
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
