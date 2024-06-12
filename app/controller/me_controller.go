package controller

import (
	"context"
	"github.com/dns2012/dealls-dating-service/app/usecase"
	schemav1 "github.com/dns2012/dealls-dating-service/proto/schema/v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *SchemaV1API) Me(ctx context.Context,  req *emptypb.Empty) (*schemav1.UserResponse, error)  {
	userClaim, err := s.TokenManager.VerifyToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.MeUsecase.Call(ctx, &usecase.MeUsecaseParams{UserID: userClaim.ID})
	if err != nil {
		return nil, err
	}

	user := &schemav1.User{
		Id: uint32(res.User.ID),
		Nickname: res.User.Nickname,
		Email:    res.User.Email,
		RegisteredAt: timestamppb.New(res.User.CreatedAt),
		FullName: res.User.Profile.FullName,
		ImageUrl: res.User.Profile.ImageUrl,
		BirthAt: res.User.Profile.BirthAt.Format("2006-01-02"),
		Gender: schemav1.Gender(res.User.Profile.Gender),
		Company: res.User.Profile.Company,
		JobTitle: res.User.Profile.JobTitle,
	}

	if !res.User.PremiumAt.IsZero() {
		user.IsVerified = true
		user.PremiumAt = timestamppb.New(res.User.PremiumAt)
	}

	return &schemav1.UserResponse{Data: user}, nil
}
