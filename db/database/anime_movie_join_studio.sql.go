// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: anime_movie_join_studio.sql

package db

import (
	"context"
)

const getAnimeMovieStudiosDirectly = `-- name: GetAnimeMovieStudiosDirectly :many
SELECT studios.id, studios.studio_name, studios.created_at
FROM studios
JOIN anime_movie_studios ON studios.id = anime_movie_studios.studio_id
WHERE anime_movie_studios.anime_id = $1
`

func (q *Queries) GetAnimeMovieStudiosDirectly(ctx context.Context, animeID int64) ([]Studio, error) {
	rows, err := q.db.Query(ctx, getAnimeMovieStudiosDirectly, animeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Studio{}
	for rows.Next() {
		var i Studio
		if err := rows.Scan(&i.ID, &i.StudioName, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
