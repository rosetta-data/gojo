// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: anime_season_join_resource.sql

package db

import (
	"context"
)

const getAnimeSeasonResourceDirectly = `-- name: GetAnimeSeasonResourceDirectly :one
SELECT anime_resources.id, anime_resources.tvdb_id, anime_resources.tmdb_id, anime_resources.imdb_id, anime_resources.livechart_id, anime_resources.anime_planet_id, anime_resources.anisearch_id, anime_resources.anidb_id, anime_resources.kitsu_id, anime_resources.mal_id, anime_resources.notify_moe_id, anime_resources.anilist_id, anime_resources.created_at
FROM anime_resources
JOIN anime_season_resources ON anime_resources.id = anime_season_resources.resource_id
WHERE anime_season_resources.season_id = $1
LIMIT 1
`

func (q *Queries) GetAnimeSeasonResourceDirectly(ctx context.Context, seasonID int64) (AnimeResource, error) {
	row := q.db.QueryRow(ctx, getAnimeSeasonResourceDirectly, seasonID)
	var i AnimeResource
	err := row.Scan(
		&i.ID,
		&i.TvdbID,
		&i.TmdbID,
		&i.ImdbID,
		&i.LivechartID,
		&i.AnimePlanetID,
		&i.AnisearchID,
		&i.AnidbID,
		&i.KitsuID,
		&i.MalID,
		&i.NotifyMoeID,
		&i.AnilistID,
		&i.CreatedAt,
	)
	return i, err
}