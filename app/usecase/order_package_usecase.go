package usecase

import (
	"context"
	"github.com/dns2012/dealls-dating-service/app/domain/entity"
	"github.com/dns2012/dealls-dating-service/app/domain/exception"
	"github.com/dns2012/dealls-dating-service/app/domain/repository"
	"google.golang.org/grpc/grpclog"
)

type OrderPackageUsecase interface {
	Call(ctx context.Context, params *OrderPackageUsecaseParams) error
}

type OrderPackageUsecaseParams struct {
	UserID uint
	PackageID uint
	TotalPayment uint
}

type orderPackageUsecaseImplementation struct {
	userRepository repository.UserRepository
	packageRepository repository.PackageRepository
	userPackageRepsitory repository.UserPackageRepository
	logger grpclog.LoggerV2
}

func (u *orderPackageUsecaseImplementation) Call(ctx context.Context, params *OrderPackageUsecaseParams) error {
	user, err := u.userRepository.FindBy(ctx, &repository.UserFilterArgs{UserID: params.UserID})
	if err != nil {
		u.logger.Errorf("Failed to get user: %v", err)
		return exception.Internal(err.Error())
	}
	if user == nil {
		return exception.NotFound("User is not found. It might be deleted.")
	}

	pkg, err := u.packageRepository.Find(ctx, params.PackageID)
	if err != nil {
		u.logger.Errorf("Failed to get package: %v", err)
		return exception.Internal(err.Error())
	}
	if pkg == nil {
		return exception.NotFound("Package is not found. It might be deleted.")
	}

	if params.TotalPayment < pkg.Price {
		return exception.Invalid("Order failed. Payment is insufficient.")
	}

	_, err = u.userPackageRepsitory.Create(ctx, &entity.UserPackage{
		UserID:    user.ID,
		PackageID: pkg.ID,
	})
	if err != nil {
		u.logger.Errorf("Failed to create user package: %v", err)
		return exception.Internal(err.Error())
	}
	return nil
}

func NewOrderPackageUsecase(
	userRepository repository.UserRepository,
	packageRepository repository.PackageRepository,
	userPackageRepsitory repository.UserPackageRepository,
	logger grpclog.LoggerV2,
) OrderPackageUsecase {
	return &orderPackageUsecaseImplementation{
		userRepository: userRepository,
		packageRepository: packageRepository,
		userPackageRepsitory: userPackageRepsitory,
		logger: logger,
	}
}