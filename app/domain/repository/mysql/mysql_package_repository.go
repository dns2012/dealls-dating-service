package mysql

import (
	"context"
	"errors"
	"github.com/dns2012/dealls-dating-service/app/domain/entity"
	"github.com/dns2012/dealls-dating-service/app/domain/repository"
	"google.golang.org/grpc/grpclog"
	"gorm.io/gorm"
)

type MysqlPackageRepository struct {
	db *gorm.DB
	logger grpclog.LoggerV2
}

func (r *MysqlPackageRepository) List(ctx context.Context) ([]*entity.Package, error) {
	query := r.db.WithContext(ctx)

	var packages []*entity.Package

	if err := query.Find(&packages).Error; err != nil {
		return nil, err
	}

	return packages, nil
}

func (r *MysqlPackageRepository) Find(ctx context.Context, id uint) (*entity.Package, error) {
	query := r.db.WithContext(ctx)

	var pkg *entity.Package

	if err := query.First(&pkg, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return pkg, nil
}

func NewMysqlPackageRepository(
	db *gorm.DB,
	logger grpclog.LoggerV2,
) repository.PackageRepository {
	return &MysqlPackageRepository{
		db:     db,
		logger: logger,
	}
}