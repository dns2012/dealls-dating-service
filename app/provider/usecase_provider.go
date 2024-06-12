package provider

import (
	"github.com/dns2012/dealls-dating-service/app/controller"
	"github.com/dns2012/dealls-dating-service/app/usecase"
)

var (
	loginUsecase = usecase.NewLoginUsecase(userRepository, passwordManager, tokenManager, logger)
	registerUsecase = usecase.NewRegisterUsecase(userRepository, passwordManager, tokenManager,logger)
	listUserUsecase = usecase.NewListUserUsecase(userRepository, preferenceRepository, logger)
	createUserPreferenceUsecase = usecase.NewCreateUserPreferenceUsecase(userRepository, preferenceRepository, logger)
	listPackageUsecase = usecase.NewListPackageUsecase(packageRepository, logger)
	orderPackageUsecase = usecase.NewOrderPackageUsecase(userRepository, packageRepository, userPackageRepository, logger)
	meUsecase = usecase.NewMeUsecase(userRepository, logger)
)

func SchemaV1UsecaseProvider() *controller.SchemaV1API {
	return controller.NewSchemaV1API(
		loginUsecase,
		registerUsecase,
		listUserUsecase,
		createUserPreferenceUsecase,
		listPackageUsecase,
		orderPackageUsecase,
		meUsecase,
		tokenManager,
	)
}