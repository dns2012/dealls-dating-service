package repository

import (
	"context"
	"github.com/dns2012/dealls-dating-service/app/domain/entity"
)

type PreferenceRepository interface {
	ListBy(ctx context.Context, args *PreferenceFilterArgs) ([]*entity.Preference, error)
	Create(ctx context.Context, preference *entity.Preference) (*entity.Preference, error)
	Count(ctx context.Context, args *PreferenceFilterArgs) (int64, error)
	FindBy(ctx context.Context, args *PreferenceFilterArgs) (*entity.Preference, error)
}

type PreferenceFilterArgs struct {
	UserID uint `json:"user_id" db:"user_id"`
	PreferenceUserID uint `json:"preference_user_id" db:"preference_user_id"`
	CreatedAt string `json:"created_at" db:"created_at"`
}