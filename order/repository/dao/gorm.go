package dao

import (
	"gorm.io/gorm"
)

type GORMArticleDao struct {
	db *gorm.DB
}

func NewGORMOrderDao(db *gorm.DB) OrderDao {
	return &GORMArticleDao{
		db: db,
	}
}
