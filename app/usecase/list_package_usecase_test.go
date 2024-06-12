package usecase_test

import (
	"context"
	"errors"
	"github.com/dns2012/dealls-dating-service/app/domain/entity"
	"github.com/dns2012/dealls-dating-service/app/domain/exception"
	"github.com/dns2012/dealls-dating-service/app/usecase"
	"github.com/dns2012/dealls-dating-service/mocks"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/grpclog"
	"io/ioutil"
	"os"
	"testing"
)

func TestListPackageUsecase_Call(t *testing.T) {
	type input struct {
		ctx context.Context
	}

	type output struct {
		res *usecase.ListPackageUsecaseResult
		err error
	}

	type testcase struct {
		name string
		in   *input
		out  *output
		on   func(*UsecaseMock, *input, *output)
	}

	packageEntity := &entity.Package{
		Name:            "Package 1",
		Description:     "Package 1 For Your Swap Freedom",
		Price:           50000,
		UnlimitedSwap:   false,
		TotalSwapPerDay: 50,
	}
	tests := []testcase{
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx: context.Background(),
			},
			out: &output{
				err: exception.Internal("INTERNAL_SERVER_ERROR"),
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.packageRepository.On("List", i.ctx).Return(nil, errors.New("INTERNAL_SERVER_ERROR"))
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
			},
			out: &output{
				res: &usecase.ListPackageUsecaseResult{
					List: []*entity.Package{packageEntity},
				},
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.packageRepository.On("List", i.ctx).Return([]*entity.Package{packageEntity}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &UsecaseMock{
				packageRepository: mocks.PackageRepository{},
				logger: grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := usecase.NewListPackageUsecase(&m.packageRepository, m.logger)
			res, err := subject.Call(tt.in.ctx)

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
