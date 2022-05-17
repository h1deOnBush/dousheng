package service

// 点赞
func (svc *Service) Like(userId, videoId int64) error {
	err := svc.dao.CreateFavorite(userId, videoId)
	if err != nil {
		return err
	}
	// 更新视频的点赞数
	err = svc.dao.UpdateVideoFavoriteCount(videoId, 1)
	return err
}

// 取消点赞
func (svc *Service) Unlike(userId, videoId int64) error {
	err := svc.dao.DeleteFavorite(userId, videoId)
	if err != nil {
		return err
	}
	// 更新视频的点赞数
	err = svc.dao.UpdateVideoFavoriteCount(videoId, -1)
	return err
}

// 获取用户点赞列表
func (svc *Service) GetFavoriteList(userId int64) ([]*Video, error) {
	videoList := []*Video{}
	videoIdList, err := svc.dao.GetFavoriteVideoIdList(userId)
	if err != nil {
		return nil, err
	}

	// 根据点赞的视频id获取视频实体集合
	videos, err := svc.dao.GetVideoListById(videoIdList)
	if err != nil {
		return nil, err
	}

	// 封装成想要返回的信息
	for _, v := range videos {
		// 获取视频作者信息
		user, err := svc.GetUserInfo(v.AuthorId, userId)
		if err != nil {
			return nil, err
		}
		vv := &Video{
			Id:            v.Id,
			Title:         v.Title,
			Author:        *user,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    true,
		}
		videoList = append(videoList, vv)
	}

	return videoList, nil
}
