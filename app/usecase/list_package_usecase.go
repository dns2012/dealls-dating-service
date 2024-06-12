package usecase

import (
	"context"
	"github.com/dns2012/dealls-dating-service/app/domain/entity"
	"github.com/dns2012/dealls-dating-service/app/domain/exception"
	"github.com/dns2012/dealls-dating-service/app/domain/repository"
	"google.golang.org/grpc/grpclog"
)

type ListPackageUsecase interface {
	Call(ctx context.Context) (*ListPackageUsecaseResult, error)
}

type ListPackageUsecaseResult struct {
	List []*entity.Package
}

type listPackageUsecaseImplementation struct {
	packageRepository repository.PackageRepository
	logger grpclog.LoggerV2
}

func (u *listPackageUsecaseImplementation) Call(ctx context.Context) (*ListPackageUsecaseResult, error) {
	list, err := u.packageRepository.List(ctx)
	if err != nil {
		u.logger.Errorf("Failed to get list packages: %v", err)
		return nil, exception.Internal(err.Error())
	}
	return &ListPackageUsecaseResult{List: list}, nil
}

func NewListPackageUsecase(
	packageRepository repository.PackageRepository,
	logger grpclog.LoggerV2,
) ListPackageUsecase {
	return &listPackageUsecaseImplementation{
		packageRepository: packageRepository,
		logger: logger,
	}
}