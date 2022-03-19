package base

import (
	"context"
	"gorm.io/gorm"
)

// Repository prepare for domain
type Repository struct {
	ctx context.Context
	db  *gorm.DB
}

func NewRepository(db *gorm.DB, dstStruct interface{}) *Repository {
	ctx := context.Background()
	return &Repository{
		ctx: ctx,
		db:  db.WithContext(ctx).Model(dstStruct),
	}
}

func (r *Repository) GetDB() *gorm.DB {
	return r.db
}

func (r *Repository) SetRepository(model interface{}) *Repository {
	ctx := context.Background()
	return &Repository{
		ctx: ctx,
		db:  r.db.WithContext(ctx).Model(model),
	}
}

// ******************** 新增修改函数 ******************//

func (r *Repository) Save(data interface{}) error {
	return r.db.Save(data).Error
}

// ******************** 删除查询函数 ******************//

// First 根据条件查询第一条记录
func (r *Repository) First(dst interface{}, filter interface{}) error {
	return r.db.Where(filter).First(dst).Error
}

// Take 随机获取一条记录
func (r *Repository) Take(dst interface{}, filter interface{}) error {
	return r.db.Where(filter).Take(dst).Error
}

// Last 随机获取一条记录
func (r *Repository) Last(dst interface{}, filter interface{}) error {
	return r.db.Where(filter).Last(dst).Error
}

// Find 根据条件获取全部
func (r *Repository) Find(dst interface{}, filter ...interface{}) error {
	return r.db.Where(filter).Find(dst).Error
}

// List 根据条件获取列表
func (r *Repository) List(dst interface{}, pageSize int64, page int64, filter ...interface{}) error {
	if len(filter) >= 2 {
		return r.db.Where(filter[0], filter[1:]...).Offset(int(page) - 1).Limit(int(pageSize)).Find(dst).Error
	} else {
		return r.db.Offset(int(page) - 1).Limit(int(pageSize)).Find(dst).Error
	}
}

// Counts 根据条件获取列表
func (r *Repository) Counts(dst *int64, filter ...interface{}) error {
	if len(filter) >= 2 {
		return r.db.Where(filter[0], filter[1:]...).Count(dst).Error
	} else {
		return r.db.Count(dst).Error
	}
}

// ******************** 删除相关函数 ******************//

func (r *Repository) Delete(dst interface{}, filter ...interface{}) error {
	return r.db.Where(filter).Delete(dst).Error
}
