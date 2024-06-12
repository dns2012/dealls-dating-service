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

func TestLoginUsecase_Call(t *testing.T) {
	type input struct {
		ctx context.Context
		req *usecase.LoginUsecaseParams
	}

	type output struct {
		res *usecase.LoginUsecaseResult
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
				req: &usecase.LoginUsecaseParams{
					Email:     "user@dealls.com",
					Password:  "Password123+",
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
			name: "NOT_FOUND",
			in: &input{
				ctx: context.Background(),
				req: &usecase.LoginUsecaseParams{
					Email:     "user@dealls.com",
					Password:  "Password123+",
				},
			},
			out: &output{
				err: exception.NotFound("User is not found. It might be deleted."),
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{Email: i.req.Email}).Return(nil, nil)
			},
		},
		{
			name: "INVALID",
			in: &input{
				ctx: context.Background(),
				req: &usecase.LoginUsecaseParams{
					Email:     "user@dealls.com",
					Password:  "Password123+",
				},
			},
			out: &output{
				err: exception.Invalid("Password is not correct. Please pass a valid password for given email."),
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{Email: i.req.Email}).Return(userEntity, nil)
				m.passwordManager.On("CheckHashPassword", i.req.Password, userEntity.Password).Return(false)
			},
		},
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx: context.Background(),
				req: &usecase.LoginUsecaseParams{
					Email:     "user@dealls.com",
					Password:  "Password123+",
				},
			},
			out: &output{
				err: exception.Internal("INTERNAL_SERVER_ERROR"),
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{Email: i.req.Email}).Return(userEntity, nil)
				m.passwordManager.On("CheckHashPassword", i.req.Password, userEntity.Password).Return(true)
				m.tokenManager.On("CreateToken", i.ctx, userEntity).Return("", errors.New("INTERNAL_SERVER_ERROR"))
			},
		},
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx: context.Background(),
				req: &usecase.LoginUsecaseParams{
					Email:     "user@dealls.com",
					Password:  "Password123+",
				},
			},
			out: &output{
				res: &usecase.LoginUsecaseResult{Token: "jwttoken"},
			},
			on: func(m *UsecaseMock, i *input, o *output) {
				m.userRepository.On("FindBy", i.ctx, &repository.UserFilterArgs{Email: i.req.Email}).Return(userEntity, nil)
				m.passwordManager.On("CheckHashPassword", i.req.Password, userEntity.Password).Return(true)
				m.tokenManager.On("CreateToken", i.ctx, userEntity).Return("jwttoken", nil)
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

			subject := usecase.NewLoginUsecase(&m.userRepository, &m.passwordManager, &m.tokenManager, m.logger)
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
