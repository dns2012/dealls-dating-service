package provider

import (
	"github.com/dns2012/dealls-dating-service/app/domain/manager"
	"github.com/dns2012/dealls-dating-service/config"
	"golang.org/x/crypto/bcrypt"
)

var (
	tokenManager = manager.NewTokenManager([]byte(config.ENV("JWT_SECRET_KEY", "secret")))
	passwordManager = manager.NewPasswordManager(bcrypt.DefaultCost)
)
