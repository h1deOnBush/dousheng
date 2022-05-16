package service

// 关注
func (svc *Service) Follow(userId, followerId int64) error {
	err := svc.dao.CreateRelation(userId, followerId)
	if err != nil {
		return err
	}
	// 更新2个用户的关注粉丝数
	err = svc.dao.UpdateUserFollowerCount(userId, 1)
	if err != nil {
		return err
	}
	err = svc.dao.UpdateUserFollowCount(followerId, 1)
	if err != nil {
		return err
	}
	return nil
}

// 取消关注
func (svc *Service) Unfollow(userId, followerId int64) error {
	err := svc.dao.DeleteRelation(userId, followerId)
	if err != nil {
		return err
	}
	// 更新2个用户的关注粉丝数
	err = svc.dao.UpdateUserFollowerCount(userId, -1)
	if err != nil {
		return err
	}
	err = svc.dao.UpdateUserFollowCount(followerId, -1)
	if err != nil {
		return err
	}
	return nil
}

// 获取关注列表
func (svc *Service) GetFollowList(userId int64) ([]*User, error) {
	userList := []*User{}
	followIdList, err := svc.dao.GetFollowUserIdList(userId)
	if err != nil {
		return nil, err
	}
	users, err := svc.dao.GetUserListById(followIdList)
	if err != nil {
		return nil, err
	}
	for _, u := range users {
		userList = append(userList, &User{
			Id:            u.Id,
			Username:      u.Username,
			FollowCount:   u.FollowCount,
			FollowerCount: u.FollowerCount,
			IsFollow:      true,
		})
	}
	return userList, nil
}

// 获取粉丝列表
func (svc *Service) GetFollowerList(userId int64) ([]*User, error) {
	userList := []*User{}
	followerIdList, err := svc.dao.GetFollowerUserIdList(userId)
	if err != nil {
		return nil, err
	}
	users, err := svc.dao.GetUserListById(followerIdList)
	if err != nil {
		return nil, err
	}
	for _, u := range users {
		// 反向查看自己是否关注了粉丝
		r, err := svc.dao.GetRelation(u.Id, userId)
		if err != nil {
			return nil, err
		}
		isFollow := false
		if r.Id > 0 {
			isFollow = true
		}
		userList = append(userList, &User{
			Id:            u.Id,
			Username:      u.Username,
			FollowCount:   u.FollowCount,
			FollowerCount: u.FollowerCount,
			IsFollow:      isFollow,
		})
	}
	return userList, nil
}
