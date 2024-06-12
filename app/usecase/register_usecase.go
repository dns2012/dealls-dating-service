package usecase

import (
	"context"
	"github.com/dns2012/dealls-dating-service/app/domain/constant"
	"github.com/dns2012/dealls-dating-service/app/domain/entity"
	"github.com/dns2012/dealls-dating-service/app/domain/exception"
	"github.com/dns2012/dealls-dating-service/app/domain/manager"
	"github.com/dns2012/dealls-dating-service/app/domain/repository"
	"google.golang.org/grpc/grpclog"
	"strings"
	"time"
)

type RegisterUsecase interface {
	Call(ctx context.Context, params *RegisterUsecaseParams) (*RegisteUsecaseResult, error)
}

type RegisterUsecaseParams struct {
	FullName string
	Email string
	Password string
	BirthDate string
	Gender uint
	Company string
	JobTitle string
}

type RegisteUsecaseResult struct {
	Token string
}

type registerUsecaseImplementation struct {
	userRepository repository.UserRepository
	passwordManager manager.PasswordManager
	tokenManager manager.TokenManager
	logger grpclog.LoggerV2
}

func (u *registerUsecaseImplementation) Call(ctx context.Context, params *RegisterUsecaseParams) (*RegisteUsecaseResult, error) {
	user, err := u.userRepository.FindBy(ctx, &repository.UserFilterArgs{Email: params.Email})
	if err != nil {
		u.logger.Errorf("Failed to get user: %v", err)
		return nil, exception.Internal(err.Error())
	}
	if user != nil {
		return nil, exception.Exist("Email already exist. Please use another email.")
	}

	password, err := u.passwordManager.CreateHashPassword(params.Password)
	if err != nil {
		u.logger.Errorf("Failed to hash password: %v", err)
		return nil, exception.Internal(err.Error())
	}

	name := strings.Fields(params.FullName)
	birthDate, _ := time.Parse("2006-01-02", params.BirthDate)
	user, err = u.userRepository.Create(ctx, &entity.User{
		Nickname:  name[0],
		Email:     params.Email,
		Password:  password,
		Profile:   entity.Profile{
			FullName: params.FullName,
			ImageUrl: constant.DUMMY_IMAGE_URL,
			BirthAt:  birthDate,
			Gender:   params.Gender,
			Company:  params.Company,
			JobTitle: params.JobTitle,
		},
	})
	if err != nil {
		u.logger.Errorf("Failed to create user: %v", err)
		return nil, exception.Internal(err.Error())
	}

	token, err := u.tokenManager.CreateToken(ctx, user)
	if err != nil {
		u.logger.Errorf("Failed to create token: %v", err)
		return nil, exception.Internal(err.Error())
	}

	return &RegisteUsecaseResult{Token: token}, nil
}

func NewRegisterUsecase(
	userRepository repository.UserRepository,
	passwordManager manager.PasswordManager,
	tokenManager manager.TokenManager,
	logger grpclog.LoggerV2,
) RegisterUsecase {
	return &registerUsecaseImplementation{
		userRepository: userRepository,
		passwordManager: passwordManager,
		tokenManager: tokenManager,
		logger: logger,
	}
}