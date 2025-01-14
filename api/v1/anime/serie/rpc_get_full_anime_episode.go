package asapiv1

import (
	"context"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	apbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/apb"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	nfpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/nfpb"
	"github.com/dj-yacine-flutter/gojo/ping"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgerrcode"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *AnimeSerieServer) GetFullAnimeEpisode(ctx context.Context, req *aspbv1.GetFullAnimeEpisodeRequest) (*aspbv1.GetFullAnimeEpisodeResponse, error) {
	var err error

	_, err = shv1.AuthorizeUser(ctx, server.tokenMaker, utils.AllRolls)
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	violations := validateGetFullAnimeEpisodeRequest(req)
	if violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	cache := &ping.CacheKey{
		ID:     req.EpisodeID,
		Target: ping.AnimeEpisode,
	}

	res := &aspbv1.GetFullAnimeEpisodeResponse{}

	var episode db.AnimeEpisode
	if err = server.ping.Handle(ctx, cache.Main(), &episode, func() error {
		episode, err = server.gojo.GetAnimeEpisode(ctx, req.GetEpisodeID())
		if err != nil {
			return shv1.ApiError("cannot get anime episode", err)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	res.AnimeEpisode = convertAnimeEpisode(episode)

	var meta db.Meta
	if err = server.ping.Handle(ctx, cache.Meta(), &meta, func() error {
		meta, err = server.gojo.GetAnimeEpisodeMetaWithLanguageDirectly(ctx, db.GetAnimeEpisodeMetaWithLanguageDirectlyParams{
			EpisodeID:  req.GetEpisodeID(),
			LanguageID: req.GetLanguageID(),
		})
		if err != nil {
			return shv1.ApiError("no anime episode found with this language ID", err)
		}

		return nil
	}); err != nil {
		return nil, err
	}

	res.EpisodeMeta = &nfpbv1.AnimeMetaResponse{
		LanguageID: req.GetLanguageID(),
		Meta:       shv1.ConvertMeta(meta),
		CreatedAt:  timestamppb.New(meta.CreatedAt),
	}

	var serverID int64
	if err = server.ping.Handle(ctx, cache.Server(), &serverID, func() error {
		serverID, err = server.gojo.GetAnimeEpisodeServerByEpisodeID(ctx, req.GetEpisodeID())
		if err != nil {
			if db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot get anime episode server ID", err)
			} else {
				return nil
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	res.ServerID = serverID

	if serverID != 0 {
		var subVideos []db.AnimeEpisodeVideo
		if err = server.ping.Handle(ctx, cache.SubV(), &subVideos, func() error {
			v, err := server.gojo.ListAnimeEpisodeServerSubVideos(ctx, res.ServerID)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot list anime episode server sub videos", err)
			}

			if len(v) > 0 {
				subVideos = make([]db.AnimeEpisodeVideo, len(v))
				for i, x := range v {
					subVideos[i], err = server.gojo.GetAnimeEpisodeVideo(ctx, x.VideoID)
					if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
						return shv1.ApiError("cannot get anime episode server sub videos", err)
					}
				}
			}

			return nil
		}); err != nil {
			return nil, err
		}

		var subTorrents []db.AnimeEpisodeTorrent
		if err = server.ping.Handle(ctx, cache.SubT(), &subTorrents, func() error {
			v, err := server.gojo.ListAnimeEpisodeServerSubTorrents(ctx, res.ServerID)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot list anime episode server sub torrents", err)
			}

			if len(v) > 0 {
				subTorrents = make([]db.AnimeEpisodeTorrent, len(v))
				for i, x := range v {
					subTorrents[i], err = server.gojo.GetAnimeEpisodeTorrent(ctx, x.TorrentID)
					if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
						return shv1.ApiError("cannot get anime episode server sub torrents", err)
					}
				}

			}

			return nil
		}); err != nil {
			return nil, err
		}

		res.Sub = &apbv1.AnimeSubDataResponse{
			Videos:   convertAnimeEpisodeVideos(subVideos),
			Torrents: convertAnimeEpisodeTorrents(subTorrents),
		}

		var dubVideos []db.AnimeEpisodeVideo
		if err = server.ping.Handle(ctx, cache.DubV(), &dubVideos, func() error {
			v, err := server.gojo.ListAnimeEpisodeServerDubVideos(ctx, res.ServerID)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot list anime episode server dub videos", err)
			}

			if len(v) > 0 {
				dubVideos = make([]db.AnimeEpisodeVideo, len(v))
				for i, x := range v {
					dubVideos[i], err = server.gojo.GetAnimeEpisodeVideo(ctx, x.VideoID)
					if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
						return shv1.ApiError("cannot get anime episode server dub videos", err)
					}
				}
			}

			return nil
		}); err != nil {
			return nil, err
		}

		var dubTorrents []db.AnimeEpisodeTorrent
		if err = server.ping.Handle(ctx, cache.DubT(), &dubTorrents, func() error {
			v, err := server.gojo.ListAnimeEpisodeServerDubTorrents(ctx, res.ServerID)
			if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
				return shv1.ApiError("cannot list anime episode server dub torrents", err)
			}

			if len(v) > 0 {
				dubTorrents = make([]db.AnimeEpisodeTorrent, len(v))
				for i, x := range v {
					dubTorrents[i], err = server.gojo.GetAnimeEpisodeTorrent(ctx, x.TorrentID)
					if err != nil && db.ErrorDB(err).Code != pgerrcode.CaseNotFound {
						return shv1.ApiError("cannot get anime episode server dub torrents", err)
					}
				}
			}

			return nil
		}); err != nil {
			return nil, err
		}

		res.Dub = &apbv1.AnimeDubDataResponse{
			Videos:   convertAnimeEpisodeVideos(dubVideos),
			Torrents: convertAnimeEpisodeTorrents(dubTorrents),
		}
	}

	return res, nil
}

func validateGetFullAnimeEpisodeRequest(req *aspbv1.GetFullAnimeEpisodeRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetEpisodeID()); err != nil {
		violations = append(violations, shv1.FieldViolation("episodeID", err))
	}

	if err := utils.ValidateInt(int64(req.GetLanguageID())); err != nil {
		violations = append(violations, shv1.FieldViolation("languageID", err))
	}

	return violations
}
