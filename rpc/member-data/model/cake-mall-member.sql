CREATE TABLE `wechat_user`
(
    `id`          bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键自增id',
    `number`      bigint(20) NOT NULL COMMENT '用户唯一number',
    `union_id`    varchar(128) NOT NULL COMMENT '用户微信unionId',
    `session_key` varchar(255) NOT NULL COMMENT '用户微信sessionKey',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_number` (`number`),
    KEY           `number_index` (`number`),
    KEY           `union_id_index` (`union_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4

CREATE TABLE `user_pwd`
(
    `id`          bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键自增id',
    `number`      bigint(20) NOT NULL COMMENT '用户唯一number',
    `password`    varchar(255) NOT NULL COMMENT '用户密码',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_number` (`number`),
    KEY           `number_index` (`number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4

CREATE TABLE `user`
(
    `id`          bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键自增id',
    `number`      bigint(20) NOT NULL COMMENT '用户唯一number',
    `user_name`   varchar(255) NOT NULL COMMENT '用户名称',
    `mobile`      varchar(255) DEFAULT '' COMMENT '用户手机号',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_number` (`number`),
    KEY           `number_index` (`number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4