package usecase_test

import (
	"context"
	"errors"
	"github.com/dns2012/dealls-dating-service/app/domain/constant"
	"github.com/dns2012/dealls-dating-service/app/domain/entity"
	"github.com/dns2012/dealls-dating-service/app/domain/exception"
	"github.com/dns2012/dealls-dating-service/app/domain/manager"
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

func TestListUserUsecase_Call(t *testing.T) {
	type input struct {
		ctx context.Context
		req *usecase.ListUserUsecaseParams
	}

	type output struct {
		res *usecase.ListUserUsecaseResult
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
				req: &usecase.ListUserUsecaseParams{
					UserID:       1,
					Page:    1,
					PageSize: 10,
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
				req: &usecase.ListUserUsecaseParams{
					UserID:  1,
					Page:    1,
					PageSize: 10,
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
				req: &usecase.ListUserUsecaseParams{
					UserID:       1,
					Page:    1,
					PageSize: 10,
				},
			},
			out: &output{
				err: exception.Internal("INTERNAL_SERVER_ERROR"),
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{UserID: i.req.UserID}).Return(userEntity, nil)
				m.preferenceRepository.On("ListBy", i.ctx, &repository.PreferenceFilterArgs{
					UserID: userEntity.ID,
					CreatedAt: time.Now().Format("2006-01-02"),
				}).Return(nil, errors.New("INTERNAL_SERVER_ERROR"))
			},
		},
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx: context.Background(),
				req: &usecase.ListUserUsecaseParams{
					UserID:       1,
					Page:    1,
					PageSize: 10,
				},
			},
			out: &output{
				err: exception.Internal("INTERNAL_SERVER_ERROR"),
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{UserID: i.req.UserID}).Return(userEntity, nil)
				m.preferenceRepository.On("ListBy", i.ctx, &repository.PreferenceFilterArgs{
					UserID: userEntity.ID,
					CreatedAt: time.Now().Format("2006-01-02"),
				}).Return([]*entity.Preference{
					{
						UserID: userEntity.ID,
						PreferenceUserID: 2,
					},
				}, nil)
				m.userRepository.On("ListBy", i.ctx, &repository.UserFilterArgs{ExcludeUserID: []uint{userEntity.ID, 2}}, manager.Pagination{
					Page:     int(i.req.Page),
					PageSize: int(i.req.PageSize),
				}).Return(nil, nil, errors.New("INTERNAL_SERVER_ERROR"))
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				req: &usecase.ListUserUsecaseParams{
					UserID:       1,
					Page:    1,
					PageSize: 10,
				},
			},
			out: &output{
				res: &usecase.ListUserUsecaseResult{
					Page: &manager.Pagination{
						Page:      1,
						PageSize:  10,
					},
					List: []*entity.User{userEntity},
				},
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{UserID: i.req.UserID}).Return(userEntity, nil)
				m.preferenceRepository.On("ListBy", i.ctx, &repository.PreferenceFilterArgs{
					UserID: userEntity.ID,
					CreatedAt: time.Now().Format("2006-01-02"),
				}).Return([]*entity.Preference{
					{
						UserID: userEntity.ID,
						PreferenceUserID: 2,
					},
				}, nil)
				pagination := manager.Pagination{
					Page:     int(i.req.Page),
					PageSize: int(i.req.PageSize),
				}
				m.userRepository.On("ListBy", i.ctx, &repository.UserFilterArgs{ExcludeUserID: []uint{userEntity.ID, 2}}, pagination).Return(&pagination, []*entity.User{userEntity}, nil)
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

			subject := usecase.NewListUserUsecase(&m.userRepository, &m.preferenceRepository, m.logger)
			res, err := subject.Call(tt.in.ctx, tt.in.req)

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
