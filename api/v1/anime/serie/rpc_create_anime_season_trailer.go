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

func (server *AnimeSerieServer) CreateAnimeSeasonTrailer(ctx context.Context, req *aspbv1.CreateAnimeSeasonTrailerRequest) (*aspbv1.CreateAnimeSeasonTrailerResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime season trailer")
	}

	if violations := validateCreateAnimeSeasonTrailerRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	arg := db.CreateAnimeSeasonTrailerTxParams{
		SeasonID: req.GetSeasonID(),
	}

	if req.SeasonTrailers != nil {
		arg.SeasonTrailersParams = make([]db.CreateAnimeTrailerParams, len(req.GetSeasonTrailers()))
		for i, v := range req.GetSeasonTrailers() {
			arg.SeasonTrailersParams[i].IsOfficial = v.IsOfficial
			arg.SeasonTrailersParams[i].HostName = v.HostName
			arg.SeasonTrailersParams[i].HostKey = v.HostKey
		}
	}

	data, err := server.gojo.CreateAnimeSeasonTrailerTx(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to create anime season trailer", err)
	}

	res := &aspbv1.CreateAnimeSeasonTrailerResponse{
		AnimeSeason:    convertAnimeSeason(data.AnimeSeason),
		SeasonTrailers: av1.ConvertAnimeTrailers(data.SeasonTrailers),
	}
	return res, nil
}

func validateCreateAnimeSeasonTrailerRequest(req *aspbv1.CreateAnimeSeasonTrailerRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetSeasonID()); err != nil {
		violations = append(violations, shv1.FieldViolation("seasonID", err))
	}

	if req.SeasonTrailers != nil {
		if len(req.GetSeasonTrailers()) > 0 {
			for i, v := range req.GetSeasonTrailers() {
				if err := utils.ValidateString(v.HostName, 1, 200); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("SeasonTrailers > hostName at index [%d]", i), err))
				}
				if err := utils.ValidateString(v.HostKey, 1, 200); err != nil {
					violations = append(violations, shv1.FieldViolation(fmt.Sprintf("SeasonTrailers > hostKey at index [%d]", i), err))
				}
			}
		}

	} else {
		violations = append(violations, shv1.FieldViolation("SeasonTrailers", errors.New("you need to send the SeasonTrailers model")))
	}

	return violations
}
