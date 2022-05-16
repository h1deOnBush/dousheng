package errcode

// 用户错误码
var (
	UsernameAlreadyExists = NewError(20000000, "用户名已存在")
	UserNotExists         = NewError(20000001, "用户不存在")
	UserPasswordWrong     = NewError(20000002, "用户密码不存在")
	UserRegisterFail      = NewError(20000003, "用户注册失败")
	UserLoginFail         = NewError(20000004, "用户登录失败")
	GetUserInfoFail       = NewError(20000005, "获取用户信息失败")
)

var (
	ErrorUploadFileFail = NewError(30000000, "视频上传失败")
	GetPublishListFail  = NewError(30000001, "获取发布失败")
	GetFeedFail         = NewError(30000002, "获取Feed流失败")
)

var (
	LikeFail            = NewError(40000000, "点赞失败")
	UnlikeFail          = NewError(40000001, "取消点赞失败")
	GetFavoriteListFail = NewError(40000002, "获取点赞列表失败")
)

var (
	FollowFail          = NewError(50000000, "关注失败")
	UnFollowFail        = NewError(50000001, "取消关注失败")
	GetFollowListFail   = NewError(50000002, "获取关注列表失败")
	GetFollowerListFail = NewError(50000003, "获取粉丝列表失败")
)

var (
	CommentFail        = NewError(60000000, "评论失败")
	DeleteCommentFail  = NewError(60000001, "删除评论失败")
	GetCommentListFail = NewError(60000002, "获取评论列表失败")
)
