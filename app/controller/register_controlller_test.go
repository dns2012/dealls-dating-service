package controller_test

import (
	"context"
	"github.com/dns2012/dealls-dating-service/app/controller"
	"github.com/dns2012/dealls-dating-service/app/domain/exception"
	"github.com/dns2012/dealls-dating-service/app/usecase"
	"github.com/dns2012/dealls-dating-service/mocks"
	schemav1 "github.com/dns2012/dealls-dating-service/proto/schema/v1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSchemaV1API_Register(t *testing.T) {
	type input struct {
		ctx context.Context
		req *schemav1.RegisterRequest
	}

	type output struct {
		res *schemav1.AuthResponse
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
				req: &schemav1.RegisterRequest{},
			},
			out: &output{},
			on: func(m *SchemaV1APIMock, i *input, o *output) {
				o.err = exception.Invalid(i.req.Validate().Error())
			},
		},
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx: context.Background(),
				req: &schemav1.RegisterRequest{
					FullName:        "Dealls User",
					Email:           "user@dealls.com",
					Password:        "Password123+",
					ConfirmPassword: "Password123+",
					BirthDate:       "1999-09-09",
					Gender:          schemav1.Gender_MALE,
					Company:         "Dealls",
					JobTitle:        "Backend Engineer",
				},
			},
			out: &output{
				err: exception.Internal("INTERNAL_SERVER_ERROR"),
			},
			on: func(m *SchemaV1APIMock, i *input, o *output) {
				m.registerUsecase.On("Call", i.ctx, &usecase.RegisterUsecaseParams{
					FullName:  i.req.GetFullName(),
					Email:     i.req.GetEmail(),
					Password:  i.req.GetPassword(),
					BirthDate: i.req.GetBirthDate(),
					Gender: uint(i.req.GetGender()),
					Company:   i.req.GetCompany(),
					JobTitle:  i.req.GetJobTitle(),
				}).Return(nil, o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				req: &schemav1.RegisterRequest{
					FullName:        "Dealls User",
					Email:           "user@dealls.com",
					Password:        "Password123+",
					ConfirmPassword: "Password123+",
					BirthDate:       "1999-09-09",
					Gender:          schemav1.Gender_MALE,
					Company:         "Dealls",
					JobTitle:        "Backend Engineer",
				},
			},
			out: &output{
				res: &schemav1.AuthResponse{
					Data: &schemav1.AuthResponseData{
						AccessToken: "jwttoken",
					},
				},
			},
			on: func(m *SchemaV1APIMock, i *input, o *output) {
				m.registerUsecase.On("Call", i.ctx, &usecase.RegisterUsecaseParams{
					FullName:  i.req.GetFullName(),
					Email:     i.req.GetEmail(),
					Password:  i.req.GetPassword(),
					BirthDate: i.req.GetBirthDate(),
					Gender: uint(i.req.GetGender()),
					Company:   i.req.GetCompany(),
					JobTitle:  i.req.GetJobTitle(),
				}).Return(&usecase.RegisteUsecaseResult{Token: o.res.Data.AccessToken}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SchemaV1APIMock{registerUsecase: mocks.RegisterUsecase{}}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := new(controller.SchemaV1API)
			subject.RegisterUsecase = &m.registerUsecase
			res, err := subject.Register(tt.in.ctx, tt.in.req)

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
