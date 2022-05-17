package dao

import (
	"github.com/jinzhu/gorm"
	"github/h1deOnBush/dousheng/internal/model"
	"time"
)

// 上传一条视频
func (d *Dao) CreateVideo(authorId int64, title, playUrl, coverUrl string) error {
	video := model.Video{
		Title:         title,
		AuthorId:      authorId,
		PlayUrl:       playUrl,
		CreatedOn:     time.Now(),
		CoverUrl:      coverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
	}
	res := d.db.Create(&video)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// 根据时间返回视频列表
func (d *Dao) GetVideoListByTimeDesc(latestTime time.Time) ([]*model.Video, error) {
	videoList := []*model.Video{}
	result := d.db.Model(&model.Video{}).Where("created_on < ?", latestTime).Order("created_on DESC").Limit(30).Find(&videoList)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	return videoList, nil
}

// 根据投稿者返回视频列表
func (d *Dao) GetVideoListByAuthor(authorId int64) ([]*model.Video, error) {
	videoList := []*model.Video{}
	result := d.db.Where("author_id=?", authorId).Find(&videoList)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	return videoList, nil
}

// 更新视频的点赞数
func (d *Dao) UpdateVideoFavoriteCount(videoId int64, cnt int) error {
	video := model.Video{Id: videoId}
	err := d.db.Model(&video).Update("favorite_count", gorm.Expr("favorite_count + ?", cnt)).Error
	return err
}

// 更新视频的评论数
func (d *Dao) UpdateVideoCommentCount(videoId int64, cnt int) error {
	video := model.Video{Id: videoId}
	err := d.db.Model(&video).Update("comment_count", gorm.Expr("comment_count + ?", cnt)).Error
	return err
}

func (d *Dao) GetVideoListById(videoIdList []int64) ([]*model.Video, error) {
	videoList := []*model.Video{}
	result := d.db.Where(videoIdList).Find(&videoList)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	return videoList, nil
}
