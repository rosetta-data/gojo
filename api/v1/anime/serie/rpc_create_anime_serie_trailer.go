package asapiv1

import (
	"context"
	"errors"
	"fmt"

	av1 "github.com/dj-yacine-flutter/gojo/api/v1/anime"
	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeSerieServer) CreateAnimeSerieTrailer(ctx context.Context, req *aspbv1.CreateAnimeSerieTrailerRequest) (*aspbv1.CreateAnimeSerieTrailerResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime Serie trailer")
	}

	if violations := validateCreateAnimeSerieTrailerRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	arg := db.CreateAnimeSerieTrailerTxParams{
		AnimeID: req.GetAnimeID(),
	}

	if req.AnimeTrailers != nil {
		arg.AnimeTrailersParams = make([]db.CreateAnimeTrailerParams, len(req.GetAnimeTrailers()))
		for i, v := range req.GetAnimeTrailers() {
			arg.AnimeTrailersParams[i].IsOfficial = v.IsOfficial
			arg.AnimeTrailersParams[i].HostName = v.HostName
			arg.AnimeTrailersParams[i].HostKey = v.HostKey
		}
	}

	data, err := server.gojo.CreateAnimeSerieTrailerTx(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to create anime Serie trailer", err)
	}

	res := &aspbv1.CreateAnimeSerieTrailerResponse{
		AnimeSerie:    convertAnimeSerie(data.AnimeSerie),
		AnimeTrailers: av1.ConvertAnimeTrailers(data.AnimeTrailers),
	}
	return res, nil
}

func validateCreateAnimeSerieTrailerRequest(req *aspbv1.CreateAnimeSerieTrailerRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shv1.FieldViolation("ID", err))
	}

	if req.AnimeTrailers != nil {
		if len(req.GetAnimeTrailers()) > 0 {
			for i, v := range req.GetAnimeTrailers() {
				if err := utils.ValidateString(v.HostName, 1, 200); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeTrailers > hostName at index [%d]", i), err))
				}
				if err := utils.ValidateString(v.HostKey, 1, 200); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("animeTrailers > hostKey at index [%d]", i), err))
				}
			}
		}

	} else {
		violations = append(violations, shv1.FieldViolation("animeTrailers", errors.New("you need to send the AnimeTrailers model")))
	}

	return violations
}
