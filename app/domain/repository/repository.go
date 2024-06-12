package repository

import (
	"github.com/dns2012/dealls-dating-service/app/domain/manager"
	"gorm.io/gorm"
	"math"
)

func Paginate(value interface{}, pagination *manager.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var rowsCount int64
	db.Model(value).Count(&rowsCount)

	pagination.RowsCount = rowsCount
	pageCount := int(math.Ceil(float64(rowsCount) / float64(pagination.PageSize)))
	pagination.PageCount = pageCount

	pagination.Prev = pagination.GetPrevPage()
	pagination.Next = pagination.GetNextPage()

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetPageSize())
	}
}
