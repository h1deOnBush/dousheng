CREATE TABLE IF NOT EXISTS  `user` (
    `id` BIGINT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(32) UNIQUE NOT NULL COMMENT '用户名',
    `password` VARCHAR(32) NOT NULL COMMENT '密码',
    `follow_count` BIGINT(10) UNSIGNED NOT NULL COMMENT '关注数',
    `follower_count` BIGINT(10) UNSIGNED NOT NULL COMMENT '粉丝数',
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

CREATE TABLE IF NOT EXISTS `video` (
    `id` BIGINT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `author_id` BIGINT(10) UNSIGNED NOT NULL COMMENT '发布者用户id',
    `play_url` VARCHAR(255) NOT NULL COMMENT '播放地址',
    `cover_url` VARCHAR(255) NOT NULL COMMENT '封面地址',
    `created_on` DATETIME NOT NULL COMMENT '创建时间',
    `favorite_count` BIGINT(10) UNSIGNED NOT NULL COMMENT '点赞数',
    `comment_count` BIGINT(10) UNSIGNED NOT NULL COMMENT '评论数',
    INDEX `author`(author_id),
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='视频表';

CREATE TABLE IF NOT EXISTS `favorite` (
    `id` BIGINT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT(10) UNSIGNED NOT NULL COMMENT '点赞用户id',
    `video_id` BIGINT(10) UNSIGNED NOT NULL COMMENT '点赞视频id',
    INDEX `user_video`(user_id, video_id),
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='点赞表';

CREATE TABLE IF NOT EXISTS `relation` (
    `id` BIGINT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT(10) UNSIGNED NOT NULL COMMENT '被关注用户id',
    `follower_id` BIGINT(10) UNSIGNED NOT NULL COMMENT '发起关注请求用户id',
    INDEX `user_follower`(user_id, follower_id),
    INDEX `follower_user`(follower_id, user_id),
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户关系表';

CREATE TABLE IF NOT EXISTS `comment` (
    `id` BIGINT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT(10) UNSIGNED NOT NULL COMMENT '评论用户id',
    `video_id` BIGINT(10) UNSIGNED NOT NULL COMMENT '评论视频id',
    `comment_text` text COMMENT '评论内容',
    `created_on` DATETIME NOT NULL COMMENT '创建时间',
    INDEX `video`(video_id),
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='评论表';