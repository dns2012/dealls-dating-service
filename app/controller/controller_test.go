package controller_test

import (
	"github.com/dns2012/dealls-dating-service/mocks"
)

type SchemaV1APIMock struct {
	loginUsecase mocks.LoginUsecase
	registerUsecase mocks.RegisterUsecase
	listUserUsecase mocks.ListUserUsecase
	createUserPreferenceUsecase mocks.CreateUserPreferenceUsecase
	listPackageUsecase mocks.ListPackageUsecase
	orderPackageUsecase mocks.OrderPackageUsecase
	meUsecase mocks.MeUsecase

	tokenManager mocks.TokenManager
}