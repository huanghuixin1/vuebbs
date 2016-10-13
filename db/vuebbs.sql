/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50505
 Source Host           : localhost
 Source Database       : vuebbs

 Target Server Type    : MySQL
 Target Server Version : 50505
 File Encoding         : utf-8

 Date: 10/13/2016 13:38:35 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `Comment`
-- ----------------------------
DROP TABLE IF EXISTS `Comment`;
CREATE TABLE `Comment` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `CDetail` varchar(512) NOT NULL COMMENT '评论内容',
  `IsDelete` smallint(6) NOT NULL DEFAULT '0',
  `CreateTime` int(11) NOT NULL COMMENT '创建时间',
  `Fk_Aid` int(11) NOT NULL COMMENT '外键_文章id',
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `articles`
-- ----------------------------
DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ATitle` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `Summary` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '内容摘要',
  `AContent` longtext COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '帖子内容',
  `Fk_Cid` int(11) NOT NULL COMMENT '外键_类别id',
  `IsDelete` smallint(1) NOT NULL,
  `CreateTime` int(11) NOT NULL COMMENT '发布时间',
  `Fk_userId` int(11) NOT NULL COMMENT '发帖用户',
  `AType` int(11) NOT NULL DEFAULT '0' COMMENT '帖子类型 0-一般帖子 1-求助帖子 2-已解决问题的帖子 3-公告',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=86 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
--  Records of `articles`
-- ----------------------------
BEGIN;
INSERT INTO `articles` VALUES ('1', '关于.net开发的展望', '圣诞快乐家里看书的加拉斯加领导', '耍的阿斯顿as大神大神大神大神大神', '1', '0', '1474935443', '1', '0'), ('2', '2关于.net开发的展望', '圣诞快乐家里看书的加拉斯加领导', '耍的阿斯顿as大神大神大神大神大神', '1', '0', '1470295353', '5', '0'), ('3', '3关于.net开发的展望', '圣诞快乐家里看书的加拉斯加领导', '耍的阿斯顿as大神大神大神大神大神', '1', '0', '1473295353', '1', '0'), ('4', '关于.net开发的展望123123123', '圣诞快乐家里看书的加拉斯加领导', '耍的阿斯顿as大神大神大神大神大神', '2', '0', '1471295353', '1', '0'), ('5', 'hhx222', '内容太摘要222', 'asdlkasjlkdajskldas', '1', '0', '1474295353', '1', '0'), ('6', '66666关于.net开发的展望123123123', '圣诞快乐家里看书的加拉斯加领导', '耍的阿斯顿as大神大神大神大神大神', '1', '0', '1474295353', '1', '0'), ('7', 'aaa关于.net开', '圣诞快乐家sadsadasd', '耍的阿斯顿as大神大神大神大神大神', '1', '0', '1474295353', '1', '0'), ('8', 'bbbb标题111', '圣诞快乐家sadsadasd', '耍的阿斯顿as大神大神大神大神大神', '1', '0', '1474295353', '1', '0'), ('9', '关于.net开发的展望', '圣诞快乐家里看书的加拉斯加领导', '耍的阿斯顿as大神大神大神大神大神', '1', '0', '1474935443', '1', '0'), ('40', '关于.net开发的展望', '圣诞快乐家里看书的加拉斯加领导', '耍的阿斯顿as大神大神大神大神大神', '1', '0', '1475974559', '1', '0'), ('41', '2关于.net开发的展望', '圣诞快乐家里看书的加拉斯加领导', '耍的阿斯顿as大神大神大神大神大神', '1', '0', '1475974559', '5', '0'), ('42', '3关于.net开发的展望', '圣诞快乐家里看书的加拉斯加领导', '耍的阿斯顿as大神大神大神大神大神', '1', '0', '1475974559', '1', '0'), ('43', '关于.net开发的展望123123123', '圣诞快乐家里看书的加拉斯加领导', '耍的阿斯顿as大神大神大神大神大神', '2', '0', '1475974559', '1', '0'), ('44', 'hhx222', '内容太摘要222', 'asdlkasjlkdajskldas', '1', '0', '1475974559', '1', '0'), ('45', '66666关于.net开发的展望123123123', '圣诞快乐家里看书的加拉斯加领导', '耍的阿斯顿as大神大神大神大神大神', '1', '0', '1475974559', '1', '0'), ('46', 'aaa关于.net开', '圣诞快乐家sadsadasd', '耍的阿斯顿as大神大神大神大神大神', '1', '0', '1475974559', '1', '0'), ('47', 'bbbb标题111', '圣诞快乐家sadsadasd', '耍的阿斯顿as大神大神大神大神大神', '1', '0', '1475974559', '1', '0'), ('48', '关于.net开发的展望', '圣诞快乐家里看书的加拉斯加领导', '耍的阿斯顿as大神大神大神大神大神', '1', '0', '1475974559', '1', '0'), ('55', '关于.net开发的展望', '圣诞快乐家里看书的加拉斯加领导', '耍的阿斯顿as大神大神大神大神大神', '1', '0', '1475974925', '1', '0'), ('56', '2关于.net开发的展望', '圣诞快乐家里看书的加拉斯加领导', '耍的阿斯顿as大神大神大神大神大神', '1', '0', '1475974925', '5', '0'), ('57', '3关于.net开发的展望', '圣诞快乐家里看书的加拉斯加领导', '耍的阿斯顿as大神大神大神大神大神', '1', '0', '1475974925', '1', '0'), ('58', '关于.net开发的展望123123123', '圣诞快乐家里看书的加拉斯加领导', '耍的阿斯顿as大神大神大神大神大神', '2', '0', '1475974925', '1', '0'), ('59', 'hhx222', '内容太摘要222', 'asdlkasjlkdajskldas', '1', '0', '1475974925', '1', '0'), ('60', '66666关于.net开发的展望123123123', '圣诞快乐家里看书的加拉斯加领导', '耍的阿斯顿as大神大神大神大神大神', '1', '0', '1475974925', '1', '0'), ('61', 'aaa关于.net开', '圣诞快乐家sadsadasd', '耍的阿斯顿as大神大神大神大神大神', '1', '0', '1475974925', '1', '0'), ('62', 'bbbb标题111', '圣诞快乐家sadsadasd', '耍的阿斯顿as大神大神大神大神大神', '1', '0', '1475974925', '1', '0'), ('63', '关于.net开发的展望', '圣诞快乐家里看书的加拉斯加领导', '耍的阿斯顿as大神大神大神大神大神', '1', '0', '1475974925', '1', '0'), ('64', '关于.net开发的展望', '圣诞快乐家里看书的加拉斯加领导', '耍的阿斯顿as大神大神大神大神大神', '1', '0', '1475974925', '1', '0'), ('65', '2关于.net开发的展望', '圣诞快乐家里看书的加拉斯加领导', '耍的阿斯顿as大神大神大神大神大神', '1', '0', '1475974925', '5', '0'), ('66', '3关于.net开发的展望', '圣诞快乐家里看书的加拉斯加领导', '耍的阿斯顿as大神大神大神大神大神', '1', '0', '1475974925', '1', '0'), ('67', '关于.net开发的展望123123123', '圣诞快乐家里看书的加拉斯加领导', '耍的阿斯顿as大神大神大神大神大神', '2', '0', '1475974925', '1', '0'), ('68', 'hhx222', '内容太摘要222', 'asdlkasjlkdajskldas', '1', '0', '1475974925', '1', '0'), ('69', '66666关于.net开发的展望123123123', '圣诞快乐家里看书的加拉斯加领导', '耍的阿斯顿as大神大神大神大神大神', '1', '0', '1475974925', '1', '0'), ('70', 'aaa关于.net开', '圣诞快乐家sadsadasd', '耍的阿斯顿as大神大神大神大神大神', '1', '0', '1475974925', '1', '0'), ('71', 'bbbb标题111', '圣诞快乐家sadsadasd', '耍的阿斯顿as大神大神大神大神大神', '1', '0', '1475974925', '1', '0'), ('72', '关于.net开发的展望', '圣诞快乐家里看书的加拉斯加领导', '耍的阿斯顿as大神大神大神大神大神', '1', '0', '1475974925', '1', '0');
COMMIT;

-- ----------------------------
--  Table structure for `categories`
-- ----------------------------
DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `CName` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '模块名字',
  `IsDelete` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否被删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
--  Records of `categories`
-- ----------------------------
BEGIN;
INSERT INTO `categories` VALUES ('1', '前端开发', b'0'), ('2', 'vps讨论', b'0'), ('3', '.net开发', b'0');
COMMIT;

-- ----------------------------
--  Table structure for `test`
-- ----------------------------
DROP TABLE IF EXISTS `test`;
CREATE TABLE `test` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `UName` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `Age` int(11) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
--  Records of `test`
-- ----------------------------
BEGIN;
INSERT INTO `test` VALUES ('37', 'hhx', '199');
COMMIT;

-- ----------------------------
--  Table structure for `userinfos`
-- ----------------------------
DROP TABLE IF EXISTS `userinfos`;
CREATE TABLE `userinfos` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `NickName` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `LoginName` varchar(16) COLLATE utf8mb4_unicode_ci NOT NULL,
  `LoginPwd` varchar(16) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '登录密码',
  `HeaderImg` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '头像',
  `EmailAddr` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `IsDelete` smallint(6) NOT NULL DEFAULT '0' COMMENT '是否被删除',
  `ULevel` int(11) NOT NULL DEFAULT '0' COMMENT '0-普通会员',
  `CreateTime` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY ```unique_loginName``` (`LoginName`) USING HASH
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
--  Records of `userinfos`
-- ----------------------------
BEGIN;
INSERT INTO `userinfos` VALUES ('1', 'hhx', 'hhx', '123', '', 'hhx__', '0', '999', '1474032388'), ('5', 'hhx2号', 'hhx22', '123', '', '123@23.123', '0', '999', '1404964116'), ('8', 'hhx2233号', 'hhx223', '123', '', '123@23.123', '0', '999', '1474512802'), ('9', 'hhx22343号', 'hhx22334', '123', '', '123@23.123', '0', '999', '1472512802');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
