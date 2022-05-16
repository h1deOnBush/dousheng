package dao

import (
	"github.com/jinzhu/gorm"
	"github/h1deOnBush/dousheng/internal/model"
)

// 关注一个用户
func (d *Dao) CreateRelation(userId, followerId int64) error {
	relation := model.Relation{
		UserId:     userId,
		FollowerId: followerId,
	}
	res := d.db.Create(&relation)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// 取关一个用户
func (d *Dao) DeleteRelation(userId, followerId int64) error {
	var relation model.Relation
	res := d.db.Where("user_id=? and follower_id=?", userId, followerId).Delete(&relation)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// 获取一条关系记录让service层判断是否关注过
func (d *Dao) GetRelation(userId, followerId int64) (*model.Relation, error) {
	var relation model.Relation
	res := d.db.Where("user_id=? and follower_id=?", userId, followerId).First(&relation)
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		return nil, res.Error
	}
	return &relation, nil
}

// 查看指定用户关注的用户id集合
func (d *Dao) GetFollowUserIdList(userId int64) ([]int64, error) {
	var userIdList []int64
	var relationList []model.Relation
	res := d.db.Select("user_id").Where("follower_id=?", userId).Find(&relationList)
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		return nil, res.Error
	}
	for _, r := range relationList {
		userIdList = append(userIdList, r.UserId)
	}
	return userIdList, nil
}

// 查看指定用户粉丝用户id集合
func (d *Dao) GetFollowerUserIdList(userId int64) ([]int64, error) {
	var userIdList []int64
	var relationList []model.Relation
	res := d.db.Select("follower_id").Where("user_id=?", userId).Find(&relationList)
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		return nil, res.Error
	}
	for _, r := range relationList {
		userIdList = append(userIdList, r.FollowerId)
	}
	return userIdList, nil
}
