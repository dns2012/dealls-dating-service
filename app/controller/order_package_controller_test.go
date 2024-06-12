package controller_test

import (
	"context"
	"github.com/dns2012/dealls-dating-service/app/controller"
	"github.com/dns2012/dealls-dating-service/app/domain/exception"
	"github.com/dns2012/dealls-dating-service/app/domain/manager"
	"github.com/dns2012/dealls-dating-service/app/usecase"
	"github.com/dns2012/dealls-dating-service/mocks"
	schemav1 "github.com/dns2012/dealls-dating-service/proto/schema/v1"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"
)

func TestSchemaV1API_OrderPackage(t *testing.T) {
	type input struct {
		ctx context.Context
		req *schemav1.OrderPackageRequest
	}

	type output struct {
		res *emptypb.Empty
		err error
	}

	type testcase struct {
		name string
		in   *input
		out  *output
		on   func(*SchemaV1APIMock, *input, *output)
	}

	tests := []testcase{
		{
			name: "BAD_REQUEST",
			in: &input{
				ctx: context.Background(),
				req: &schemav1.OrderPackageRequest{},
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
				req: &schemav1.OrderPackageRequest{
					Id: 1,
					TotalPayment: 10000,
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
				req: &schemav1.OrderPackageRequest{
					Id: 1,
					TotalPayment: 10000,
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
				m.orderPackageUsecase.On("Call", i.ctx, &usecase.OrderPackageUsecaseParams{
					UserID:       1,
					PackageID:    uint(i.req.GetId()),
					TotalPayment: uint(i.req.GetTotalPayment()),
				}).Return(o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				req: &schemav1.OrderPackageRequest{
					Id: 1,
					TotalPayment: 10000,
				},
			},
			out: &output{
				res: &emptypb.Empty{},
			},
			on: func(m *SchemaV1APIMock, i *input, o *output) {
				m.tokenManager.On("VerifyToken", i.ctx).Return(&manager.UserClaims{
					ID:               1,
					Nickname:         "User",
					Email:            "user@dealls.com",
					RegisteredClaims: jwt.RegisteredClaims{},
				}, nil)
				m.orderPackageUsecase.On("Call", i.ctx, &usecase.OrderPackageUsecaseParams{
					UserID:       1,
					PackageID:    uint(i.req.GetId()),
					TotalPayment: uint(i.req.GetTotalPayment()),
				}).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SchemaV1APIMock{
				orderPackageUsecase: mocks.OrderPackageUsecase{},
				tokenManager: mocks.TokenManager{},
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := new(controller.SchemaV1API)
			subject.OrderPackageUsecase = &m.orderPackageUsecase
			subject.TokenManager = &m.tokenManager
			res, err := subject.OrderPackage(tt.in.ctx, tt.in.req)

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
