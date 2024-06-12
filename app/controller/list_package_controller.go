package controller

import (
	"context"
	schemav1 "github.com/dns2012/dealls-dating-service/proto/schema/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *SchemaV1API) ListPackage(ctx context.Context, req *emptypb.Empty) (*schemav1.ListPackageResponse, error)  {
	res, err := s.ListPackageUsecase.Call(ctx)
	if err != nil {
		return nil, err
	}

	var packages []*schemav1.Package
	for _, i := range res.List {
		packages = append(packages, &schemav1.Package{
			Id:              uint32(i.ID),
			Name:            i.Name,
			Description:     i.Description,
			Price: uint32(i.Price),
			UnlimitedSwap:   i.UnlimitedSwap,
			TotalSwapPerDay: uint32(i.TotalSwapPerDay),
		})
	}

	return &schemav1.ListPackageResponse{Data: packages}, nil
}
