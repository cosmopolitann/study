/*
 Navicat Premium Data Transfer

 Source Server         : xiaolong
 Source Server Type    : SQLite
 Source Server Version : 3030001
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3030001
 File Encoding         : 65001

 Date: 05/07/2021 11:16:34
*/

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS "article";

CREATE TABLE "article" (
  "id" VARCHAR (64) NOT NULL,--id
  "user_id" VARCHAR (64) NOT NULL,--用户id
  "accesstory" VARCHAR (600),--附件
  "accesstory_type" INTEGER (10) NOT NULL,--附件类型
  "text" text NOT NULL,--文本
  "tag" VARCHAR (64) NOT NULL,--标签
  "ptime" INTEGER (10) NOT NULL DEFAULT ( strftime( '%s', 'now' ) ),--创建时间
  "play_num" INTEGER (10) NOT NULL DEFAULT ( 0 ),--播放次数
  "share_num" INTEGER (10) NOT NULL DEFAULT ( 0 ),--分享次数
  "title" VARCHAR (128) NOT NULL,--标题
  "thumbnail" VARCHAR (128) NOT NULL,--缩略图
  "file_name" VARCHAR (128) NOT NULL,--文件名字
  "file_size" VARCHAR (128) NOT NULL,--文件大小
  PRIMARY KEY ("id"),--主键 id
  UNIQUE ("id" ASC)--唯一键 id
);
PRAGMA foreign_keys = true;
