CREATE TABLE IF NOT EXISTS wechat_user
(
    `id`          BIGINT       NOT NULL AUTO_INCREMENT COMMENT '主键自增id',
    `number`      BIGINT       NOT NULL COMMENT '用户唯一number',
    `union_id`    VARCHAR(255) NOT NULL COMMENT '用户微信unionId',
    `session_key` VARCHAR(255) NOT NULL COMMENT '用户微信sessionKey',
    `createTime`  timestamp    NULL DEFAULT CURRENT_TIMESTAMP,
    `updateTime`  timestamp    NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;