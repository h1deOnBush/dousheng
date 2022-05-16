package dao

import (
	"github.com/jinzhu/gorm"
	"github/h1deOnBush/dousheng/internal/model"
	"time"
)

// 新增一条评论
func (d *Dao) CreateComment(userId, videoId int64, commentText string) error {
	comment := model.Comment{
		UserId:      userId,
		VideoId:     videoId,
		CommentText: commentText,
		CreatedOn:   time.Now(),
	}
	res := d.db.Create(&comment)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// 删除一条评论
func (d *Dao) DeleteComment(commentId int64) error {
	var comment model.Comment
	res := d.db.Where("id=?", commentId).Delete(&comment)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// 查询某个视频的所有评论，根据发布时间排序
func (d *Dao) GetCommentListByVideoIdOrderByTimeDesc(videoId int64) ([]*model.Comment, error) {
	commentList := []*model.Comment{}
	result := d.db.Model(&model.Video{}).Where("video_id=?", videoId).Order("created_on DESC").Find(&commentList)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	return commentList, nil
}
