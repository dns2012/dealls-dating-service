package controller

import (
	"github.com/dns2012/dealls-dating-service/app/domain/manager"
	"github.com/dns2012/dealls-dating-service/app/usecase"
	schemav1 "github.com/dns2012/dealls-dating-service/proto/schema/v1"
)

type SchemaV1API struct {
	schemav1.UnimplementedAuthSchemaServer
	schemav1.UnimplementedUserSchemaServer
	schemav1.UnimplementedPackageSchemaServer

	LoginUsecase                usecase.LoginUsecase
	RegisterUsecase             usecase.RegisterUsecase
	ListUserUsecase             usecase.ListUserUsecase
	CreateUserPreferenceUsecase usecase.CreateUserPreferenceUsecase
	ListPackageUsecase          usecase.ListPackageUsecase
	OrderPackageUsecase         usecase.OrderPackageUsecase
	MeUsecase                   usecase.MeUsecase

	TokenManager manager.TokenManager
}

func NewSchemaV1API(
	loginUsecase usecase.LoginUsecase,
	registerUsecase usecase.RegisterUsecase,
	listUserUsecase usecase.ListUserUsecase,
	createUserPreferenceUsecase usecase.CreateUserPreferenceUsecase,
	listPackageUsecase usecase.ListPackageUsecase,
	orderPackageUsecase usecase.OrderPackageUsecase,
	meUsecase usecase.MeUsecase,
	tokenManager manager.TokenManager,
) *SchemaV1API {
	return &SchemaV1API{
		LoginUsecase:                loginUsecase,
		RegisterUsecase:             registerUsecase,
		ListUserUsecase:             listUserUsecase,
		CreateUserPreferenceUsecase: createUserPreferenceUsecase,
		ListPackageUsecase:          listPackageUsecase,
		OrderPackageUsecase:         orderPackageUsecase,
		MeUsecase:                   meUsecase,
		TokenManager:                tokenManager,
	}
}