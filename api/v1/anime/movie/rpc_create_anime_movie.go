package amapiv1

import (
	"context"
	"errors"
	"time"

	av1 "github.com/dj-yacine-flutter/gojo/api/v1/anime"
	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	ampbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ampb"
	nfpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/nfpb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *AnimeMovieServer) CreateAnimeMovie(ctx context.Context, req *ampbv1.CreateAnimeMovieRequest) (*ampbv1.CreateAnimeMovieResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime movie")
	}

	if violations := validateCreateAnimeMovieRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	arg := db.CreateAnimeMovieTxParams{
		CreateAnimeMovieParams: db.CreateAnimeMovieParams{
			OriginalTitle:     req.GetAnimeMovie().GetOriginalTitle(),
			Aired:             req.GetAnimeMovie().GetAired().AsTime(),
			ReleaseYear:       req.GetAnimeMovie().GetReleaseYear(),
			Rating:            req.GetAnimeMovie().GetRating(),
			Duration:          req.GetAnimeMovie().GetDuration().AsDuration(),
			PortraitPoster:    req.GetAnimeMovie().GetPortraitPoster(),
			PortraitBlurHash:  req.GetAnimeMovie().GetPortraitBlurHash(),
			LandscapePoster:   req.GetAnimeMovie().GetLandscapePoster(),
			LandscapeBlurHash: req.GetAnimeMovie().GetLandscapeBlurHash(),
		},
		CreateAnimeResourceParams: db.CreateAnimeResourceParams{
			TvdbID:        req.GetAnimeResources().GetTvdbID(),
			TmdbID:        req.GetAnimeResources().GetTmdbID(),
			ImdbID:        req.GetAnimeResources().GetImdbID(),
			LivechartID:   req.GetAnimeResources().GetLivechartID(),
			AnimePlanetID: req.GetAnimeResources().GetAnimePlanetID(),
			AnisearchID:   req.GetAnimeResources().GetAnisearchID(),
			AnidbID:       req.GetAnimeResources().GetAnidbID(),
			KitsuID:       req.GetAnimeResources().GetKitsuID(),
			MalID:         req.GetAnimeResources().GetMalID(),
			NotifyMoeID:   req.GetAnimeResources().GetNotifyMoeID(),
			AnilistID:     req.GetAnimeResources().GetAnilistID(),
		},
		CreateAnimeLinkParams: db.CreateAnimeLinkParams{
			OfficialWebsite: req.GetAnimeLinks().GetOfficialWebsite(),
			WikipediaUrl:    req.GetAnimeLinks().GetWikipediaUrl(),
			CrunchyrollUrl:  req.GetAnimeLinks().GetCrunchyrollUrl(),
			SocialMedia:     req.GetAnimeLinks().GetSocialMedia(),
		},
	}

	arg.CreateAnimeMetasParams = make([]db.AnimeMetaTxParam, len(req.AnimeMetas))
	for i, v := range req.AnimeMetas {
		arg.CreateAnimeMetasParams[i] = db.AnimeMetaTxParam{
			LanguageID: v.GetLanguageID(),
			CreateMetaParams: db.CreateMetaParams{
				Title:    v.GetMeta().GetTitle(),
				Overview: v.GetMeta().GetOverview(),
			},
		}
	}

	data, err := server.gojo.CreateAnimeMovieTx(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("cannot create anime movie", err)
	}

	res := &ampbv1.CreateAnimeMovieResponse{
		AnimeMovie:     convertAnimeMovie(data.AnimeMovie),
		AnimeResources: av1.ConvertAnimeResource(data.AnimeResource),
		AnimeLinks:     av1.ConvertAnimeLink(data.AnimeLink),
	}

	titles := make([]string, len(data.AnimeMetas))
	res.AnimeMetas = make([]*nfpbv1.AnimeMetaResponse, len(data.AnimeMetas))
	for i, v := range data.AnimeMetas {
		res.AnimeMetas[i] = &nfpbv1.AnimeMetaResponse{
			Meta:       shv1.ConvertMeta(v.Meta),
			LanguageID: v.LanguageID,
			CreatedAt:  timestamppb.New(v.Meta.CreatedAt),
		}
		titles[i] = v.Meta.Title
	}

	server.meilisearch.AddDocuments(&utils.Document{
		ID:     data.AnimeMovie.ID,
		Titles: utils.RemoveDuplicatesTitles(titles),
	})

	return res, nil
}

func validateCreateAnimeMovieRequest(req *ampbv1.CreateAnimeMovieRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if req.AnimeMovie != nil {
		if err := utils.ValidateString(req.GetAnimeMovie().GetOriginalTitle(), 2, 500); err != nil {
			violations = append(violations, shv1.FieldViolation("originalTitle", err))
		}

		if err := utils.ValidateDate(req.GetAnimeMovie().GetAired().AsTime().Format(time.DateOnly)); err != nil {
			violations = append(violations, shv1.FieldViolation("aired", err))
		}

		if err := utils.ValidateYear(req.GetAnimeMovie().GetReleaseYear()); err != nil {
			violations = append(violations, shv1.FieldViolation("releaseYear", err))
		}

		if err := utils.ValidateString(req.GetAnimeMovie().GetRating(), 2, 30); err != nil {
			violations = append(violations, shv1.FieldViolation("rating", err))
		}

		if err := utils.ValidateDuration(req.GetAnimeMovie().GetDuration().AsDuration().String()); err != nil {
			violations = append(violations, shv1.FieldViolation("duration", err))
		}

		if err := utils.ValidateImage(req.GetAnimeMovie().GetPortraitPoster()); err != nil {
			violations = append(violations, shv1.FieldViolation("portraitPoster", err))
		}

		if err := utils.ValidateString(req.GetAnimeMovie().GetPortraitBlurHash(), 0, 100); err != nil {
			violations = append(violations, shv1.FieldViolation("portraitBlurHash", err))
		}

		if err := utils.ValidateImage(req.GetAnimeMovie().GetLandscapePoster()); err != nil {
			violations = append(violations, shv1.FieldViolation("landscapePoster", err))
		}

		if err := utils.ValidateString(req.GetAnimeMovie().GetLandscapeBlurHash(), 0, 100); err != nil {
			violations = append(violations, shv1.FieldViolation("landscapeBlurHash", err))
		}

	} else {
		violations = append(violations, shv1.FieldViolation("animeMovie", errors.New("you need to send the animeMovie model")))
	}

	if req.AnimeResources != nil {
		if err := utils.ValidateInt(int64(req.GetAnimeResources().GetTvdbID())); err != nil {
			violations = append(violations, shv1.FieldViolation("tvdbID", err))
		}

		if err := utils.ValidateInt(int64(req.GetAnimeResources().GetTmdbID())); err != nil {
			violations = append(violations, shv1.FieldViolation("tmdbID", err))
		}

		if err := utils.ValidateInt(int64(req.GetAnimeResources().GetLivechartID())); err != nil {
			violations = append(violations, shv1.FieldViolation("livechartID", err))
		}

		if err := utils.ValidateInt(int64(req.GetAnimeResources().GetAnidbID())); err != nil {
			violations = append(violations, shv1.FieldViolation("anidbID", err))
		}

		if err := utils.ValidateInt(int64(req.GetAnimeResources().GetAnisearchID())); err != nil {
			violations = append(violations, shv1.FieldViolation("anisearchID", err))
		}

		if err := utils.ValidateInt(int64(req.GetAnimeResources().GetKitsuID())); err != nil {
			violations = append(violations, shv1.FieldViolation("kitsuID", err))
		}

		if err := utils.ValidateInt(int64(req.GetAnimeResources().GetMalID())); err != nil {
			violations = append(violations, shv1.FieldViolation("malID", err))
		}

		if err := utils.ValidateInt(int64(req.GetAnimeResources().GetAnilistID())); err != nil {
			violations = append(violations, shv1.FieldViolation("anilistID", err))
		}

	} else {
		violations = append(violations, shv1.FieldViolation("animeResources", errors.New("you need to send the AnimeResources model")))
	}

	if req.AnimeLinks != nil {
		if err := utils.ValidateURL(req.GetAnimeLinks().GetOfficialWebsite(), ""); err != nil {
			violations = append(violations, shv1.FieldViolation("officialWebsite", err))
		}

		if err := utils.ValidateURL(req.GetAnimeLinks().GetCrunchyrollUrl(), "crunchyroll"); err != nil {
			violations = append(violations, shv1.FieldViolation("crunchyrollUrl", err))
		}

		if err := utils.ValidateURL(req.GetAnimeLinks().GetWikipediaUrl(), "wikipedia"); err != nil {
			violations = append(violations, shv1.FieldViolation("wikipediaUrl", err))
		}

		if len(req.GetAnimeLinks().GetSocialMedia()) > 0 {
			for _, v := range req.GetAnimeLinks().GetSocialMedia() {
				if err := utils.ValidateURL(v, ""); err != nil {
					violations = append(violations, shv1.FieldViolation("socialMedia", err))
				}
			}
		}

	} else {
		violations = append(violations, shv1.FieldViolation("animeLinks", errors.New("you need to send the AnimeLinks model")))
	}

	if req.AnimeMetas != nil {
		for _, v := range req.AnimeMetas {
			if err := utils.ValidateInt(int64(v.GetLanguageID())); err != nil {
				violations = append(violations, shv1.FieldViolation("languageID", err))
			}

			if err := utils.ValidateString(v.GetMeta().GetTitle(), 2, 500); err != nil {
				violations = append(violations, shv1.FieldViolation("title", err))
			}

			if err := utils.ValidateString(v.GetMeta().GetOverview(), 5, 5000); err != nil {
				violations = append(violations, shv1.FieldViolation("overview", err))
			}
		}

	} else {
		violations = append(violations, shv1.FieldViolation("animeMetas", errors.New("give at least one metadata")))
	}

	return violations
}
