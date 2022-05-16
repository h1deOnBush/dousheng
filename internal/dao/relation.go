package dao

import "github/h1deOnBush/dousheng/internal/model"

// 关注一个用户
func (d *Dao) CreateRelation(userId, toUserId int64) error {
	return nil
}

// 取关一个用户
func (d *Dao) DeleteRelation(userId, toUserId int64) error {
	return nil
}

// 获取一条关系记录让service判断是否关注过
func (d *Dao) GetRelation(userId, followerId int64) (*model.Relation, error) {
	return nil, nil
}

// 查看指定用户关注的用户id集合
func (d *Dao) GetFollowUserIdList(userId int64) ([]int64, error) {
	return []int64{}, nil
}

// 查看指定用户粉丝用户id集合
func (d *Dao) GetFollowerUserIdList(userId int64) ([]int64, error) {
	return []int64{}, nil
}
