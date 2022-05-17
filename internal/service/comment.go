package service

type Comment struct {
	Id         int64  `json:"id"`
	User       User   `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

func (svc *Service) CommentOn(userId, videoId int64, commentText string) error {
	err := svc.dao.CreateComment(userId, videoId, commentText)
	if err != nil {
		return nil
	}
	// 更新对应视频的评论数
	err = svc.dao.UpdateVideoCommentCount(videoId, 1)
	return err
}

func (svc *Service) DeleteComment(commentId int64, videoId int64) error {
	err := svc.dao.DeleteComment(commentId)
	if err != nil {
		return nil
	}
	// 更新对应视频的评论数
	err = svc.dao.UpdateVideoCommentCount(videoId, -1)
	return err
}

func (svc *Service) GetCommentList(userId, videoId int64) ([]*Comment, error) {
	commentList := []*Comment{}
	comments, err := svc.dao.GetCommentListByVideoIdOrderByTimeDesc(videoId)
	if err != nil {
		return nil, err
	}
	for _, c := range comments {
		user, err := svc.GetUserInfo(c.UserId, userId)
		if err != nil {
			return nil, err
		}
		commentList = append(commentList, &Comment{
			Id:         c.Id,
			User:       *user,
			Content:    c.CommentText,
			CreateDate: c.CreatedOn.Format("2006-01-02 15:04:05"),
		})
	}
	return commentList, nil
}
