package usecase

import (
	"context"
	"github.com/dns2012/dealls-dating-service/app/domain/entity"
	"github.com/dns2012/dealls-dating-service/app/domain/exception"
	"github.com/dns2012/dealls-dating-service/app/domain/repository"
	"google.golang.org/grpc/grpclog"
)

type MeUsecase interface {
	Call(ctx context.Context, params *MeUsecaseParams) (*MeUsecaseResult, error)
}

type MeUsecaseParams struct {
	UserID uint
}

type MeUsecaseResult struct {
	User *entity.User
}

type meUsecaseImplementation struct {
	userRepository repository.UserRepository
	logger grpclog.LoggerV2
}

func (u *meUsecaseImplementation) Call(ctx context.Context, params *MeUsecaseParams) (*MeUsecaseResult, error) {
	user, err := u.userRepository.FindBy(ctx, &repository.UserFilterArgs{UserID: params.UserID})
	if err != nil {
		u.logger.Errorf("Failed to get user: %v", err)
		return nil, exception.Internal(err.Error())
	}
	if user == nil {
		return nil, exception.NotFound("User is not found. It might be deleted.")
	}

	return &MeUsecaseResult{User: user}, nil
}

func NewMeUsecase(
	userRepository repository.UserRepository,
	logger grpclog.LoggerV2,
) MeUsecase {
	return &meUsecaseImplementation{
		userRepository: userRepository,
		logger: logger,
	}
}