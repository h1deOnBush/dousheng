package model

type Relation struct {
	Id         int64 `json:"id"`
	UserId     int64 `json:"user_id"`
	FollowerId int64 `json:"follower_id"`
}

func (r Relation) TableName() string {
	return "relation"
}
