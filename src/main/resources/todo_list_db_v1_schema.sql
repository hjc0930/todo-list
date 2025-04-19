-- ----------------------------
-- Copyright (c) 2025-2026 jiacheng all rights reserved.
-- ----------------------------

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

# Dump of table tb_newbee_mall_admin_user

# ------------------------------------------------------------
DROP TABLE IF EXISTS `tb_todo_list_user`;
CREATE TABLE `tb_todo_list_user` (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'User ID',
  `login_user_name` varchar(50) NOT NULL COMMENT 'Login User Name',
  `login_password` varchar(50) NOT NULL COMMENT 'Login Password',
  `phone` varchar(20) NOT NULL COMMENT 'Phone',
  `locked` tinyint(4) NOT NULL DEFAULT '0' COMMENT 'Locked (0: No, 1: Yes)',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time',
  `create_user_id` int(11) NOT NULL DEFAULT '0' COMMENT 'Create User ID',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Update Time',
  `update_user_id` int(11) NOT NULL DEFAULT '0' COMMENT 'Update User ID',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

# -- ----------------------------
DROP TABLE IF EXISTS `tb_todo_list_task`;
CREATE TABLE `tb_todo_list_task` (
  `task_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'Task ID',
  `task_name` varchar(100) NOT NULL COMMENT 'Task Name',
  `task_description` text NOT NULL COMMENT 'Task Description',
  `user_id` bigint(20) NOT NULL COMMENT 'User ID',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT 'Status (0: Pending, 1: Completed)',
  `is_deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT 'Is Deleted (0: No, 1: Yes)',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Create Time',
  `create_user_id` int(11) NOT NULL DEFAULT '0' COMMENT 'Create User ID',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Update Time',
  `update_user_id` int(11) NOT NULL DEFAULT '0' COMMENT 'Update User ID',
  PRIMARY KEY (`task_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

