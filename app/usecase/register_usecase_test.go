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

func TestRegisterUsecase_Call(t *testing.T) {
	type input struct {
		ctx context.Context
		req *usecase.RegisterUsecaseParams
	}

	type output struct {
		res *usecase.RegisteUsecaseResult
		err error
	}

	type testcase struct {
		name string
		in   *input
		out  *output
		on   func(*UsecaseMock, *input, *output)
	}

	tests := []testcase{
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx: context.Background(),
				req: &usecase.RegisterUsecaseParams{
					FullName:  "User Dealls",
					Email:     "user@dealls.com",
					Password:  "Password123+",
					BirthDate: "1999-09-09",
					Gender:    uint(schemav1.Gender_MALE),
					Company:   "Dealls",
					JobTitle:  "Backend Engineer",
				},
			},
			out: &output{
				err: exception.Internal("INTERNAL_SERVER_ERROR"),
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{Email: i.req.Email}).Return(nil, errors.New("INTERNAL_SERVER_ERROR"))
			},
		},
		{
			name: "ALREADY_EXIST",
			in: &input{
				ctx: context.Background(),
				req: &usecase.RegisterUsecaseParams{
					FullName:  "User Dealls",
					Email:     "user@dealls.com",
					Password:  "Password123+",
					BirthDate: "1999-09-09",
					Gender:    uint(schemav1.Gender_MALE),
					Company:   "Dealls",
					JobTitle:  "Backend Engineer",
				},
			},
			out: &output{
				err: exception.Exist("Email already exist. Please use another email."),
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{Email: i.req.Email}).Return(&entity.User{
					Model: gorm.Model{ID: 1},
				}, nil)
			},
		},
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx: context.Background(),
				req: &usecase.RegisterUsecaseParams{
					FullName:  "User Dealls",
					Email:     "user@dealls.com",
					Password:  "Password123+",
					BirthDate: "1999-09-09",
					Gender:    uint(schemav1.Gender_MALE),
					Company:   "Dealls",
					JobTitle:  "Backend Engineer",
				},
			},
			out: &output{
				err: exception.Internal("INTERNAL_SERVER_ERROR"),
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{Email: i.req.Email}).Return(nil, nil)
				m.passwordManager.On("CreateHashPassword", i.req.Password).Return("", errors.New("INTERNAL_SERVER_ERROR"))
			},
		},
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx: context.Background(),
				req: &usecase.RegisterUsecaseParams{
					FullName:  "User Dealls",
					Email:     "user@dealls.com",
					Password:  "Password123+",
					BirthDate: "1999-09-09",
					Gender:    uint(schemav1.Gender_MALE),
					Company:   "Dealls",
					JobTitle:  "Backend Engineer",
				},
			},
			out: &output{
				err: exception.Internal("INTERNAL_SERVER_ERROR"),
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{Email: i.req.Email}).Return(nil, nil)
				m.passwordManager.On("CreateHashPassword", i.req.Password).Return("passwordhashed", nil)
				name := strings.Fields(i.req.FullName)
				birthDate, _ := time.Parse("2006-01-02", i.req.BirthDate)
				user := &entity.User{
					Nickname:    name[0],
					Email:       i.req.Email,
					Password:    "passwordhashed",
					Profile:     entity.Profile{
						FullName: i.req.FullName,
						ImageUrl: constant.DUMMY_IMAGE_URL,
						BirthAt:  birthDate,
						Gender:   i.req.Gender,
						Company:  i.req.Company,
						JobTitle: i.req.JobTitle,
					},
				}
				m.userRepository.On("Create", i.ctx, user).Return(nil, errors.New("INTERNAL_SERVER_ERROR"))
			},
		},
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx: context.Background(),
				req: &usecase.RegisterUsecaseParams{
					FullName:  "User Dealls",
					Email:     "user@dealls.com",
					Password:  "Password123+",
					BirthDate: "1999-09-09",
					Gender:    uint(schemav1.Gender_MALE),
					Company:   "Dealls",
					JobTitle:  "Backend Engineer",
				},
			},
			out: &output{
				err: exception.Internal("INTERNAL_SERVER_ERROR"),
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{Email: i.req.Email}).Return(nil, nil)
				m.passwordManager.On("CreateHashPassword", i.req.Password).Return("passwordhashed", nil)
				name := strings.Fields(i.req.FullName)
				birthDate, _ := time.Parse("2006-01-02", i.req.BirthDate)
				user := &entity.User{
					Nickname: name[0],
					Email:    i.req.Email,
					Password: "passwordhashed",
					Profile: entity.Profile{
						FullName: i.req.FullName,
						ImageUrl: constant.DUMMY_IMAGE_URL,
						BirthAt:  birthDate,
						Gender:   i.req.Gender,
						Company:  i.req.Company,
						JobTitle: i.req.JobTitle,
					},
				}
				m.userRepository.On("Create", i.ctx, user).Return(user, nil)
				m.tokenManager.On("CreateToken", i.ctx, user).Return("", errors.New("INTERNAL_SERVER_ERROR"))
			},
		},
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx: context.Background(),
				req: &usecase.RegisterUsecaseParams{
					FullName:  "User Dealls",
					Email:     "user@dealls.com",
					Password:  "Password123+",
					BirthDate: "1999-09-09",
					Gender:    uint(schemav1.Gender_MALE),
					Company:   "Dealls",
					JobTitle:  "Backend Engineer",
				},
			},
			out: &output{
				res: &usecase.RegisteUsecaseResult{Token: "jwttoken"},
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{Email: i.req.Email}).Return(nil, nil)
				m.passwordManager.On("CreateHashPassword", i.req.Password).Return("passwordhashed", nil)
				name := strings.Fields(i.req.FullName)
				birthDate, _ := time.Parse("2006-01-02", i.req.BirthDate)
				user := &entity.User{
					Nickname: name[0],
					Email:    i.req.Email,
					Password: "passwordhashed",
					Profile: entity.Profile{
						FullName: i.req.FullName,
						ImageUrl: constant.DUMMY_IMAGE_URL,
						BirthAt:  birthDate,
						Gender:   i.req.Gender,
						Company:  i.req.Company,
						JobTitle: i.req.JobTitle,
					},
				}
				m.userRepository.On("Create", i.ctx, user).Return(user, nil)
				m.tokenManager.On("CreateToken", i.ctx, user).Return("jwttoken", nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &UsecaseMock{
				userRepository: mocks.UserRepository{},
				passwordManager: mocks.PasswordManager{},
				tokenManager: mocks.TokenManager{},
				logger: grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard),
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := usecase.NewRegisterUsecase(&m.userRepository, &m.passwordManager, &m.tokenManager, m.logger)
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
