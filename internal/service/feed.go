package service

import "time"

func (svc *Service) Feed(latestTime time.Time, UserId int64) ([]*Video, int64, error) {
	videoList := []*Video{}

	videos, err := svc.dao.GetVideoListByTime(latestTime)
	if err != nil {
		return nil, 0, err
	}
	var nextTime int64
	if len(videos) > 0 {
		for _, v := range videos {
			user, err := svc.GetUserInfo(v.AuthorId, UserId)
			vv := &Video{
				Id:            v.Id,
				Author:        *user,
				PlayUrl:       v.PlayUrl,
				CoverUrl:      v.CoverUrl,
				FavoriteCount: v.FavoriteCount,
				CommentCount:  v.CommentCount,
				IsFavorite:    false,
			}

			// 在登录情况下查询自己是否点赞过该作品
			f, err := svc.dao.GetFavorite(v.Id, UserId)
			if err != nil {
				return nil, 0, err
			}
			if f.Id != 0 {
				vv.IsFavorite = true
			}
			videoList = append(videoList, vv)
		}
		nextTime = videos[len(videos)-1].CreatedOn.Unix()
	}
	return videoList, nextTime, nil
}
