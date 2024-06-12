package repository

import (
	"context"
	"github.com/dns2012/dealls-dating-service/app/domain/entity"
)

type PackageRepository interface {
	List(ctx context.Context) ([]*entity.Package, error)
	Find(ctx context.Context, id uint) (*entity.Package, error)
}