package controller

import (
	"context"
	"github.com/dns2012/dealls-dating-service/app/domain/exception"
	"github.com/dns2012/dealls-dating-service/app/usecase"
	schemav1 "github.com/dns2012/dealls-dating-service/proto/schema/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *SchemaV1API) CreateUserPreference(ctx context.Context, req *schemav1.CreateUserPreferenceRequest) (*emptypb.Empty, error)  {
	if err := req.Validate(); err != nil {
		return nil, exception.Invalid(err.Error())
	}

	userClaim, err := s.TokenManager.VerifyToken(ctx)
	if err != nil {
		return nil, err
	}

	if userClaim.ID == uint(req.GetPreferenceUserId()) {
		return nil, exception.Invalid("Invalid user preference. Please use another user preference.")
	}

	err = s.CreateUserPreferenceUsecase.Call(ctx, &usecase.CreateUserPreferenceUsecaseParams{
		UserID:           userClaim.ID,
		PreferenceUserID: uint(req.GetPreferenceUserId()),
		PreferenceType:   uint(req.GetPreferenceType()),
	})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
