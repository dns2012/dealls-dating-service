package mysql

import (
	"context"
	"errors"
	"github.com/dns2012/dealls-dating-service/app/domain/entity"
	"github.com/dns2012/dealls-dating-service/app/domain/manager"
	"github.com/dns2012/dealls-dating-service/app/domain/repository"
	"google.golang.org/grpc/grpclog"
	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	db *gorm.DB
	logger grpclog.LoggerV2
}

func (r *MysqlUserRepository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	tx := r.db.WithContext(ctx).Begin()

	if err := tx.Omit("premium_at").Create(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *MysqlUserRepository) FindBy(ctx context.Context, args *repository.UserFilterArgs) (*entity.User, error) {
	query := r.db.WithContext(ctx)

	var user *entity.User

	if args.UserID > 0 {
		query = query.Where("id = ?", args.UserID)
	}

	if len(args.Email) > 0 {
		query = query.Where("email = ?", args.Email)
	}

	if err := query.Preload("Profile").Preload("UserPackage.Package").First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *MysqlUserRepository) ListBy(ctx context.Context, args *repository.UserFilterArgs, pagination manager.Pagination) (*manager.Pagination, []*entity.User, error) {
	var users []*entity.User

	query := r.db.WithContext(ctx)

	if len(args.ExcludeUserID) > 0 {
		query = query.Where("id NOT IN (?)", args.ExcludeUserID)
	}

	if err := query.Scopes(repository.Paginate(users, &pagination, query)).Preload("Profile").Find(&users).Error; err != nil {
		return nil, nil, err
	}

	return &pagination, users, nil
}

func NewMysqlUserRepository(
	db *gorm.DB,
	logger grpclog.LoggerV2,
) repository.UserRepository {
	return &MysqlUserRepository{
		db:     db,
		logger: logger,
	}
}