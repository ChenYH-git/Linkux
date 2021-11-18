# CREATE TABLE `user` (
#                     `id` bigint(20) NOT NULL AUTO_INCREMENT,
#                     `contribution` bigint(20) NOT NULL ,
#                     `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL ,
#                     `user_id` varchar(128) NOT NULL ,
#                     `pic_link` varchar(256) NOT NULL,
#                     PRIMARY KEY (`id`),
#                     UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
# ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;
#
# CREATE TABLE `collection` (
#                     `id` bigint(20) NOT NULL AUTO_INCREMENT,
#                     `post_id` bigint(20) NOT NULL COMMENT '收藏的帖子id',
#                     `user_id` varchar(128) NOT NULL ,
#                     PRIMARY KEY (`id`),
#                     KEY `idx_user_id` (`user_id`)
# ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;
#
# CREATE TABLE `follow` (
#                               `id` bigint(20) NOT NULL AUTO_INCREMENT,
#                               `follow_id` varchar(128) NOT NULL COMMENT '关注的作者id',
#                               `user_id` varchar(128) NOT NULL ,
#                               PRIMARY KEY (`id`),
#                               KEY `idx_user_id` (`user_id`)
# ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;
#
# CREATE TABLE `followed` (
#                               `id` bigint(20) NOT NULL AUTO_INCREMENT,
#                               `followed_id` varchar(128) NOT NULL COMMENT '粉丝id',
#                               `user_id` varchar(128) NOT NULL ,
#                               PRIMARY KEY (`id`),
#                               KEY `idx_user_id` (`user_id`)
# ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;
#
# DROP TABLE IF EXISTS `label`;
# CREATE TABLE label (
#     `id` int(11) NOT NULL AUTO_INCREMENT,
#     `label_id` int(10) unsigned NOT NULL ,
#     `label_name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL ,
#     `introduction` varchar(256) COLLATE utf8mb4_general_ci NOT NULL ,
#     `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
#     `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
#      PRIMARY KEY (`id`),
#      UNIQUE KEY `idx_label_id` (`label_id`),
#      UNIQUE KEY `idx_label_name` (`label_name`)
# ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 Collate = utf8mb4_general_ci;
#
# DROP TABLE IF EXISTS `post`;
# CREATE TABLE  `post` (
#     `id` bigint(20) NOT NULL AUTO_INCREMENT,
#     `post_id` bigint(20) NOT NULL COMMENT '帖子id',
#     `trans_id` bigint(20) NOT NUlL DEFAULT '0' COMMENT '翻译任务的id',
#     `title` varchar(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
#     `content` varchar(8192) COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
#     `author_id` bigint(20) NOT NULL COMMENT '作者的用户id',
#     `label_id` bigint(20) NOT NULL COMMENT '所属标签',
#     `collect_num` bigint(20) NOT NULL DEFAULT '0' COMMENT '帖子收藏量',
#     `viewd_num` bigint(20) NOT NULL DEFAULT '0' COMMENT '帖子观看量',
#     `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '帖子状态',
#     `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
#     `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
#     PRIMARY KEY (`id`),
#     UNIQUE KEY `idx_post_id` (`post_id`),
#     KEY `idx_author_id` (`author_id`),
#     KEY `idx_community_id` (`label_id`)
# ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
#
#
# DROP TABLE IF EXISTS `trans`;
# CREATE TABLE  `trans` (
#     `id` bigint(20) NOT NULL AUTO_INCREMENT,
#     `trans_id` bigint(20) NOT NULL COMMENT '翻译任务id',
#     `title` varchar(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
#     `content` varchar(8192) COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
#     `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '任务状态',
#     `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
#     `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
#     PRIMARY KEY (`id`),
#     UNIQUE KEY `idx_trans_id` (`trans_id`)
# ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
#
#
# INSERT INTO `label` VALUES ('1', '1', 'Ubuntu', '-', '2021-11-13 17:13:00', '2021-11-13 17:13:00');
# INSERT INTO `label` VALUES ('2', '2', 'Mint', '-', '2021-11-13 17:13:00', '2021-11-13 17:13:00');
# INSERT INTO `label` VALUES ('3', '3', 'kali', '-', '2021-11-13 17:13:00', '2021-11-13 17:13:00');
# INSERT INTO `label` VALUES ('4', '4', 'MX Linux', '-', '2021-11-13 17:13:00', '2021-11-13 17:13:00')
#
