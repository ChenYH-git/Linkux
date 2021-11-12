# CREATE TABLE `user` (
#                     `id` bigint(20) NOT NULL AUTO_INCREMENT,
#                     `contribution` bigint(20) NOT NULL ,
#                     `user_id` varchar(128) NOT NULL ,
#                     PRIMARY KEY (`id`),
#                     UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
# ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;
#
# DROP TABLE IF EXISTS `label`;
# CREATE TABLE label (
#     `id` int(11) NOT NULL AUTO_INCREMENT,
#     `community_id` int(10) unsigned NOT NULL ,
#     `community_name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL ,
#     `introduction` varchar(256) COLLATE utf8mb4_general_ci NOT NULL ,
#     `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
#     `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
#      PRIMARY KEY (`id`),
#      UNIQUE KEY `idx_community_id` (`community_id`),
#      UNIQUE KEY `idx_commnity_name` (`community_name`)
# ) ENGINE = InnoB DEFAULT CHARSET = utf8mb4 Collate = utf8mb4_general_ci;

DROP TABLE IF EXISTS `post`;
CREATE TABLE  `post` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `post_id` bigint(20) NOT NULL COMMENT '帖子id',
    `title` varchar(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
    `content` varchar(8192) COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
    `author_id` bigint(20) NOT NULL COMMENT '作者的用户id',
    `label_id` bigint(20) NOT NULL COMMENT '所属标签',
    `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '帖子状态',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_post_id` (`post_id`),
    KEY `idx_author_id` (`author_id`),
    KEY `idx_community_id` (`label_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
#
#
# INSERT INTO `community` VALUES ('1', '1', 'Go', 'Golang', '2021-03-13 17:13:00', '2021-03-13 17:13:00')
# INSERT INTO `community` VALUES ('2', '2', 'leetcode', '刷题！', '2021-03-13 17:13:00', '2021-03-13 17:13:00')
# INSERT INTO `community` VALUES ('3', '3', 'csgo', 'Rush Rush A', '2021-03-13 17:13:00', '2021-03-13 17:13:00')
# INSERT INTO `community` VALUES ('4', '4', 'LOL', 'lol赛高', '2021-03-13 17:13:00', '2021-03-13 17:13:00')

