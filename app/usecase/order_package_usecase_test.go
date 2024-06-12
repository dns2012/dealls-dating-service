package usecase_test

import (
	"context"
	"errors"
	"github.com/dns2012/dealls-dating-service/app/domain/constant"
	"github.com/dns2012/dealls-dating-service/app/domain/entity"
	"github.com/dns2012/dealls-dating-service/app/domain/exception"
	"github.com/dns2012/dealls-dating-service/app/domain/repository"
	"github.com/dns2012/dealls-dating-service/app/usecase"
	"github.com/dns2012/dealls-dating-service/mocks"
	schemav1 "github.com/dns2012/dealls-dating-service/proto/schema/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/grpclog"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func TestOrderPackageUsecase_Call(t *testing.T) {
	type input struct {
		ctx context.Context
		req *usecase.OrderPackageUsecaseParams
	}

	type output struct {
		err error
	}

	type testcase struct {
		name string
		in   *input
		out  *output
		on   func(*UsecaseMock, *input, *output)
	}

	fullName := "User Dealls"
	name := strings.Fields(fullName)
	birthDate, _ := time.Parse("2006-01-02", "1999-09-09")
	userEntity := &entity.User{
		Nickname:    name[0],
		Email:       "user@dealls.com",
		Password:    "passwordhashed",
		Profile:     entity.Profile{
			FullName: fullName,
			ImageUrl: constant.DUMMY_IMAGE_URL,
			BirthAt:  birthDate,
			Gender:   uint(schemav1.Gender_MALE),
			Company:  "Dealls",
			JobTitle: "Backend Engineer",
		},
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
				req: &usecase.OrderPackageUsecaseParams{
					UserID:       1,
					PackageID:    1,
					TotalPayment: 50000,
				},
			},
			out: &output{
				err: exception.Internal("INTERNAL_SERVER_ERROR"),
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{UserID: i.req.UserID}).Return(nil, errors.New("INTERNAL_SERVER_ERROR"))
			},
		},
		{
			name: "NOT_FOUND",
			in: &input{
				ctx: context.Background(),
				req: &usecase.OrderPackageUsecaseParams{
					UserID:       1,
					PackageID:    1,
					TotalPayment: 50000,
				},
			},
			out: &output{
				err: exception.NotFound("User is not found. It might be deleted."),
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{UserID: i.req.UserID}).Return(nil, nil)
			},
		},
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx: context.Background(),
				req: &usecase.OrderPackageUsecaseParams{
					UserID:       1,
					PackageID:    1,
					TotalPayment: 50000,
				},
			},
			out: &output{
				err: exception.Internal("INTERNAL_SERVER_ERROR"),
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{UserID: i.req.UserID}).Return(userEntity, nil)
				m.packageRepository.On("Find", i.ctx, i.req.PackageID).Return(nil, errors.New("INTERNAL_SERVER_ERROR"))
			},
		},
		{
			name: "NOT_FOUND",
			in: &input{
				ctx: context.Background(),
				req: &usecase.OrderPackageUsecaseParams{
					UserID:       1,
					PackageID:    1,
					TotalPayment: 50000,
				},
			},
			out: &output{
				err: exception.NotFound("Package is not found. It might be deleted."),
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{UserID: i.req.UserID}).Return(userEntity, nil)
				m.packageRepository.On("Find", i.ctx, i.req.PackageID).Return(nil, nil)
			},
		},
		{
			name: "INVALID",
			in: &input{
				ctx: context.Background(),
				req: &usecase.OrderPackageUsecaseParams{
					UserID:       1,
					PackageID:    1,
					TotalPayment: 40000,
				},
			},
			out: &output{
				err: exception.Invalid("Order failed. Payment is insufficient."),
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{UserID: i.req.UserID}).Return(userEntity, nil)
				m.packageRepository.On("Find", i.ctx, i.req.PackageID).Return(packageEntity, nil)
			},
		},
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx: context.Background(),
				req: &usecase.OrderPackageUsecaseParams{
					UserID:       1,
					PackageID:    1,
					TotalPayment: 50000,
				},
			},
			out: &output{
				err: exception.Internal("INTERNAL_SERVER_ERROR"),
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{UserID: i.req.UserID}).Return(userEntity, nil)
				m.packageRepository.On("Find", i.ctx, i.req.PackageID).Return(packageEntity, nil)
				m.userPackageRepository.On("Create", i.ctx, &entity.UserPackage{
					UserID:    userEntity.ID,
					PackageID: packageEntity.ID,
				}).Return(nil, errors.New("INTERNAL_SERVER_ERROR"))
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				req: &usecase.OrderPackageUsecaseParams{
					UserID:       1,
					PackageID:    1,
					TotalPayment: 50000,
				},
			},
			out: &output{
				err: nil,
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{UserID: i.req.UserID}).Return(userEntity, nil)
				m.packageRepository.On("Find", i.ctx, i.req.PackageID).Return(packageEntity, nil)
				m.userPackageRepository.On("Create", i.ctx, &entity.UserPackage{
					UserID:    userEntity.ID,
					PackageID: packageEntity.ID,
				}).Return(nil, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &UsecaseMock{
				userRepository: mocks.UserRepository{},
				packageRepository: mocks.PackageRepository{},
				userPackageRepository: mocks.UserPackageRepository{},
				logger: grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := usecase.NewOrderPackageUsecase(&m.userRepository, &m.packageRepository, &m.userPackageRepository, m.logger)
			err := subject.Call(tt.in.ctx, tt.in.req)

			if tt.out.err != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.out.err, err)
			}
		})
	}
}
