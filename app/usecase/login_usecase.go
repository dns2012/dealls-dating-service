package usecase

import (
	"context"
	"github.com/dns2012/dealls-dating-service/app/domain/exception"
	"github.com/dns2012/dealls-dating-service/app/domain/manager"
	"github.com/dns2012/dealls-dating-service/app/domain/repository"
	"google.golang.org/grpc/grpclog"
)

type LoginUsecase interface {
	Call(ctx context.Context, params *LoginUsecaseParams) (*LoginUsecaseResult, error)
}

type LoginUsecaseParams struct {
	Email string
	Password string
}

type LoginUsecaseResult struct {
	Token string
}

type loginUsecaseImplementation struct {
	userRepository repository.UserRepository
	passwordManager manager.PasswordManager
	tokenManager manager.TokenManager
	logger grpclog.LoggerV2
}

func (u *loginUsecaseImplementation) Call(ctx context.Context, params *LoginUsecaseParams) (*LoginUsecaseResult, error) {
	user, err := u.userRepository.FindBy(ctx, &repository.UserFilterArgs{Email: params.Email})
	if err != nil {
		u.logger.Errorf("Failed to get user: %v", err)
		return nil, exception.Internal(err.Error())
	}
	if user == nil {
		return nil, exception.NotFound("User is not found. It might be deleted.")
	}

	if !u.passwordManager.CheckHashPassword(params.Password, user.Password) {
		return nil, exception.Invalid("Password is not correct. Please pass a valid password for given email.")
	}

	token, err := u.tokenManager.CreateToken(ctx, user)
	if err != nil {
		u.logger.Errorf("Failed to create token: %v", err)
		return nil, exception.Internal(err.Error())
	}

	return &LoginUsecaseResult{Token: token}, nil
}

func NewLoginUsecase(
	userRepository repository.UserRepository,
	passwordManager manager.PasswordManager,
	tokenManager manager.TokenManager,
	logger grpclog.LoggerV2,
) LoginUsecase {
	return &loginUsecaseImplementation{
		userRepository: userRepository,
		passwordManager: passwordManager,
		tokenManager: tokenManager,
		logger: logger,
	}
}