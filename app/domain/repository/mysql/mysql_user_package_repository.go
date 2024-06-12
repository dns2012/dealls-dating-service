package mysql

import (
	"context"
	"github.com/dns2012/dealls-dating-service/app/domain/entity"
	"github.com/dns2012/dealls-dating-service/app/domain/repository"
	"google.golang.org/grpc/grpclog"
	"gorm.io/gorm"
	"time"
)

type MysqlUserPackageRepository struct {
	db *gorm.DB
	logger grpclog.LoggerV2
}

func (r *MysqlUserPackageRepository) Create(ctx context.Context, userPackage *entity.UserPackage) (*entity.UserPackage, error) {
	tx := r.db.WithContext(ctx).Begin()

	if err := tx.Delete(&entity.UserPackage{}, "user_id = ?", userPackage.UserID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Create(&userPackage).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var user *entity.User

	if err := tx.First(&user, userPackage.UserID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if user.PremiumAt.IsZero() {
		user.PremiumAt = time.Now()
		if err := tx.Save(&user).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return userPackage, nil
}

func NewMysqlUserPackageRepository(
	db *gorm.DB,
	logger grpclog.LoggerV2,
) repository.UserPackageRepository {
	return &MysqlUserPackageRepository{
		db:     db,
		logger: logger,
	}
}