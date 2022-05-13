CREATE TABLE IF NOT EXISTS `user` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(100) NOT NULL COMMENT '用户名',
    `password` varchar(100) NOT NULL COMMENT '用户密码',
    `follow_count` bigint(20) NOT NULL COMMENT '关注数',
    `follower_count` bigint(20) NOT NULL COMMENT '粉丝数',
    `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `is_follow` tinyint(3) unsigned DEFAULT '0' COMMENT '是否关注 0为未关注、1为已关注',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

CREATE TABLE IF NOT EXISTS `video` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `author_id` bigint(20) unsigned NOT NULL COMMENT '发布者用户id',
    `play_url` varchar(255) NOT NULL COMMENT '视频播放网址',
    `cover_url` varchar(255) NOT NULL COMMENT '视频封面网址',
    `favorite_count` bigint(20) NOT NULL COMMENT '被点赞数',
    `comment_count` bigint(20) NOT NULL COMMENT '评论数',
    `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
    `is_favorite` tinyint(3) unsigned DEFAULT '0' COMMENT '是否点赞 0为未点赞、1为已点赞',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='视频表';