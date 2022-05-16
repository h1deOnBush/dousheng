package model

import "time"

type Comment struct {
	Id          int64     `json:"id,omitempty"`
	UserId      int64     `json:"user_id,omitempty"`
	VideoId     int64     `json:"video_id"`
	CommentText string    `json:"comment_text"`
	CreatedOn   time.Time `json:"created_on"`
}

func (c Comment) TableName() string {
	return "comment"
}
