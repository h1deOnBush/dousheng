package dao

import (
	"github.com/jinzhu/gorm"
	"github/h1deOnBush/dousheng/internal/model"
)

// 点赞视频, 点赞前需要先查看是否已经点赞
func (d *Dao) CreateFavorite(userId, videoId int64) error {
	favorite := model.Favorite{
		UserId:  userId,
		VideoId: videoId,
	}
	res := d.db.Create(&favorite)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// 取消点赞一个视频
func (d *Dao) DeleteFavorite(userId, videoId int64) error {
	var favorite model.Favorite
	res := d.db.Where("user_id=? and video_id=?", userId, videoId).Delete(&favorite)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// 获取一条点赞记录让service判断是否点赞过
func (d *Dao) GetFavorite(userId, videoId int64) (*model.Favorite, error) {
	var favorite model.Favorite
	res := d.db.Where("user_id=? and video_id=?", userId, videoId).First(&favorite)
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		return nil, res.Error
	}
	return &favorite, nil
}

// 查看指定用户点赞的视频集合
func (d *Dao) GetFavoriteVideoIdList(userId int64) ([]int64, error) {
	var videoIdList []int64
	var favoriteList []model.Favorite
	res := d.db.Table(model.Favorite{}.TableName()).Select("video_id").Where("user_id=?", userId).Find(&favoriteList)
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		return nil, res.Error
	}
	for _, f := range favoriteList {
		videoIdList = append(videoIdList, f.VideoId)
	}
	return videoIdList, nil
}
