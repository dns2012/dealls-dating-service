package repository

import (
	"context"
	"github.com/dns2012/dealls-dating-service/app/domain/entity"
)

type UserPackageRepository interface {
	Create(ctx context.Context, userPackage *entity.UserPackage) (*entity.UserPackage, error)
}