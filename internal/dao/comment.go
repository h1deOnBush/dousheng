package dao

import (
	"github/h1deOnBush/dousheng/internal/model"
	"time"
)

// 新增一条评论
func (d *Dao) CreateComment(userId, videoId int64, commentText string, createdOn time.Time) error {
	return nil
}

// 删除一条评论
func (d *Dao) DeleteComment(commentId int64) error {
	return nil
}

// 查询某个视频的所有评论，根据发布时间排序
func (d *Dao) GetCommentListByVideoIdOrderByCreatedOn(videoId int64) ([]*model.Comment, error) {
	return nil, nil
}
