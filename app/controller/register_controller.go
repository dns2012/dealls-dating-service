package controller

import (
	"context"
	"github.com/dns2012/dealls-dating-service/app/domain/exception"
	"github.com/dns2012/dealls-dating-service/app/usecase"
	schemav1 "github.com/dns2012/dealls-dating-service/proto/schema/v1"
)

func (s *SchemaV1API) Register(ctx context.Context, req *schemav1.RegisterRequest) (*schemav1.AuthResponse, error)  {
	if err := req.Validate(); err != nil {
		return nil, exception.Invalid(err.Error())
	}

	res, err := s.RegisterUsecase.Call(ctx, &usecase.RegisterUsecaseParams{
		FullName: req.GetFullName(),
		Email:     req.GetEmail(),
		Password:  req.GetPassword(),
		BirthDate: req.GetBirthDate(),
		Gender: uint(req.GetGender()),
		Company:   req.GetCompany(),
		JobTitle:  req.GetJobTitle(),
	})
	if err != nil {
		return nil, err
	}

	return &schemav1.AuthResponse{
		Data: &schemav1.AuthResponseData{
			AccessToken: res.Token,
		},
	}, nil
}
