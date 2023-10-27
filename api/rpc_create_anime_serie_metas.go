package api

import (
	"context"
	"errors"

	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateAnimeSerieMetas(ctx context.Context, req *pb.CreateAnimeSerieMetasRequest) (*pb.CreateAnimeSerieMetasResponse, error) {
	authPayload, err := server.authorizeUser(ctx, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, unAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime Serie metadata")
	}

	if violations := validateCreateAnimeSerieMetasRequest(req); violations != nil {
		return nil, invalidArgumentError(violations)
	}

	var DBAM = make([]db.CreateAnimeSerieMetaTxParam, len(req.AnimeMetas))
	for i, am := range req.AnimeMetas {
		DBAM[i] = db.CreateAnimeSerieMetaTxParam{
			LanguageID: am.GetLanguageID(),
			CreateMetaParams: db.CreateMetaParams{
				Title:    am.GetMeta().GetTitle(),
				Overview: am.GetMeta().GetOverview(),
			},
		}
	}

	arg := db.CreateAnimeSerieMetasTxParams{
		AnimeID:                       req.GetAnimeID(),
		CreateAnimeSerieMetasTxParams: DBAM,
	}

	metas, err := server.gojo.CreateAnimeSerieMetasTx(ctx, arg)
	if err != nil {
		db.ErrorSQL(err)
		return nil, status.Errorf(codes.Internal, "failed to create anime serie metadata : %s", err)
	}

	var PBAM = make([]*pb.AnimeMetaResponse, len(metas.CreateAnimeSerieMetasTxResults))

	for i, am := range metas.CreateAnimeSerieMetasTxResults {
		PBAM[i] = &pb.AnimeMetaResponse{
			Meta:      ConvertMeta(am.Meta),
			Language:  ConvertLanguage(am.Language),
			CreatedAt: timestamppb.New(am.Meta.CreatedAt),
		}
	}

	res := &pb.CreateAnimeSerieMetasResponse{
		AnimeID:    req.GetAnimeID(),
		AnimeMetas: PBAM,
	}
	return res, nil
}

func validateCreateAnimeSerieMetasRequest(req *pb.CreateAnimeSerieMetasRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, fieldViolation("animeID", err))
	}

	if req.AnimeMetas != nil {
		for _, am := range req.AnimeMetas {
			if err := utils.ValidateInt(int64(am.GetLanguageID())); err != nil {
				violations = append(violations, fieldViolation("languageID", err))
			}

			if err := utils.ValidateString(am.GetMeta().GetTitle(), 2, 500); err != nil {
				violations = append(violations, fieldViolation("title", err))
			}

			if err := utils.ValidateString(am.GetMeta().GetOverview(), 5, 5000); err != nil {
				violations = append(violations, fieldViolation("overview", err))
			}
		}

	} else {
		violations = append(violations, fieldViolation("animeMetas", errors.New("give at least one metadata")))
	}

	return violations
}