package controller_test

import (
	"context"
	"github.com/dns2012/dealls-dating-service/app/controller"
	"github.com/dns2012/dealls-dating-service/app/domain/entity"
	"github.com/dns2012/dealls-dating-service/app/domain/exception"
	"github.com/dns2012/dealls-dating-service/app/domain/manager"
	"github.com/dns2012/dealls-dating-service/app/usecase"
	"github.com/dns2012/dealls-dating-service/mocks"
	schemav1 "github.com/dns2012/dealls-dating-service/proto/schema/v1"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestSchemaV1API_Me(t *testing.T) {
	type input struct {
		ctx context.Context
		req *emptypb.Empty
	}

	type output struct {
		res *schemav1.UserResponse
		err error
	}

	type testcase struct {
		name string
		in   *input
		out  *output
		on   func(*SchemaV1APIMock, *input, *output)
	}

	grpcTimeNow := timestamppb.Now()
	tests := []testcase{
		{
			name: "UNAUTHENTICATED",
			in: &input{
				ctx: context.Background(),
				req: &emptypb.Empty{},
			},
			out: &output{
				err: exception.Unauthenticated("UNAUTHENTICATED"),
			},
			on: func(m *SchemaV1APIMock, i *input, o *output) {
				m.tokenManager.On("VerifyToken", i.ctx).Return(nil, o.err)
			},
		},
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx: context.Background(),
				req: &emptypb.Empty{},
			},
			out: &output{
				err: exception.Internal("INTERNAL_SERVER_ERROR"),
			},
			on: func(m *SchemaV1APIMock, i *input, o *output) {
				m.tokenManager.On("VerifyToken", i.ctx).Return(&manager.UserClaims{
					ID:               1,
					Nickname:         "User",
					Email:            "user@dealls.com",
					RegisteredClaims: jwt.RegisteredClaims{},
				}, nil)
				m.meUsecase.On("Call", i.ctx, &usecase.MeUsecaseParams{UserID: 1}).Return(nil, o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				req: &emptypb.Empty{},
			},
			out: &output{
				res: &schemav1.UserResponse{
					Data: &schemav1.User{
						Id: 1,
						Nickname:     "User",
						Email:        "user@dealls.com",
						IsVerified:   true,
						PremiumAt:    grpcTimeNow,
						RegisteredAt: grpcTimeNow,
						FullName:     "User Dealls",
						ImageUrl:     "https://image.com",
						BirthAt:      "1999-09-09",
						Gender:       schemav1.Gender_MALE,
						Company:      "Dealls",
						JobTitle:     "Backend Engineer",
					},
				},
			},
			on: func(m *SchemaV1APIMock, i *input, o *output) {
				m.tokenManager.On("VerifyToken", i.ctx).Return(&manager.UserClaims{
					ID:               1,
					Nickname:         "User",
					Email:            "user@dealls.com",
					RegisteredClaims: jwt.RegisteredClaims{},
				}, nil)

				birthDate, _ := time.Parse("2006-01-02", o.res.GetData().GetBirthAt())
				m.meUsecase.On("Call", i.ctx, &usecase.MeUsecaseParams{UserID: 1}).Return(&usecase.MeUsecaseResult{
					User: &entity.User{
						Model: gorm.Model{
							ID: uint(o.res.GetData().GetId()),
							CreatedAt: o.res.GetData().GetRegisteredAt().AsTime(),
						},
						PremiumAt:   o.res.GetData().GetPremiumAt().AsTime(),
						Nickname:    o.res.GetData().GetNickname(),
						Email:       o.res.GetData().GetEmail(),
						Profile:     entity.Profile{
							UserID:   uint(o.res.GetData().GetId()),
							FullName: o.res.GetData().GetFullName(),
							ImageUrl: o.res.GetData().GetImageUrl(),
							BirthAt:  birthDate,
							Gender: uint(o.res.GetData().GetGender()),
							Company:  o.res.GetData().GetCompany(),
							JobTitle: o.res.GetData().GetJobTitle(),
						},
					},
				}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SchemaV1APIMock{
				meUsecase: mocks.MeUsecase{},
				tokenManager: mocks.TokenManager{},
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := new(controller.SchemaV1API)
			subject.MeUsecase = &m.meUsecase
			subject.TokenManager = &m.tokenManager
			res, err := subject.Me(tt.in.ctx, tt.in.req)

			if tt.out.err != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.out.err, err)
			}

			if tt.out.res != nil {
				assert.NotNil(t, tt.out.res)
				assert.Equal(t, tt.out.res, res)
			}
		})
	}
}
