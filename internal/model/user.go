package model

type User struct {
	Id            int64  `json:"id,omitempty"`
	Username          string `json:"name,omitempty"`
	Password	  string `json:"password,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

func (u User) TableName() string {
	return "user"
}