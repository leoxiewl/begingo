CREATE TABLE `users` (
                         `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
                         `nickname` varchar(255) NOT NULL DEFAULT '' COMMENT '用户昵称',
                         `email` varchar(255) NOT NULL COMMENT '邮箱',
                         `password` varchar(255) NOT NULL COMMENT '密码',
                         `avatar` varchar(255) NOT NULL  DEFAULT '' COMMENT '用户头像',
                         `gender`        tinyint       NOT NULL DEFAULT 0 COMMENT '性别 0-未知，1-男，2-女',
                         `user_role`     varchar(256)  NOT NULL DEFAULT 'user' COMMENT '用户角色：user / admin',
                         `create_at`   datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                         `update_at`   datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                         `delete_at`   datetime DEFAULT NULL COMMENT '删除时间',
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;