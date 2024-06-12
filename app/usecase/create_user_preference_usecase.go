package usecase

import (
	"context"
	"github.com/dns2012/dealls-dating-service/app/domain/constant"
	"github.com/dns2012/dealls-dating-service/app/domain/entity"
	"github.com/dns2012/dealls-dating-service/app/domain/exception"
	"github.com/dns2012/dealls-dating-service/app/domain/repository"
	"google.golang.org/grpc/grpclog"
	"time"
)

type CreateUserPreferenceUsecase interface {
	Call(ctx context.Context, params *CreateUserPreferenceUsecaseParams) error
}

type CreateUserPreferenceUsecaseParams struct {
	UserID uint
	PreferenceUserID uint
	PreferenceType uint
}

type createUserPreferenceUsecaseImplementation struct {
	userRepository repository.UserRepository
	preferenceRepository repository.PreferenceRepository
	logger grpclog.LoggerV2
}

func (u *createUserPreferenceUsecaseImplementation) Call(ctx context.Context, params *CreateUserPreferenceUsecaseParams) error {
	user, err := u.userRepository.FindBy(ctx, &repository.UserFilterArgs{UserID: params.UserID})
	if err != nil {
		u.logger.Errorf("Failed to get user: %v", err)
		return exception.Internal(err.Error())
	}
	if user == nil {
		return exception.NotFound("User is not found. It might be deleted.")
	}

	preferenceUser, err := u.userRepository.FindBy(ctx, &repository.UserFilterArgs{UserID: params.PreferenceUserID})
	if err != nil {
		u.logger.Errorf("Failed to get preference user: %v", err)
		return exception.Internal(err.Error())
	}
	if preferenceUser == nil {
		return exception.NotFound("Preference user is not found. It might be deleted.")
	}

	preferenceCount, err := u.preferenceRepository.Count(ctx, &repository.PreferenceFilterArgs{
		UserID: user.ID,
		CreatedAt: time.Now().Format("2006-01-02"),
	})
	if err != nil {
		u.logger.Errorf("Failed to count preferences: %v", err)
		return exception.Internal(err.Error())
	}

	if user.UserPackage.PackageID > 0 {
		if !user.UserPackage.Package.UnlimitedSwap && uint(preferenceCount) >= user.UserPackage.Package.TotalSwapPerDay {
			return exception.Invalid("Preference quota has been exceeded. Please try again in the next day.")
		}
	} else {
		if preferenceCount >= constant.PREFERENCE_QUOTA {
			return exception.Invalid("Preference quota has been exceeded. Please try again in the next day.")
		}
	}

	preference, err := u.preferenceRepository.FindBy(ctx, &repository.PreferenceFilterArgs{
		UserID: user.ID,
		PreferenceUserID: preferenceUser.ID,
		CreatedAt: time.Now().Format("2006-01-02"),
	})
	if err != nil {
		u.logger.Errorf("Failed to count preferences: %v", err)
		return exception.Internal(err.Error())
	}
	if preference != nil {
		return exception.Exist("Preference user already exists. Please user another user preference.")
	}

	_, err = u.preferenceRepository.Create(ctx, &entity.Preference{
		UserID:           user.ID,
		PreferenceUserID: preferenceUser.ID,
		PreferenceType: params.PreferenceType,
	})
	if err != nil {
		u.logger.Errorf("Failed to create preference: %v", err)
		return exception.Internal(err.Error())
	}

	return nil
}

func NewCreateUserPreferenceUsecase(
	userRepository repository.UserRepository,
	preferenceRepository repository.PreferenceRepository,
	logger grpclog.LoggerV2,
) CreateUserPreferenceUsecase {
	return &createUserPreferenceUsecaseImplementation{
		userRepository: userRepository,
		preferenceRepository: preferenceRepository,
		logger: logger,
	}
}