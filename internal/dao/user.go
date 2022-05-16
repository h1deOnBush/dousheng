package dao

import (
	"github.com/jinzhu/gorm"
	"github/h1deOnBush/dousheng/internal/model"
)

// 注册用户
func (d *Dao) CreateUser(username, password string) (int64, error) {
	user := model.User{Username: username, Password: password, FollowCount: 0, FollowerCount: 0}
	res := d.db.Create(&user)
	if res.Error != nil {
		return 0, res.Error
	}
	return user.Id, nil
}

// 根据uid查找user信息
func (d *Dao) GetUserById(uid int64) (*model.User, error) {
	user := model.User{Id: uid}

	result := d.db.First(&user)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return &user, result.Error
	}
	return &user, nil
}

// 根据username查询user
func (d *Dao) GetUserByName(username string) (*model.User, error) {
	var user model.User
	result := d.db.Model(&model.User{}).Where("username=?", username).First(&user)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	return &user, nil
}

// 根据id列表批量返回用户
func (d *Dao) GetUserListById(userIdList []int64) ([]*model.User, error) {
	userList := []*model.User{}
	result := d.db.Where(userIdList).Find(&userList)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	return userList, nil
}

// 修改指定用户的关注数
func (d *Dao) UpdateUserFollowCount(userId int64, cnt int) error {
	user := model.User{Id: userId}
	err := d.db.Model(&user).Update("follow_count", gorm.Expr("follow_count + ?", cnt)).Error
	return err
}

// 修改指定用户的粉丝数
func (d *Dao) UpdateUserFollowerCount(userId int64, cnt int) error {
	user := model.User{Id: userId}
	err := d.db.Model(&user).Update("follower_count", gorm.Expr("follower_count  + ? ", cnt)).Error
	return err
}
