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
	"gorm.io/gorm"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func TestCreateUserPreference_Call(t *testing.T) {
	type input struct {
		ctx context.Context
		req *usecase.CreateUserPreferenceUsecaseParams
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

	tests := []testcase{
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx: context.Background(),
				req: &usecase.CreateUserPreferenceUsecaseParams{
					UserID:       1,
					PreferenceUserID:    2,
					PreferenceType: uint(schemav1.PreferenceType_LIKE),
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
				req: &usecase.CreateUserPreferenceUsecaseParams{
					UserID:       1,
					PreferenceUserID:    2,
					PreferenceType: uint(schemav1.PreferenceType_LIKE),
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
				req: &usecase.CreateUserPreferenceUsecaseParams{
					UserID:       1,
					PreferenceUserID:    2,
					PreferenceType: uint(schemav1.PreferenceType_LIKE),
				},
			},
			out: &output{
				err: exception.Internal("INTERNAL_SERVER_ERROR"),
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{UserID: i.req.UserID}).Return(userEntity, nil)
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{UserID: i.req.PreferenceUserID}).Return(nil, errors.New("INTERNAL_SERVER_ERROR"))
			},
		},
		{
			name: "NOT_FOUND",
			in: &input{
				ctx: context.Background(),
				req: &usecase.CreateUserPreferenceUsecaseParams{
					UserID:       1,
					PreferenceUserID:    2,
					PreferenceType: uint(schemav1.PreferenceType_LIKE),
				},
			},
			out: &output{
				err: exception.NotFound("Preference user is not found. It might be deleted."),
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{UserID: i.req.UserID}).Return(userEntity, nil)
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{UserID: i.req.PreferenceUserID}).Return(nil, nil)
			},
		},
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx: context.Background(),
				req: &usecase.CreateUserPreferenceUsecaseParams{
					UserID:       1,
					PreferenceUserID:    2,
					PreferenceType: uint(schemav1.PreferenceType_LIKE),
				},
			},
			out: &output{
				err: exception.Internal("INTERNAL_SERVER_ERROR"),
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{UserID: i.req.UserID}).Return(userEntity, nil)
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{UserID: i.req.PreferenceUserID}).Return(userEntity, nil)
				m.preferenceRepository.On("Count", i.ctx, &repository.PreferenceFilterArgs{
					UserID: userEntity.ID,
					CreatedAt: time.Now().Format("2006-01-02"),
				}).Return(int64(0), errors.New("INTERNAL_SERVER_ERROR"))
			},
		},
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx: context.Background(),
				req: &usecase.CreateUserPreferenceUsecaseParams{
					UserID:       1,
					PreferenceUserID:    2,
					PreferenceType: uint(schemav1.PreferenceType_LIKE),
				},
			},
			out: &output{
				err: exception.Invalid("Preference quota has been exceeded. Please try again in the next day."),
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{UserID: i.req.UserID}).Return(userEntity, nil)
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{UserID: i.req.PreferenceUserID}).Return(userEntity, nil)
				m.preferenceRepository.On("Count", i.ctx, &repository.PreferenceFilterArgs{
					UserID: userEntity.ID,
					CreatedAt: time.Now().Format("2006-01-02"),
				}).Return(int64(10), nil)
			},
		},
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx: context.Background(),
				req: &usecase.CreateUserPreferenceUsecaseParams{
					UserID:       1,
					PreferenceUserID:    2,
					PreferenceType: uint(schemav1.PreferenceType_LIKE),
				},
			},
			out: &output{
				err: exception.Invalid("Preference quota has been exceeded. Please try again in the next day."),
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				userEntity.UserPackage = entity.UserPackage{
					Model: gorm.Model{ID: 1},
					PackageID: 1,
					Package: entity.Package{
						Model: gorm.Model{ID: 1},
						UnlimitedSwap: false,
						TotalSwapPerDay: 15,
					},
				}
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{UserID: i.req.UserID}).Return(userEntity, nil)
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{UserID: i.req.PreferenceUserID}).Return(userEntity, nil)
				m.preferenceRepository.On("Count", i.ctx, &repository.PreferenceFilterArgs{
					UserID: userEntity.ID,
					CreatedAt: time.Now().Format("2006-01-02"),
				}).Return(int64(15), nil)
			},
		},
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx: context.Background(),
				req: &usecase.CreateUserPreferenceUsecaseParams{
					UserID:       1,
					PreferenceUserID:    2,
					PreferenceType: uint(schemav1.PreferenceType_LIKE),
				},
			},
			out: &output{
				err: exception.Internal("INTERNAL_SERVER_ERROR"),
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{UserID: i.req.UserID}).Return(userEntity, nil)
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{UserID: i.req.PreferenceUserID}).Return(userEntity, nil)
				m.preferenceRepository.On("Count", i.ctx, &repository.PreferenceFilterArgs{
					UserID: userEntity.ID,
					CreatedAt: time.Now().Format("2006-01-02"),
				}).Return(int64(9), nil)
				m.preferenceRepository.On("FindBy", i.ctx, &repository.PreferenceFilterArgs{
					UserID: userEntity.ID,
					PreferenceUserID: userEntity.ID,
					CreatedAt: time.Now().Format("2006-01-02"),
				}).Return(nil, errors.New("INTERNAL_SERVER_ERROR"))
			},
		},
		{
			name: "ALREADY_EXIST",
			in: &input{
				ctx: context.Background(),
				req: &usecase.CreateUserPreferenceUsecaseParams{
					UserID:       1,
					PreferenceUserID:    2,
					PreferenceType: uint(schemav1.PreferenceType_LIKE),
				},
			},
			out: &output{
				err: exception.Exist("Preference user already exists. Please user another user preference."),
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{UserID: i.req.UserID}).Return(userEntity, nil)
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{UserID: i.req.PreferenceUserID}).Return(userEntity, nil)
				m.preferenceRepository.On("Count", i.ctx, &repository.PreferenceFilterArgs{
					UserID: userEntity.ID,
					CreatedAt: time.Now().Format("2006-01-02"),
				}).Return(int64(9), nil)
				m.preferenceRepository.On("FindBy", i.ctx, &repository.PreferenceFilterArgs{
					UserID: userEntity.ID,
					PreferenceUserID: userEntity.ID,
					CreatedAt: time.Now().Format("2006-01-02"),
				}).Return(&entity.Preference{UserID: userEntity.ID, PreferenceUserID: userEntity.ID}, nil)
			},
		},
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx: context.Background(),
				req: &usecase.CreateUserPreferenceUsecaseParams{
					UserID:       1,
					PreferenceUserID:    2,
					PreferenceType: uint(schemav1.PreferenceType_LIKE),
				},
			},
			out: &output{
				err: exception.Internal("INTERNAL_SERVER_ERROR"),
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{UserID: i.req.UserID}).Return(userEntity, nil)
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{UserID: i.req.PreferenceUserID}).Return(userEntity, nil)
				m.preferenceRepository.On("Count", i.ctx, &repository.PreferenceFilterArgs{
					UserID: userEntity.ID,
					CreatedAt: time.Now().Format("2006-01-02"),
				}).Return(int64(9), nil)
				m.preferenceRepository.On("FindBy", i.ctx, &repository.PreferenceFilterArgs{
					UserID: userEntity.ID,
					PreferenceUserID: userEntity.ID,
					CreatedAt: time.Now().Format("2006-01-02"),
				}).Return(nil, nil)
				m.preferenceRepository.On("Create", i.ctx, &entity.Preference{
					UserID:           userEntity.ID,
					PreferenceUserID: userEntity.ID,
					PreferenceType: i.req.PreferenceType,
				}).Return(nil, errors.New("INTERNAL_SERVER_ERROR"))
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				req: &usecase.CreateUserPreferenceUsecaseParams{
					UserID:       1,
					PreferenceUserID:    2,
					PreferenceType: uint(schemav1.PreferenceType_LIKE),
				},
			},
			out: &output{
				err: nil,
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{UserID: i.req.UserID}).Return(userEntity, nil)
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{UserID: i.req.PreferenceUserID}).Return(userEntity, nil)
				m.preferenceRepository.On("Count", i.ctx, &repository.PreferenceFilterArgs{
					UserID: userEntity.ID,
					CreatedAt: time.Now().Format("2006-01-02"),
				}).Return(int64(9), nil)
				m.preferenceRepository.On("FindBy", i.ctx, &repository.PreferenceFilterArgs{
					UserID: userEntity.ID,
					PreferenceUserID: userEntity.ID,
					CreatedAt: time.Now().Format("2006-01-02"),
				}).Return(nil, nil)
				m.preferenceRepository.On("Create", i.ctx, &entity.Preference{
					UserID:           userEntity.ID,
					PreferenceUserID: userEntity.ID,
					PreferenceType: i.req.PreferenceType,
				}).Return(nil, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &UsecaseMock{
				userRepository: mocks.UserRepository{},
				preferenceRepository: mocks.PreferenceRepository{},
				logger: grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := usecase.NewCreateUserPreferenceUsecase(&m.userRepository, &m.preferenceRepository, m.logger)
			err := subject.Call(tt.in.ctx, tt.in.req)

			if tt.out.err != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.out.err, err)
			}
		})
	}
}