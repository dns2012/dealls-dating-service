package usecase

import (
	"context"
	"github.com/dns2012/dealls-dating-service/app/domain/entity"
	"github.com/dns2012/dealls-dating-service/app/domain/exception"
	"github.com/dns2012/dealls-dating-service/app/domain/manager"
	"github.com/dns2012/dealls-dating-service/app/domain/repository"
	"google.golang.org/grpc/grpclog"
	"time"
)

type ListUserUsecase interface {
	Call(ctx context.Context, params *ListUserUsecaseParams) (*ListUserUsecaseResult, error)
}

type ListUserUsecaseParams struct {
	UserID uint
	Page uint
	PageSize uint
}

type ListUserUsecaseResult struct {
	Page *manager.Pagination
	List []*entity.User
}

type listUserUsecaseImplementation struct {
	userRepository repository.UserRepository
	preferenceRepository repository.PreferenceRepository
	logger grpclog.LoggerV2
}

func (u *listUserUsecaseImplementation) Call(ctx context.Context, params *ListUserUsecaseParams) (*ListUserUsecaseResult, error) {
	user, err := u.userRepository.FindBy(ctx, &repository.UserFilterArgs{UserID: params.UserID})
	if err != nil {
		u.logger.Errorf("Failed to get user: %v", err)
		return nil, exception.Internal(err.Error())
	}
	if user == nil {
		return nil, exception.NotFound("User is not found. It might be deleted.")
	}

	preferences, err := u.preferenceRepository.ListBy(ctx, &repository.PreferenceFilterArgs{
		UserID: user.ID,
		CreatedAt: time.Now().Format("2006-01-02"),
	})
	if err != nil {
		u.logger.Errorf("Failed to get list preferences: %v", err)
		return nil, exception.Internal(err.Error())
	}

	preferenceTargetIDs := []uint{user.ID}
	for _, i := range preferences {
		preferenceTargetIDs = append(preferenceTargetIDs, i.PreferenceUserID)
	}

	page, list, err := u.userRepository.ListBy(
		ctx,
		&repository.UserFilterArgs{ExcludeUserID: preferenceTargetIDs},
		manager.Pagination{
			Page: int(params.Page),
			PageSize: int(params.PageSize),
		},
	)
	if err != nil {
		u.logger.Errorf("Failed to get list users: %v", err)
		return nil, exception.Internal(err.Error())
	}

	return &ListUserUsecaseResult{Page: page, List: list}, nil
}

func NewListUserUsecase(
	userRepository repository.UserRepository,
	preferenceRepository repository.PreferenceRepository,
	logger grpclog.LoggerV2,
) ListUserUsecase {
	return &listUserUsecaseImplementation{
		userRepository: userRepository,
		preferenceRepository: preferenceRepository,
		logger: logger,
	}
}