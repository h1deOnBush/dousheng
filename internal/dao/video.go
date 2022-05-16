package dao

import (
	"github/h1deOnBush/dousheng/internal/model"
	"time"
)

// 上传一条视频
func (d *Dao) CreateVideo() error {
	return nil
}

// 根据时间返回视频列表
func (d *Dao) GetVideoListByTime(latestTime time.Time) ([]*model.Video, error) {
	return nil, nil
}

// 根据投稿者返回视频列表
func (d *Dao) GetVideoListByAuthor(authorId int64) ([]*model.Video, error) {
	return nil, nil
}

// 更新视频的点赞数
func (d *Dao) UpdateVideoFavoriteCount(videoId int64, cnt int) error {
	return nil
}

// 更新视频的评论数
func (d *Dao) UpdateVideoCommentCount(videoId int64, cnt int) error {
	return nil
}
