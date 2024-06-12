package repository

import (
	"context"
	"github.com/dns2012/dealls-dating-service/app/domain/entity"
	"github.com/dns2012/dealls-dating-service/app/domain/manager"
)

type UserRepository interface {
	FindBy(ctx context.Context, args *UserFilterArgs) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	ListBy(ctx context.Context, args *UserFilterArgs, pagination manager.Pagination) (*manager.Pagination, []*entity.User, error)
}

type UserFilterArgs struct {
	Email string `db:"email"`
	UserID uint `db:"user_id"`
	ExcludeUserID []uint
}