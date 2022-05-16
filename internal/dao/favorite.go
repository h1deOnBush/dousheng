package dao

// 点赞视频
func (d *Dao) CreateFavorite(userId, videoId int64) error {
	return nil
}

// 取消点赞一个视频
func (d *Dao) DeleteFavorite(userId, videoId int64) error {
	return nil
}

// 获取一条点赞记录让service判断是否点赞过
func (d *Dao) GetFavorite(userId, videoId int64) error {
	return nil
}

// 查看指定用户点赞的视频集合
func (d *Dao) GetFavoriteVideoIdList(userId int64) ([]int64, error) {
	return []int64{}, nil
}
