package service

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	Title         string `json:"title"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
}

func (svc *Service) Publish(authorId int64, title, playUrl, coverUrl string) error {
	err := svc.dao.CreateVideo(authorId, title, playUrl, coverUrl)
	if err != nil {
		return err
	}
	return nil
}

// 查询指定用户的发布列表
func (svc *Service) PublishList(authorId int64) ([]*Video, error) {
	videoList := []*Video{}
	user, err := svc.GetUserInfo(authorId, authorId)
	if err != nil {
		return nil, err
	}

	// 查询指定用户发布的所有视频
	videos, err := svc.dao.GetVideoListByAuthor(authorId)
	if err != nil {
		return nil, err
	}
	if len(videos) != 0 {
		// 封装成想要返回的信息
		for _, v := range videos {
			vv := &Video{
				Id:            v.Id,
				Title:         v.Title,
				Author:        *user,
				PlayUrl:       v.PlayUrl,
				CoverUrl:      v.CoverUrl,
				FavoriteCount: v.FavoriteCount,
				CommentCount:  v.CommentCount,
				IsFavorite:    false,
			}
			// 查询自己是否点赞过自己的视频
			f, err := svc.dao.GetFavorite(authorId, v.Id)
			if err != nil {
				return nil, err
			}
			if f.Id != 0 {
				vv.IsFavorite = true
			}
			videoList = append(videoList, vv)
		}
	}
	return videoList, nil
}
