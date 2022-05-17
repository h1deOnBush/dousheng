package service

import (
	"crypto/md5"
	"fmt"
	"github/h1deOnBush/dousheng/pkg/errcode"
	"io"
)

// service层的user结构体设计
type User struct {
	Id            int64  `json:"id"`
	Username      string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

// 用户注册
func (svc *Service) Register(username, password string) (int64, error) {
	// 先查询是否有同名用户
	user, err := svc.dao.GetUserByName(username)
	if err != nil {
		return 0, err
	}
	if user.Id > 0 {
		return 0, errcode.UsernameAlreadyExists
	}
	// 将密码进行md5加密后进行存储
	h := md5.New()
	if _, err = io.WriteString(h, password); err != nil {
		return 0, err
	}
	password = fmt.Sprintf("%x", h.Sum(nil))
	userId, err := svc.dao.CreateUser(username, password)
	return userId, err
}

// 用户认证, 返回的userId用于做jwt加密
func (svc *Service) CheckUser(username, password string) (int64, error) {
	user, err := svc.dao.GetUserByName(username)
	if err != nil {
		return 0, err
	}
	if user.Id == 0 {
		return 0, errcode.UserNotExists
	}

	// 密码进行加密后再比对
	h := md5.New()
	if _, err = io.WriteString(h, password); err != nil {
		return 0, err
	}
	password = fmt.Sprintf("%x", h.Sum(nil))

	if user.Password != password {
		return 0, errcode.UserPasswordWrong
	}

	return user.Id, nil
}

// 获取用户信息
func (svc *Service) GetUserInfo(userId, myId int64) (*User, error) {
	user, err := svc.dao.GetUserById(userId)

	if err != nil {
		return nil, err
	}
	if user.Username == "" {
		return nil, errcode.UserNotExists
	}
	u := &User{
		Id:            user.Id,
		Username:      user.Username,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      false,
	}
	// 查看自己是否关注该用户, 如果是自己查询自己直接返回。如果myId==0代表没有登录
	if userId == myId || myId == 0 {
		return u, nil
	}
	r, err := svc.dao.GetRelation(userId, myId)
	if err != nil {
		return nil, err
	}
	if r.Id != 0 {
		u.IsFollow = true
	}
	return u, nil
}
