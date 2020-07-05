/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50726
Source Host           : localhost:3306
Source Database       : ginpom

Target Server Type    : MYSQL
Target Server Version : 50726
File Encoding         : 65001

Date: 2020-07-05 23:50:47
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '姓名',
  `addr` varchar(255) NOT NULL DEFAULT '' COMMENT '住址',
  `age` smallint(4) NOT NULL DEFAULT '0' COMMENT '年龄',
  `birth` varchar(100) NOT NULL DEFAULT '' COMMENT '生日',
  `sex` smallint(4) NOT NULL DEFAULT '0' COMMENT '性别',
  `update_at` datetime NOT NULL DEFAULT '2020-01-01 00:00:00' COMMENT '更新时间',
  `create_at` datetime NOT NULL DEFAULT '2020-01-01 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';

-- ----------------------------
-- Records of user
-- ----------------------------
