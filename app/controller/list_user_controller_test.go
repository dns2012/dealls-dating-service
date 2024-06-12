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
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestSchemaV1API_ListUser(t *testing.T) {
	type input struct {
		ctx context.Context
		req *schemav1.ListUserRequest
	}

	type output struct {
		res *schemav1.ListUserResponse
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
			name: "BAD_REQUEST",
			in: &input{
				ctx: context.Background(),
				req: &schemav1.ListUserRequest{},
			},
			out: &output{},
			on: func(m *SchemaV1APIMock, i *input, o *output) {
				o.err = exception.Invalid(i.req.Validate().Error())
			},
		},
		{
			name: "UNAUTHENTICATED",
			in: &input{
				ctx: context.Background(),
				req: &schemav1.ListUserRequest{
					Page: 1,
					PageSize: 10,
				},
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
				req: &schemav1.ListUserRequest{
					Page: 1,
					PageSize: 10,
				},
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
				m.listUserUsecase.On("Call", i.ctx, &usecase.ListUserUsecaseParams{
					UserID:   1,
					Page:     uint(i.req.GetPage()),
					PageSize: uint(i.req.GetPageSize()),
				}).Return(nil, o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				req: &schemav1.ListUserRequest{
					Page: 1,
					PageSize: 10,
				},
			},
			out: &output{
				res: &schemav1.ListUserResponse{
					Page: &schemav1.Page{
						Current:   1,
						Size:      10,
						Next:      2,
						Prev:      1,
						Count:     2,
						RowsCount: 20,
					},
					Data: []*schemav1.User{
						{
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
			},
			on: func(m *SchemaV1APIMock, i *input, o *output) {
				m.tokenManager.On("VerifyToken", i.ctx).Return(&manager.UserClaims{
					ID:               1,
					Nickname:         "User",
					Email:            "user@dealls.com",
					RegisteredClaims: jwt.RegisteredClaims{},
				}, nil)
				birthDate, _ := time.Parse("2006-01-02", o.res.GetData()[0].GetBirthAt())
				m.listUserUsecase.On("Call", i.ctx, &usecase.ListUserUsecaseParams{
					UserID:   1,
					Page:     uint(i.req.GetPage()),
					PageSize: uint(i.req.GetPageSize()),
				}).Return(&usecase.ListUserUsecaseResult{
					Page: &manager.Pagination{
						Page:      int(o.res.GetPage().GetCurrent()),
						PageSize:  int(o.res.GetPage().GetSize()),
						PageCount: int(o.res.GetPage().GetCount()),
						Next:      int(o.res.GetPage().GetNext()),
						Prev:      int(o.res.GetPage().GetPrev()),
						RowsCount: int64(o.res.GetPage().GetRowsCount()),
					},
					List: []*entity.User{
						{
							Model: gorm.Model{
								ID: uint(o.res.GetData()[0].GetId()),
								CreatedAt: o.res.GetData()[0].GetRegisteredAt().AsTime(),
							},
							PremiumAt:   o.res.GetData()[0].GetPremiumAt().AsTime(),
							Nickname:    o.res.GetData()[0].GetNickname(),
							Email:       o.res.GetData()[0].GetEmail(),
							Profile:     entity.Profile{
								UserID:   uint(o.res.GetData()[0].GetId()),
								FullName: o.res.GetData()[0].GetFullName(),
								ImageUrl: o.res.GetData()[0].GetImageUrl(),
								BirthAt:  birthDate,
								Gender: uint(o.res.GetData()[0].GetGender()),
								Company:  o.res.GetData()[0].GetCompany(),
								JobTitle: o.res.GetData()[0].GetJobTitle(),
							},
						},
					},
				}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SchemaV1APIMock{
				listUserUsecase: mocks.ListUserUsecase{},
				tokenManager: mocks.TokenManager{},
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := new(controller.SchemaV1API)
			subject.ListUserUsecase = &m.listUserUsecase
			subject.TokenManager = &m.tokenManager
			res, err := subject.ListUser(tt.in.ctx, tt.in.req)

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
