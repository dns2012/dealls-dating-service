package mysql

import (
	"context"
	"errors"
	"fmt"
	"github.com/dns2012/dealls-dating-service/app/domain/entity"
	"github.com/dns2012/dealls-dating-service/app/domain/repository"
	"google.golang.org/grpc/grpclog"
	"gorm.io/gorm"
)

type MysqlPreferenceRepository struct {
	db *gorm.DB
	logger grpclog.LoggerV2
}

func (r *MysqlPreferenceRepository) ListBy(ctx context.Context, args *repository.PreferenceFilterArgs) ([]*entity.Preference, error) {
	query := r.db.WithContext(ctx)

	var preferences []*entity.Preference

	if args.UserID > 0 {
		query = query.Where("user_id = ?", args.UserID)
	}

	if len(args.CreatedAt) > 0 {
		query = query.Where("created_at BETWEEN ? AND ?", fmt.Sprintf("%s 00:00:01", args.CreatedAt), fmt.Sprintf("%s 23:59:59", args.CreatedAt))
	}

	if err := query.Find(&preferences).Error; err != nil {
		return nil, err
	}
	return preferences, nil
}

func (r *MysqlPreferenceRepository) Create(ctx context.Context, preference *entity.Preference) (*entity.Preference, error) {
	tx := r.db.WithContext(ctx).Begin()

	if err := tx.Create(&preference).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return preference, nil
}

func (r *MysqlPreferenceRepository) Count(ctx context.Context, args *repository.PreferenceFilterArgs) (int64, error) {
	query := r.db.WithContext(ctx)

	var preference int64

	if args.UserID > 0 {
		query = query.Where("user_id = ?", args.UserID)
	}

	if len(args.CreatedAt) > 0 {
		query = query.Where("created_at BETWEEN ? AND ?", fmt.Sprintf("%s 00:00:01", args.CreatedAt), fmt.Sprintf("%s 23:59:59", args.CreatedAt))
	}

	if err := query.Model(&entity.Preference{}).Count(&preference).Error; err != nil {
		return 0, err
	}
	return preference, nil
}


func (r *MysqlPreferenceRepository) FindBy(ctx context.Context, args *repository.PreferenceFilterArgs) (*entity.Preference, error) {
	query := r.db.WithContext(ctx)

	var preference *entity.Preference

	if args.UserID > 0 {
		query = query.Where("user_id = ?", args.UserID)
	}

	if args.PreferenceUserID > 0 {
		query = query.Where("preference_user_id = ?", args.PreferenceUserID)
	}

	if len(args.CreatedAt) > 0 {
		query = query.Where("created_at BETWEEN ? AND ?", fmt.Sprintf("%s 00:00:01", args.CreatedAt), fmt.Sprintf("%s 23:59:59", args.CreatedAt))
	}

	if err := query.First(&preference).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return preference, nil
}

func NewMysqlPreferenceRepository(
	db *gorm.DB,
	logger grpclog.LoggerV2,
) repository.PreferenceRepository {
	return &MysqlPreferenceRepository{
		db:     db,
		logger: logger,
	}
}