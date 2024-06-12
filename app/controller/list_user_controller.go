package controller

import (
	"context"
	"github.com/dns2012/dealls-dating-service/app/domain/exception"
	"github.com/dns2012/dealls-dating-service/app/usecase"
	schemav1 "github.com/dns2012/dealls-dating-service/proto/schema/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *SchemaV1API) ListUser(ctx context.Context, req *schemav1.ListUserRequest) (*schemav1.ListUserResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.Invalid(err.Error())
	}

	userClaim, err := s.TokenManager.VerifyToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.ListUserUsecase.Call(ctx, &usecase.ListUserUsecaseParams{
		UserID:   userClaim.ID,
		Page:     uint(req.GetPage()),
		PageSize: uint(req.GetPageSize()),
	})
	if err != nil {
		return nil, err
	}

	var users []*schemav1.User

	for _, i := range res.List {
		user := &schemav1.User{
			Id: uint32(i.ID),
			Nickname: i.Nickname,
			Email:    i.Email,
			RegisteredAt: timestamppb.New(i.CreatedAt),
			FullName: i.Profile.FullName,
			ImageUrl: i.Profile.ImageUrl,
			BirthAt: i.Profile.BirthAt.Format("2006-01-02"),
			Gender: schemav1.Gender(i.Profile.Gender),
			Company: i.Profile.Company,
			JobTitle: i.Profile.JobTitle,
		}

		if !i.PremiumAt.IsZero() {
			user.IsVerified = true
			user.PremiumAt = timestamppb.New(i.PremiumAt)
		}
		users = append(users, user)
	}

	return &schemav1.ListUserResponse{
		Page: &schemav1.Page{
			Current: uint32(res.Page.Page),
			Size:      uint32(res.Page.PageSize),
			Next:      uint32(res.Page.Next),
			Prev:      uint32(res.Page.Prev),
			Count:     uint32(res.Page.PageCount),
			RowsCount: uint32(res.Page.RowsCount),
		},
		Data: users,
	}, nil
}
