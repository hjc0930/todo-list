-- todo_list.todo_items definition
CREATE TABLE
  `todo_items` (
    `id` bigint (20) NOT NULL AUTO_INCREMENT,
    `list_id` bigint (20) NOT NULL COMMENT '所属列表ID',
    `content` text NOT NULL COMMENT '事项内容',
    `position` int (11) DEFAULT '0' COMMENT '列表内排序',
    `is_completed` tinyint (4) DEFAULT '0' COMMENT '完成状态：0-未完成，1-已完成',
    `created_time` datetime DEFAULT CURRENT_TIMESTAMP,
    `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_time` datetime DEFAULT NULL COMMENT '软删除时间',
    `create_user` bigint (20) DEFAULT NULL COMMENT '创建人ID',
    `update_user` bigint (20) DEFAULT NULL COMMENT '更新人ID',
    `delete_user` bigint (20) DEFAULT NULL COMMENT '删除人ID',
    PRIMARY KEY (`id`)
  ) ENGINE = InnoDB AUTO_INCREMENT = 2 DEFAULT CHARSET = utf8mb4;

-- todo_list.todo_lists definition
CREATE TABLE
  `todo_lists` (
    `id` bigint (20) NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NOT NULL COMMENT '列表名称',
    `status` tinyint (4) DEFAULT '0' COMMENT '状态：0-进行中，1-已完成，2-已归档',
    `created_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_time` datetime DEFAULT NULL COMMENT '软删除时间',
    `create_user` bigint (20) DEFAULT NULL COMMENT '创建人ID',
    `update_user` bigint (20) DEFAULT NULL COMMENT '更新人ID',
    `delete_user` bigint (20) DEFAULT NULL COMMENT '删除人ID',
    PRIMARY KEY (`id`)
  ) ENGINE = InnoDB AUTO_INCREMENT = 3 DEFAULT CHARSET = utf8mb4;

-- todo_list.user_auths definition
CREATE TABLE
  `user_auths` (
    `id` bigint (20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint (20) NOT NULL COMMENT '关联用户ID',
    `identity_type` varchar(20) NOT NULL COMMENT '登录类型（EMAIL/PHONE/WEICHAT）',
    `identifier` varchar(100) NOT NULL COMMENT '标识（邮箱/手机号/第三方ID）',
    `credential` varchar(200) NOT NULL COMMENT '凭证（哈希密码/Token）',
    `verified` tinyint (4) DEFAULT '0' COMMENT '是否验证：0-否，1-是',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_auths` (`user_id`, `identifier`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- todo_list.users definition
CREATE TABLE
  `users` (
    `id` bigint (20) NOT NULL AUTO_INCREMENT,
    `username` varchar(50) NOT NULL COMMENT '昵称',
    `avatar` varchar(200) DEFAULT NULL COMMENT '头像URL',
    `created_time` datetime DEFAULT CURRENT_TIMESTAMP,
    `updated_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_time` datetime DEFAULT NULL COMMENT '软删除时间',
    `create_user` bigint (20) DEFAULT NULL COMMENT '创建人ID',
    `update_user` bigint (20) DEFAULT NULL COMMENT '更新人ID',
    `delete_user` bigint (20) DEFAULT NULL COMMENT '删除人ID',
    PRIMARY KEY (`id`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;