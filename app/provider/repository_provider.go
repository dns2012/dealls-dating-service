package provider

import (
	"github.com/dns2012/dealls-dating-service/app/domain/repository/mysql"
	"github.com/dns2012/dealls-dating-service/config"
	"google.golang.org/grpc/grpclog"
	"io/ioutil"
	"os"
)

var (
	logger = grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	userRepository        = mysql.NewMysqlUserRepository(config.GormDB(), logger)
	preferenceRepository  = mysql.NewMysqlPreferenceRepository(config.GormDB(), logger)
	packageRepository     = mysql.NewMysqlPackageRepository(config.GormDB(), logger)
	userPackageRepository = mysql.NewMysqlUserPackageRepository(config.GormDB(), logger)
)
