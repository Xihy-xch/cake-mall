CREATE TABLE `wechat_user`
(
    `id`          bigint       NOT NULL AUTO_INCREMENT COMMENT '主键自增id',
    `number`      bigint       NOT NULL COMMENT '用户唯一number',
    `union_id`    varchar(255) NOT NULL COMMENT '用户微信unionId',
    `session_key` varchar(255) NOT NULL COMMENT '用户微信sessionKey',
    `createTime`  timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updateTime`  timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY           `union_id_index` (`union_id`),
    KEY           `number_index` (`number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `user_pwd`
(
    `id`         bigint       NOT NULL AUTO_INCREMENT COMMENT '主键自增id',
    `number`     bigint       NOT NULL COMMENT '用户唯一number',
    `password`   varchar(255) NOT NULL COMMENT '用户密码',
    `createTime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updateTime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY          `number_index` (`number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci