package model

type Favorite struct {
	Id      int64 `json:"id"`
	UserId  int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
}

func (r Favorite) TableName() string {
	return "favorite"
}
