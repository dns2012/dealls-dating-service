package usecase_test

import (
	"github.com/dns2012/dealls-dating-service/mocks"
	"google.golang.org/grpc/grpclog"
)

type UsecaseMock struct {
	userRepository mocks.UserRepository
	preferenceRepository mocks.PreferenceRepository
	packageRepository mocks.PackageRepository
	userPackageRepository mocks.UserPackageRepository

	passwordManager mocks.PasswordManager
	tokenManager mocks.TokenManager

	logger grpclog.LoggerV2
}
