package controller

import (
	"context"
	global "github/h1deOnBush/dousheng/gloabal"
	"github/h1deOnBush/dousheng/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
