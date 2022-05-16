package model

import "time"

type Video struct {
	Id            int64     `json:"id,omitempty"`
	AuthorId      int64     `json:"author_id,omitempty"`
	PlayUrl       string    `json:"play_url,omitempty"`
	CreatedOn     time.Time `json:"created_on,omitempty"`
	CoverUrl      string    `json:"cover_url,omitempty"`
	FavoriteCount int64     `json:"favorite_count,omitempty"`
	CommentCount  int64     `json:"comment_count,omitempty"`
}

func (v Video) TableName() string {
	return "video"
}
