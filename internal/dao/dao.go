package dao

import (
	"github.com/jinzhu/gorm"
)

type Dao struct {
	db *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{db: engine}
}
