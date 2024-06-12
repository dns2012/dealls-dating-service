package controller_test

import (
	"context"
	"github.com/dns2012/dealls-dating-service/app/controller"
	"github.com/dns2012/dealls-dating-service/app/domain/entity"
	"github.com/dns2012/dealls-dating-service/app/domain/exception"
	"github.com/dns2012/dealls-dating-service/app/usecase"
	"github.com/dns2012/dealls-dating-service/mocks"
	schemav1 "github.com/dns2012/dealls-dating-service/proto/schema/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	"testing"
)

func TestSchemaV1API_ListPackage(t *testing.T) {
	type input struct {
		ctx context.Context
		req *emptypb.Empty
	}

	type output struct {
		res *schemav1.ListPackageResponse
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
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx: context.Background(),
				req: &emptypb.Empty{},
			},
			out: &output{
				err: exception.Internal("INTERNAL_SERVER_ERROR"),
			},
			on: func(m *SchemaV1APIMock, i *input, o *output) {
				m.listPackageUsecase.On("Call", i.ctx).Return(nil, o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				req: &emptypb.Empty{},
			},
			out: &output{
				res: &schemav1.ListPackageResponse{
					Data: []*schemav1.Package{
						{
							Id: 1,
							Name: "Package 1",
							Description: "Package 1 For Your Swap Freedom",
							Price: 50000,
							UnlimitedSwap: false,
							TotalSwapPerDay: 50,
						},
					},
				},
			},
			on: func(m *SchemaV1APIMock, i *input, o *output) {
				m.listPackageUsecase.On("Call", i.ctx).Return(&usecase.ListPackageUsecaseResult{
					List: []*entity.Package{
						{
							Model:           gorm.Model{ID: uint(o.res.GetData()[0].GetId())},
							Name:            o.res.GetData()[0].GetName(),
							Description:     o.res.GetData()[0].GetDescription(),
							Price:           uint(o.res.GetData()[0].GetPrice()),
							UnlimitedSwap:   o.res.GetData()[0].GetUnlimitedSwap(),
							TotalSwapPerDay: uint(o.res.GetData()[0].GetTotalSwapPerDay()),
						},
					},
				}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SchemaV1APIMock{listPackageUsecase: mocks.ListPackageUsecase{}}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := new(controller.SchemaV1API)
			subject.ListPackageUsecase = &m.listPackageUsecase
			res, err := subject.ListPackage(tt.in.ctx, tt.in.req)

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